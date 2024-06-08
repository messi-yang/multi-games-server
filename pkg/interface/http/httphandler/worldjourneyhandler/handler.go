package worldjourneyhandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httpsession"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	ErrCommandIsNotExecutedByOwnPlayer = fmt.Errorf("the command is not executed by its own player")
)

var websocketUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type HttpHandler struct {
	redisServerMessageMediator redisservermessagemediator.Mediator
}

func NewHttpHandler(redisServerMessageMediator redisservermessagemediator.Mediator) *HttpHandler {
	return &HttpHandler{
		redisServerMessageMediator: redisServerMessageMediator,
	}
}

func (httpHandler *HttpHandler) StartJourney(c *gin.Context) {
	socketConn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to upgrade http to socket")
		return
	}
	defer socketConn.Close()

	respondMessageLocker := &sync.RWMutex{}
	respondMessage := func(jsonObj any) {
		respondMessageLocker.Lock()
		defer respondMessageLocker.Unlock()

		messageJsonInBytes := jsonutil.Marshal(jsonObj)
		err = socketConn.WriteMessage(2, messageJsonInBytes)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	var closeConnFlag sync.WaitGroup
	closeConnFlag.Add(1)
	closeConnection := func() {
		closeConnFlag.Done()
	}

	var worldIdDto uuid.UUID
	var myPlayerIdDto uuid.UUID
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)

	respondServerEvent := func(serverEvent any) {
		respondMessage(serverEvent)
	}

	broadcastServerEvent := func(serverEvent any) {
		httpHandler.redisServerMessageMediator.Send(
			newWorldMessageChannel(worldIdDto),
			jsonutil.Marshal(worldMessage{
				SenderId:    myPlayerIdDto,
				ServerEvent: serverEvent,
			}),
		)
	}

	respondAndBroadcastServerEvent := func(serverEvent any) {
		respondServerEvent(serverEvent)
		broadcastServerEvent(serverEvent)
	}

	sendServerEventToPlayer := func(playerIdDto uuid.UUID, serverEvent any) {
		httpHandler.redisServerMessageMediator.Send(
			newPlayerMessageChannel(worldIdDto, playerIdDto),
			jsonutil.Marshal(playerMessage{
				ServerEvent: serverEvent,
			}),
		)
	}

	generatePlayerJoinedServerEvent := func(playerDto dto.PlayerDto) playerJoinedServerEvent {
		return playerJoinedServerEvent{
			Name:   playerJoinedServerEventName,
			Player: playerDto,
		}
	}

	generateWorldEnteredServerEvent := func(worldDto dto.WorldDto, unitDtos []dto.UnitDto, playerDtos []dto.PlayerDto) worldEnteredServerEvent {
		return worldEnteredServerEvent{
			Name:       worldEnteredServerEventName,
			World:      viewmodel.WorldViewModel(worldDto),
			Units:      unitDtos,
			MyPlayerId: myPlayerIdDto,
			Players:    playerDtos,
		}
	}

	generatePlayerLeftServerEvent := func(playerIdDto uuid.UUID) playerLeftServerEvent {
		return playerLeftServerEvent{
			Name:     playerLeftServerEventName,
			PlayerId: playerIdDto,
		}
	}

	generateCommandReceivedServerEvent := func(command any) commandReceivedServerEvent {
		return commandReceivedServerEvent{
			Name:    commandReceivedServerEventName,
			Command: command,
		}
	}

	generateCommandFailedServerEvent := func(commandId uuid.UUID) commandFailedServerEvent {
		return commandFailedServerEvent{
			Name:      commandFailedServerEventName,
			CommandId: commandId,
		}
	}

	generateP2pOfferReceivedServerEvent := func(peerPlayerId uuid.UUID, iceCandidates []any, offer any) p2pOfferReceivedServerEvent {
		return p2pOfferReceivedServerEvent{
			Name:          p2pOfferReceivedServerEventName,
			PeerPlayerId:  peerPlayerId,
			IceCandidates: iceCandidates,
			Offer:         offer,
		}
	}

	generateP2pAnswerReceivedServerEvent := func(peerPlayerId uuid.UUID, iceCandidates []any, offer any) p2pAnswerReceivedServerEvent {
		return p2pAnswerReceivedServerEvent{
			Name:          p2pAnswerReceivedServerEventName,
			PeerPlayerId:  peerPlayerId,
			IceCandidates: iceCandidates,
			Answer:        offer,
		}
	}
	generateErroredServerEvent := func(err error) erroredServerEvent {
		return erroredServerEvent{
			Name:    erroredServerEventName,
			Message: err.Error(),
		}
	}

	if worldIdDto, err = uuid.Parse(c.Request.URL.Query().Get("id")); err != nil {
		respondServerEvent(generateErroredServerEvent(err))
		closeConnection()
		return
	}

	myPlayerDto, err := httpHandler.createPlayer(worldIdDto, authorizedUserIdDto)
	if err != nil {
		respondServerEvent(generateErroredServerEvent(err))
		closeConnection()
		return
	}
	myPlayerIdDto = myPlayerDto.Id
	broadcastServerEvent(generatePlayerJoinedServerEvent(myPlayerDto))

	// Subscribe to messages broadcasted from other websocket servers
	worldMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newWorldMessageChannel(worldIdDto),
		func(messageBytes []byte) {
			message, err := jsonutil.Unmarshal[worldMessage](messageBytes)
			if err != nil {
				return
			}
			if message.SenderId == myPlayerIdDto {
				return
			}

			respondMessage(message.ServerEvent)
		},
	)
	defer worldMessageUnusbscriber()

	// Subscribe to messages sent specifically to your player
	playerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newPlayerMessageChannel(worldIdDto, myPlayerIdDto),
		func(messageBytes []byte) {
			message, err := jsonutil.Unmarshal[playerMessage](messageBytes)
			if err != nil {
				return
			}

			respondMessage(message.ServerEvent)
		},
	)
	defer playerMessageUnusbscriber()

	safelyLeaveWorldInAllCases := func() {
		if err = httpHandler.removePlayer(worldIdDto, myPlayerIdDto); err != nil {
			fmt.Println(err)
		}
		broadcastServerEvent(generatePlayerLeftServerEvent(myPlayerIdDto))

	}
	defer safelyLeaveWorldInAllCases()

	worldDto, unitDtos, playerDtos, err := httpHandler.getWorldInformation(worldIdDto)
	if err != nil {
		respondServerEvent(generateErroredServerEvent(err))
		closeConnection()
		return
	}
	respondServerEvent(generateWorldEnteredServerEvent(worldDto, unitDtos, playerDtos))

	go func() {
		for {
			_, message, err := socketConn.ReadMessage()
			if err != nil {
				respondServerEvent(generateErroredServerEvent(err))
				closeConnection()
				return
			}

			clientEvent, err := jsonutil.Unmarshal[clientEvent](message)
			if err != nil {
				respondServerEvent(generateErroredServerEvent(err))
				return
			}

			switch clientEvent.Name {
			case pingClientEventName:
				continue
			case p2pOfferSentClientEventName:
				clientEvent, err := jsonutil.Unmarshal[p2pOfferSentClientEvent](message)
				if err != nil {
					respondServerEvent(generateErroredServerEvent(err))
					return
				}
				sendServerEventToPlayer(clientEvent.PeerPlayerId, generateP2pOfferReceivedServerEvent(myPlayerIdDto, clientEvent.IceCandidates, clientEvent.Offer))
			case p2pAnswerSentClientEventName:
				clientEvent, err := jsonutil.Unmarshal[p2pAnswerSentClientEvent](message)
				if err != nil {
					respondServerEvent(generateErroredServerEvent(err))
					return
				}
				sendServerEventToPlayer(clientEvent.PeerPlayerId, generateP2pAnswerReceivedServerEvent(myPlayerIdDto, clientEvent.IceCandidates, clientEvent.Answer))
			case commandSentClientEventName:
				clientEvent, err := jsonutil.Unmarshal[commandSentClientEvent](message)
				if err != nil {
					respondServerEvent(generateErroredServerEvent(err))
					return
				}
				sendServerEventToPlayer(clientEvent.PeerPlayerId, generateCommandReceivedServerEvent(clientEvent.Command))
			case commandRequestedClientEventName:
				clientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[command]](message)
				if err != nil {
					respondServerEvent(generateErroredServerEvent(err))
					return
				}
				switch clientEvent.Command.Name {
				case createStaticUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createStaticUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateStaticUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
					); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case createFenceUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createFenceUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateFenceUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
					); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case createPortalUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createPortalUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreatePortalUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
					); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case createLinkUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createLinkUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateLinkUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
						commandDto.Label,
						commandDto.Url,
					); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case createEmbedUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createEmbedUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateEmbedUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
						commandDto.Label,
						commandDto.EmbedCode,
					); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case rotateUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[rotateUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRotateUnitCommand(commandDto.UnitId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case removeStaticUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeStaticUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveStaticUnitCommand(commandDto.UnitId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case removeFenceUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeFenceUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveFenceUnitCommand(commandDto.UniId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case removePortalUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removePortalUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemovePortalUnitCommand(commandDto.UnitId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case removeLinkUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeLinkUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveLinkUnitCommand(commandDto.UnitId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case removeEmbedUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeEmbedUnitCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveEmbedUnitCommand(commandDto.UnitId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				default:
				}
			}

		}
	}()

	closeConnFlag.Wait()
}

func (httpHandler *HttpHandler) executeCreateStaticUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()
	createStaticUnitUseCase := usecase.ProvideCreateStaticUnitUseCase(uow)
	if err := createStaticUnitUseCase.Execute(idDto, worldIdDto, itemIdDto, positionDto, directionDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveStaticUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()
	removeStaticUnitUseCase := usecase.ProvideRemoveStaticUnitUseCase(uow)
	if err := removeStaticUnitUseCase.Execute(idDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeCreateFenceUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()
	createFenceUnitUseCase := usecase.ProvideCreateFenceUnitUseCase(uow)
	if err := createFenceUnitUseCase.Execute(idDto, worldIdDto, itemIdDto, positionDto, directionDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveFenceUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()
	removeFenceUnitUseCase := usecase.ProvideRemoveFenceUnitUseCase(uow)
	if err := removeFenceUnitUseCase.Execute(idDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeCreatePortalUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()
	createPortalUnitUseCase := usecase.ProvideCreatePortalUnitUseCase(uow)
	if err := createPortalUnitUseCase.Execute(idDto, worldIdDto, itemIdDto, positionDto, directionDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemovePortalUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()
	removePortalUnitUseCase := usecase.ProvideRemovePortalUnitUseCase(uow)
	if err := removePortalUnitUseCase.Execute(idDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeCreateLinkUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto dto.PositionDto,
	directionDto int8,
	labelDto *string,
	urlDto string,
) error {
	uow := pguow.NewUow()
	createLinkUnitUseCase := usecase.ProvideCreateLinkUnitUseCase(uow)
	if err := createLinkUnitUseCase.Execute(idDto, worldIdDto, itemIdDto, positionDto, directionDto, labelDto, urlDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveLinkUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()
	removeLinkUnitUseCase := usecase.ProvideRemoveLinkUnitUseCase(uow)
	if err := removeLinkUnitUseCase.Execute(idDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeCreateEmbedUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto dto.PositionDto,
	directionDto int8,
	labelDto *string,
	embedCodeDto string,
) error {
	uow := pguow.NewUow()
	createEmbedUnitUseCase := usecase.ProvideCreateEmbedUnitUseCase(uow)
	if err := createEmbedUnitUseCase.Execute(idDto, worldIdDto, itemIdDto, positionDto, directionDto, labelDto, embedCodeDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveEmbedUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()
	removeEmbedUnitUseCase := usecase.ProvideRemoveEmbedUnitUseCase(uow)
	if err := removeEmbedUnitUseCase.Execute(idDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRotateUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	rotateUnitUseCase := usecase.ProvideRotateUnitUseCase(uow)
	if err := rotateUnitUseCase.Execute(idDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) removePlayer(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	uow := pguow.NewUow()
	removePlayerUseCase := usecase.ProvideRemovePlayerUseCase(uow)
	if err := removePlayerUseCase.Execute(worldIdDto, playerIdDto); err != nil {
		uow.RevertChanges()
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) createPlayer(worldIdDto uuid.UUID, userIdDto *uuid.UUID) (newPlayerDto dto.PlayerDto, err error) {
	uow := pguow.NewUow()

	createPlayerUseCase := usecase.ProvideCreatePlayerUseCase(uow)
	newPlayerDto, err = createPlayerUseCase.Execute(worldIdDto, userIdDto)
	if err != nil {
		uow.RevertChanges()
		return newPlayerDto, err
	}
	uow.SaveChanges()

	fmt.Println(newPlayerDto)

	return newPlayerDto, nil
}

func (httpHandler *HttpHandler) getWorldInformation(worldIdDto uuid.UUID) (
	worldDto dto.WorldDto, unitDtos []dto.UnitDto, playerDtos []dto.PlayerDto, err error,
) {
	uow := pguow.NewDummyUow()

	getWorldInformationUseCase := usecase.ProvideGetWorldInformationUseCase(uow)
	return getWorldInformationUseCase.Execute(worldIdDto)
}
