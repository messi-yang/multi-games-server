package worldjourneyhandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	iam_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	world_dto "github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/embedunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/fenceunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/linkunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/playerappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/portalunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/staticunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/unitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	world_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
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
	var playerIdDto uuid.UUID
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)

	respondServerEvent := func(serverEvent any) {
		respondMessage(serverEvent)
	}

	broadcastServerEvent := func(serverEvent any) {
		httpHandler.redisServerMessageMediator.Send(
			newWorldMessageChannel(worldIdDto),
			jsonutil.Marshal(worldMessage{
				SenderId:    playerIdDto,
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

	generatePlayerJoinedServerEvent := func(playerDto world_dto.PlayerDto) playerJoinedServerEvent {
		return playerJoinedServerEvent{
			Name:   playerJoinedServerEventName,
			Player: playerDto,
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

	playerIdDto, err = httpHandler.enterWorld(worldIdDto, authorizedUserIdDto)
	if err != nil {
		respondServerEvent(generateErroredServerEvent(err))
		closeConnection()
		return
	}

	// Subscribe to messages broadcasted from other websocket servers
	worldMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newWorldMessageChannel(worldIdDto),
		func(messageBytes []byte) {
			message, err := jsonutil.Unmarshal[worldMessage](messageBytes)
			if err != nil {
				return
			}
			if message.SenderId == playerIdDto {
				return
			}

			respondMessage(message.ServerEvent)
		},
	)
	defer worldMessageUnusbscriber()

	// Subscribe to messages sent specifically to your player
	playerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newPlayerMessageChannel(worldIdDto, playerIdDto),
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
		if err = httpHandler.executeLeaveWorldCommand(worldIdDto, playerIdDto); err != nil {
			fmt.Println(err)
		}
		broadcastServerEvent(generatePlayerLeftServerEvent(playerIdDto))

	}
	defer safelyLeaveWorldInAllCases()

	if err = httpHandler.respondWorldEnteredServerEvent(worldIdDto, playerIdDto, respondMessage); err != nil {
		respondServerEvent(generateErroredServerEvent(err))
		closeConnection()
		return
	}

	playerDto, err := httpHandler.queryPlayer(worldIdDto, playerIdDto)
	if err != nil {
		respondServerEvent(generateErroredServerEvent(err))
		closeConnection()
		return
	}
	broadcastServerEvent(generatePlayerJoinedServerEvent(playerDto))

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
				sendServerEventToPlayer(clientEvent.PeerPlayerId, generateP2pOfferReceivedServerEvent(playerIdDto, clientEvent.IceCandidates, clientEvent.Offer))
			case p2pAnswerSentClientEventName:
				clientEvent, err := jsonutil.Unmarshal[p2pAnswerSentClientEvent](message)
				if err != nil {
					respondServerEvent(generateErroredServerEvent(err))
					return
				}
				sendServerEventToPlayer(clientEvent.PeerPlayerId, generateP2pAnswerReceivedServerEvent(playerIdDto, clientEvent.IceCandidates, clientEvent.Answer))
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
				case changePlayerActionCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[changePlayerActionCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if commandDto.PlayerId != playerIdDto {
						respondServerEvent(generateErroredServerEvent(ErrCommandIsNotExecutedByOwnPlayer))
						return
					}

					if err = httpHandler.executeChangePlayerActionCommand(
						worldIdDto, commandDto.PlayerId, commandDto.Action,
					); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case sendPlayerIntoPortalCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[sendPlayerIntoPortalCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if commandDto.PlayerId != playerIdDto {
						respondServerEvent(generateErroredServerEvent(ErrCommandIsNotExecutedByOwnPlayer))
						return
					}

					if err = httpHandler.executeSendPlayerIntoPortalCommand(worldIdDto, commandDto.PlayerId, commandDto.UnitId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
				case changePlayerHeldItemCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[changePlayerHeldItemCommand]](message)
					if err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if commandDto.PlayerId != playerIdDto {
						respondServerEvent(generateErroredServerEvent(ErrCommandIsNotExecutedByOwnPlayer))
						return
					}

					if err = httpHandler.executeChangePlayerHeldItemCommand(worldIdDto, commandDto.PlayerId, commandDto.ItemId); err != nil {
						respondServerEvent(generateErroredServerEvent(err))
						respondAndBroadcastServerEvent(generateCommandFailedServerEvent(commandDto.Id))
						break
					}
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

func (httpHandler *HttpHandler) queryPlayer(worldIdDto uuid.UUID, playerIdDto uuid.UUID) (playerDto world_dto.PlayerDto, err error) {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	return playerAppService.GetPlayer(playerappsrv.GetPlayerQuery{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	})
}

func (httpHandler *HttpHandler) executeChangePlayerActionCommand(
	worldIdDto uuid.UUID, playerIdDto uuid.UUID, actionDto world_dto.PlayerActionDto,
) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	if err := playerAppService.ChangePlayerAction(playerappsrv.ChangePlayerActionCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		Action:   actionDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) executeSendPlayerIntoPortalCommand(
	worldIdDto uuid.UUID,
	playerIdDto uuid.UUID,
	unitIdDto uuid.UUID,
) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	return playerAppService.SendPlayerIntoPortal(playerappsrv.SendPlayerIntoPortalCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		UnitId:   unitIdDto,
	})
}

func (httpHandler *HttpHandler) executeChangePlayerHeldItemCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, itemIdDto uuid.UUID) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	if err := playerAppService.ChangePlayerHeldItem(playerappsrv.ChangePlayerHeldItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		ItemId:   itemIdDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) executeCreateStaticUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	staticUnitAppService := world_provide_dependency.ProvideStaticUnitAppService(uow)
	if err := staticUnitAppService.CreateStaticUnit(staticunitappsrv.CreateStaticUnitCommand{
		Id:        idDto,
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveStaticUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	staticUnitAppService := world_provide_dependency.ProvideStaticUnitAppService(uow)
	if err := staticUnitAppService.RemoveStaticUnit(staticunitappsrv.RemoveStaticUnitCommand{
		Id: idDto,
	}); err != nil {
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
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	fenceUnitAppService := world_provide_dependency.ProvideFenceUnitAppService(uow)
	if err := fenceUnitAppService.CreateFenceUnit(fenceunitappsrv.CreateFenceUnitCommand{
		Id:        idDto,
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveFenceUnitCommand(unidIdDto uuid.UUID) error {
	uow := pguow.NewUow()

	fenceUnitAppService := world_provide_dependency.ProvideFenceUnitAppService(uow)
	if err := fenceUnitAppService.RemoveFenceUnit(fenceunitappsrv.RemoveFenceUnitCommand{
		Id: unidIdDto,
	}); err != nil {
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
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	portalUnitAppService := world_provide_dependency.ProvidePortalUnitAppService(uow)
	if err := portalUnitAppService.CreatePortalUnit(portalunitappsrv.CreatePortalUnitCommand{
		Id:        idDto,
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemovePortalUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	portalUnitAppService := world_provide_dependency.ProvidePortalUnitAppService(uow)
	if err := portalUnitAppService.RemovePortalUnit(portalunitappsrv.RemovePortalUnitCommand{
		Id: idDto,
	}); err != nil {
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
	positionDto world_dto.PositionDto,
	directionDto int8,
	label *string,
	url string,
) error {
	uow := pguow.NewUow()

	linkUnitAppService := world_provide_dependency.ProvideLinkUnitAppService(uow)
	if err := linkUnitAppService.CreateLinkUnit(linkunitappsrv.CreateLinkUnitCommand{
		Id:        idDto,
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
		Label:     label,
		Url:       url,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveLinkUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	linkUnitAppService := world_provide_dependency.ProvideLinkUnitAppService(uow)
	if err := linkUnitAppService.RemoveLinkUnit(linkunitappsrv.RemoveLinkUnitCommand{
		Id: idDto,
	}); err != nil {
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
	positionDto world_dto.PositionDto,
	directionDto int8,
	label *string,
	embedCode string,
) error {
	uow := pguow.NewUow()

	embedUnitAppService := world_provide_dependency.ProvideEmbedUnitAppService(uow)
	if err := embedUnitAppService.CreateEmbedUnit(embedunitappsrv.CreateEmbedUnitCommand{
		Id:        idDto,
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
		Label:     label,
		EmbedCode: embedCode,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveEmbedUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	embedUnitAppService := world_provide_dependency.ProvideEmbedUnitAppService(uow)
	if err := embedUnitAppService.RemoveEmbedUnit(embedunitappsrv.RemoveEmbedUnitCommand{
		Id: idDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRotateUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.RotateUnit(unitappsrv.RotateUnitCommand{
		Id: idDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeLeaveWorldCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	if err := playerAppService.LeaveWorld(playerappsrv.LeaveWorldCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) enterWorld(worldIdDto uuid.UUID, userIdDto *uuid.UUID) (playerIdDto uuid.UUID, err error) {
	uow := pguow.NewUow()

	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	itemAppService := world_provide_dependency.ProvideItemAppService(uow)
	userAppService := iam_provide_dependency.ProvideUserAppService(uow)

	itemDtos, err := itemAppService.QueryItems(itemappsrv.QueryItemsQuery{})
	if err != nil {
		uow.RevertChanges()
		return playerIdDto, err
	}

	playerName := "Guest"
	if userIdDto != nil {
		user, err := userAppService.GetUser(userappsrv.GetUserQuery{
			UserId: *userIdDto,
		})
		if err != nil {
			uow.RevertChanges()
			return playerIdDto, err
		}
		playerName = user.FriendlyName
	}

	if playerIdDto, err = playerAppService.EnterWorld(playerappsrv.EnterWorldCommand{
		WorldId:          worldIdDto,
		PlayerName:       playerName,
		PlayerHeldItemId: itemDtos[0].Id,
	}); err != nil {
		uow.RevertChanges()
		return playerIdDto, err
	}

	uow.SaveChanges()

	return playerIdDto, nil
}

func (httpHandler *HttpHandler) respondWorldEnteredServerEvent(worldIdDto uuid.UUID, playerIdDto uuid.UUID, respondMessage func(any)) error {
	uow := pguow.NewDummyUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	worldAppService := world_provide_dependency.ProvideWorldAppService(uow)
	playerAppService := world_provide_dependency.ProvidePlayerAppService()

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		return err
	}

	unitDtos, err := unitAppService.GetUnits(
		unitappsrv.GetUnitsQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return err
	}

	playerDtos, err := playerAppService.GetPlayers(
		playerappsrv.GetPlayersQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return err
	}

	respondMessage(worldEnteredServerEvent{
		Name:       worldEnteredServerEventName,
		World:      viewmodel.WorldViewModel(worldDto),
		Units:      unitDtos,
		MyPlayerId: playerIdDto,
		Players:    playerDtos,
	})
	return nil
}
