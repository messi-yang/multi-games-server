package worldjourneyhandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
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
			jsonutil.Marshal(worldMessage[any]{
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
			jsonutil.Marshal(playerMessage[any]{
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

	generateWorldEnteredServerEvent := func(worldDto dto.WorldDto, playerDtos []dto.PlayerDto) worldEnteredServerEvent {
		return worldEnteredServerEvent{
			Name:       worldEnteredServerEventName,
			World:      viewmodel.WorldViewModel(worldDto),
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
			message, err := jsonutil.Unmarshal[worldMessage[serverEvent]](messageBytes)
			if err != nil {
				return
			}
			if message.SenderId == myPlayerIdDto {
				return
			}

			if message.ServerEvent.Name == worldEnteredServerEventName {
				worldEnteredServerMessage, err := jsonutil.Unmarshal[worldMessage[worldEnteredServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(worldEnteredServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == playerJoinedServerEventName {
				playerJoinedServerEvent, err := jsonutil.Unmarshal[worldMessage[playerJoinedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(playerJoinedServerEvent.ServerEvent)
			} else if message.ServerEvent.Name == playerLeftServerEventName {
				playerLeftServerMessage, err := jsonutil.Unmarshal[worldMessage[playerLeftServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(playerLeftServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == commandReceivedServerEventName {
				commandReceivedServerMessage, err := jsonutil.Unmarshal[worldMessage[commandReceivedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(commandReceivedServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == commandFailedServerEventName {
				commandFailedServerMessage, err := jsonutil.Unmarshal[worldMessage[commandFailedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(commandFailedServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == erroredServerEventName {
				erroredServerMessage, err := jsonutil.Unmarshal[worldMessage[erroredServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(erroredServerMessage.ServerEvent)
			}
		},
	)
	defer worldMessageUnusbscriber()

	// Subscribe to messages sent specifically to your player
	playerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newPlayerMessageChannel(worldIdDto, myPlayerIdDto),
		func(messageBytes []byte) {
			message, err := jsonutil.Unmarshal[playerMessage[serverEvent]](messageBytes)
			if err != nil {
				return
			}

			if message.ServerEvent.Name == p2pAnswerReceivedServerEventName {
				p2pAnswerReceivedServerMessage, err := jsonutil.Unmarshal[playerMessage[p2pAnswerReceivedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(p2pAnswerReceivedServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == p2pOfferReceivedServerEventName {
				p2pOfferReceivedServerMessage, err := jsonutil.Unmarshal[playerMessage[p2pOfferReceivedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(p2pOfferReceivedServerMessage.ServerEvent)
			}
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

	worldDto, playerDtos, err := httpHandler.getWorldInformation(worldIdDto)
	if err != nil {
		respondServerEvent(generateErroredServerEvent(err))
		closeConnection()
		return
	}
	respondServerEvent(generateWorldEnteredServerEvent(worldDto, playerDtos))

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
			case commandExecutedClientEventName:
				clientEvent, err := jsonutil.Unmarshal[commandExecutedClientEvent](message)
				if err != nil {
					respondServerEvent(generateErroredServerEvent(err))
					return
				}

				if err = httpHandler.executeCommand(worldIdDto, clientEvent.Command); err != nil {
					respondServerEvent(generateErroredServerEvent(err))
					respondAndBroadcastServerEvent(generateCommandFailedServerEvent(clientEvent.Command.Id))
				}
			}

		}
	}()

	closeConnFlag.Wait()
}

func (httpHandler *HttpHandler) executeCommand(
	worldIdDto uuid.UUID,
	commandDto dto.CommandDto,
) error {
	uow := pguow.NewUow()
	executeCommandUseCase := usecase.ProvideExecuteCommandUseCase(uow)
	if err := executeCommandUseCase.Execute(worldIdDto, commandDto); err != nil {
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

	return newPlayerDto, nil
}

func (httpHandler *HttpHandler) getWorldInformation(worldIdDto uuid.UUID) (
	worldDto dto.WorldDto, playerDtos []dto.PlayerDto, err error,
) {
	uow := pguow.NewDummyUow()

	getWorldInformationUseCase := usecase.ProvideGetWorldInformationUseCase(uow)
	return getWorldInformationUseCase.Execute(worldIdDto)
}
