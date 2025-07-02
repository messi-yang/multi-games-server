package roomservicehandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/application/usecase"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httpsession"
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

func (httpHandler *HttpHandler) StartService(c *gin.Context) {
	socketConn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to upgrade http to socket")
		return
	}
	defer socketConn.Close()

	guestPlayerName := c.Request.URL.Query().Get("name")
	if guestPlayerName == "" {
		guestPlayerName = "Guest"
	}

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

	var roomIdDto uuid.UUID
	var myPlayerIdDto uuid.UUID

	respondServerEvent := func(serverEvent any) {
		respondMessage(serverEvent)
	}

	broadcastServerEvent := func(serverEvent any) {
		httpHandler.redisServerMessageMediator.Send(
			newRoomMessageChannel(roomIdDto),
			jsonutil.Marshal(roomMessage[any]{
				SenderId:    myPlayerIdDto,
				ServerEvent: serverEvent,
			}),
		)
	}

	respondAndBroadcastServerEvent := func(serverEvent any) {
		respondServerEvent(serverEvent)
		broadcastServerEvent(serverEvent)
	}

	if roomIdDto, err = uuid.Parse(c.Request.URL.Query().Get("id")); err != nil {
		respondServerEvent(httpHandler.generateErroredServerEvent(err))
		closeConnection()
		return
	}

	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)
	myPlayerDto, err := httpHandler.createPlayer(roomIdDto, authorizedUserIdDto, guestPlayerName)
	if err != nil {
		respondServerEvent(httpHandler.generateErroredServerEvent(err))
		closeConnection()
		return
	}

	myPlayerIdDto = myPlayerDto.Id
	broadcastServerEvent(httpHandler.generatePlayerJoinedServerEvent(myPlayerDto))

	safelyLeaveRoomInAllCases := func() {
		if err = httpHandler.removePlayer(roomIdDto, myPlayerIdDto); err != nil {
			fmt.Println(err)
		}
		broadcastServerEvent(httpHandler.generatePlayerLeftServerEvent(myPlayerIdDto))

	}
	defer safelyLeaveRoomInAllCases()

	roomDto, gameDto, commandDtos, playerDtos, err := httpHandler.getRoomInformation(roomIdDto)
	if err != nil {
		respondServerEvent(httpHandler.generateErroredServerEvent(err))
		closeConnection()
		return
	}
	respondServerEvent(httpHandler.generateRoomJoinedServerEvent(roomDto, gameDto, commandDtos, myPlayerIdDto, playerDtos))

	// Subscribe to messages sent specifically to your player
	playerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newPlayerMessageChannel(roomIdDto, myPlayerIdDto),
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
			} else if message.ServerEvent.Name == commandReceivedServerEventName {
				commandReceivedServerMessage, err := jsonutil.Unmarshal[playerMessage[commandReceivedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(commandReceivedServerMessage.ServerEvent)
			}
		},
	)
	defer playerMessageUnusbscriber()

	// Subscribe to messages broadcasted from other websocket servers
	roomMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newRoomMessageChannel(roomIdDto),
		func(messageBytes []byte) {
			message, err := jsonutil.Unmarshal[roomMessage[serverEvent]](messageBytes)
			if err != nil {
				return
			}
			if message.SenderId == myPlayerIdDto {
				return
			}

			if message.ServerEvent.Name == roomJoinedServerEventName {
				roomJoinedServerMessage, err := jsonutil.Unmarshal[roomMessage[roomJoinedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(roomJoinedServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == gameStartedServerEventName {
				gameStartedServerMessage, err := jsonutil.Unmarshal[roomMessage[gameStartedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(gameStartedServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == newGameSetupServerEventName {
				newGameSetupServerMessage, err := jsonutil.Unmarshal[roomMessage[newGameSetupServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(newGameSetupServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == playerJoinedServerEventName {
				playerJoinedServerEvent, err := jsonutil.Unmarshal[roomMessage[playerJoinedServerEvent]](messageBytes)
				if err != nil {
					return
				}

				respondMessage(playerJoinedServerEvent.ServerEvent)
			} else if message.ServerEvent.Name == playerLeftServerEventName {
				playerLeftServerMessage, err := jsonutil.Unmarshal[roomMessage[playerLeftServerEvent]](messageBytes)
				if err != nil {
					return
				}

				respondMessage(playerLeftServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == commandReceivedServerEventName {
				commandReceivedServerMessage, err := jsonutil.Unmarshal[roomMessage[commandReceivedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(commandReceivedServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == commandFailedServerEventName {
				commandFailedServerMessage, err := jsonutil.Unmarshal[roomMessage[commandFailedServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(commandFailedServerMessage.ServerEvent)
			} else if message.ServerEvent.Name == erroredServerEventName {
				erroredServerMessage, err := jsonutil.Unmarshal[roomMessage[erroredServerEvent]](messageBytes)
				if err != nil {
					return
				}
				respondMessage(erroredServerMessage.ServerEvent)
			}
		},
	)
	defer roomMessageUnusbscriber()

	go func() {
		for {
			_, message, err := socketConn.ReadMessage()
			if err != nil {
				respondServerEvent(httpHandler.generateErroredServerEvent(err))
				closeConnection()
				return
			}

			clientEvent, err := jsonutil.Unmarshal[clientEvent](message)
			if err != nil {
				respondServerEvent(httpHandler.generateErroredServerEvent(err))
				return
			}

			switch clientEvent.Name {
			case pingClientEventName:
				continue
			case startGameRequestedClientEventName:
				clientEvent, err := jsonutil.Unmarshal[startGameRequestedClientEvent](message)
				if err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				if gameDto, err = httpHandler.startGame(roomIdDto, clientEvent.GameId, clientEvent.GameState); err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				respondAndBroadcastServerEvent(httpHandler.generateGameStartedServerEvent(gameDto))
			case setupNewGameRequestedClientEventName:
				clientEvent, err := jsonutil.Unmarshal[setupNewGameRequestedClientEvent](message)
				if err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				if gameDto, err = httpHandler.setupNewGame(roomIdDto, clientEvent.GameName); err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				respondAndBroadcastServerEvent(httpHandler.generateNewGameSetupServerEvent(gameDto))
			case p2pOfferSentClientEventName:
				clientEvent, err := jsonutil.Unmarshal[p2pOfferSentClientEvent](message)
				if err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				httpHandler.sendServerEventToPlayer(roomIdDto, clientEvent.PeerPlayerId, httpHandler.generateP2pOfferReceivedServerEvent(myPlayerIdDto, clientEvent.IceCandidates, clientEvent.Offer))
			case p2pAnswerSentClientEventName:
				clientEvent, err := jsonutil.Unmarshal[p2pAnswerSentClientEvent](message)
				if err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				httpHandler.sendServerEventToPlayer(roomIdDto, clientEvent.PeerPlayerId, httpHandler.generateP2pAnswerReceivedServerEvent(myPlayerIdDto, clientEvent.IceCandidates, clientEvent.Answer))
			case p2pConnectedClientEventName:
				_, err := jsonutil.Unmarshal[p2pConnectedClientEvent](message)
				if err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				// TODO: Must have something that we can do with it
			case commandSentClientEventName:
				clientEvent, err := jsonutil.Unmarshal[commandSentClientEvent](message)
				if err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}
				httpHandler.sendServerEventToPlayer(roomIdDto, clientEvent.PeerPlayerId, httpHandler.generateCommandReceivedServerEvent(clientEvent.Command))
			case commandExecutedClientEventName:
				clientEvent, err := jsonutil.Unmarshal[commandExecutedClientEvent](message)
				if err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					return
				}

				if err = httpHandler.saveCommand(clientEvent.Command); err != nil {
					respondServerEvent(httpHandler.generateErroredServerEvent(err))
					respondAndBroadcastServerEvent(httpHandler.generateCommandFailedServerEvent(clientEvent.Command.Id))
				}
			}

		}
	}()

	closeConnFlag.Wait()
}

func (httpHandler *HttpHandler) saveCommand(
	commandDto dto.CommandDto,
) error {
	uow := pguow.NewUow()
	saveCommandUseCase := usecase.ProvideSaveCommandUseCase(uow)
	if err := saveCommandUseCase.Execute(commandDto); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) removePlayer(roomIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	uow := pguow.NewUow()
	removePlayerUseCase := usecase.ProvideRemovePlayerUseCase(uow)
	if err := removePlayerUseCase.Execute(roomIdDto, playerIdDto); err != nil {
		uow.RevertChanges()
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) createPlayer(roomIdDto uuid.UUID, userIdDto *uuid.UUID, guestPlayerName string) (newPlayerDto dto.PlayerDto, err error) {
	uow := pguow.NewUow()

	createPlayerUseCase := usecase.ProvideCreatePlayerUseCase(uow)
	newPlayerDto, err = createPlayerUseCase.Execute(roomIdDto, userIdDto, guestPlayerName)
	if err != nil {
		uow.RevertChanges()
		return newPlayerDto, err
	}
	uow.SaveChanges()

	return newPlayerDto, nil
}

func (httpHandler *HttpHandler) getRoomPlayers(roomIdDto uuid.UUID) (
	playerDtos []dto.PlayerDto, err error,
) {
	uow := pguow.NewDummyUow()
	getRoomPlayersUseCase := usecase.ProvideGetRoomPlayersUseCase(uow)
	return getRoomPlayersUseCase.Execute(roomIdDto)
}

func (httpHandler *HttpHandler) getRoomInformation(roomIdDto uuid.UUID) (
	roomDto dto.RoomDto, gameDto dto.GameDto, commandDtos []dto.CommandDto, playerDtos []dto.PlayerDto, err error,
) {
	uow := pguow.NewDummyUow()

	getRoomInformationUseCase := usecase.ProvideGetRoomInformationUseCase(uow)
	return getRoomInformationUseCase.Execute(roomIdDto)
}

func (httpHandler *HttpHandler) startGame(roomIdDto uuid.UUID, gameId uuid.UUID, gameState map[string]interface{}) (gameDto dto.GameDto, err error) {
	uow := pguow.NewDummyUow()
	startGameUseCase := usecase.ProvideStartGameUseCase(uow)
	return startGameUseCase.Execute(roomIdDto, gameId, gameState)
}

func (httpHandler *HttpHandler) setupNewGame(roomId uuid.UUID, gameName string) (gameDto dto.GameDto, err error) {
	uow := pguow.NewDummyUow()
	setupNewGameUseCase := usecase.ProvideSetupNewGameUseCase(uow)
	return setupNewGameUseCase.Execute(roomId, gameName)
}

func (httpHandler *HttpHandler) sendServerEventToPlayer(roomIdDto uuid.UUID, playerIdDto uuid.UUID, serverEvent any) {
	httpHandler.redisServerMessageMediator.Send(
		newPlayerMessageChannel(roomIdDto, playerIdDto),
		jsonutil.Marshal(playerMessage[any]{
			ServerEvent: serverEvent,
		}),
	)
}

func (httpHandler *HttpHandler) generatePlayerJoinedServerEvent(playerDto dto.PlayerDto) playerJoinedServerEvent {
	return playerJoinedServerEvent{
		Name:   playerJoinedServerEventName,
		Player: playerDto,
	}
}

func (httpHandler *HttpHandler) generateRoomJoinedServerEvent(roomDto dto.RoomDto, gameDto dto.GameDto, commandDtos []dto.CommandDto, myPlayerIdDto uuid.UUID, playerDtos []dto.PlayerDto) roomJoinedServerEvent {
	return roomJoinedServerEvent{
		Name:       roomJoinedServerEventName,
		Game:       gameDto,
		Commands:   commandDtos,
		Room:       roomDto,
		MyPlayerId: myPlayerIdDto,
		Players:    playerDtos,
	}
}

func (httpHandler *HttpHandler) generatePlayerLeftServerEvent(playerIdDto uuid.UUID) playerLeftServerEvent {
	return playerLeftServerEvent{
		Name:     playerLeftServerEventName,
		PlayerId: playerIdDto,
	}
}

func (httpHandler *HttpHandler) generateCommandReceivedServerEvent(command dto.CommandDto) commandReceivedServerEvent {
	return commandReceivedServerEvent{
		Name:    commandReceivedServerEventName,
		Command: command,
	}
}

func (httpHandler *HttpHandler) generateCommandFailedServerEvent(commandId uuid.UUID) commandFailedServerEvent {
	return commandFailedServerEvent{
		Name:      commandFailedServerEventName,
		CommandId: commandId,
	}
}

func (httpHandler *HttpHandler) generateGameStartedServerEvent(gameDto dto.GameDto) gameStartedServerEvent {
	return gameStartedServerEvent{
		Name: gameStartedServerEventName,
		Game: gameDto,
	}
}

func (httpHandler *HttpHandler) generateNewGameSetupServerEvent(gameDto dto.GameDto) newGameSetupServerEvent {
	return newGameSetupServerEvent{
		Name: newGameSetupServerEventName,
		Game: gameDto,
	}
}

func (httpHandler *HttpHandler) generateP2pOfferReceivedServerEvent(peerPlayerId uuid.UUID, iceCandidates []any, offer any) p2pOfferReceivedServerEvent {
	return p2pOfferReceivedServerEvent{
		Name:          p2pOfferReceivedServerEventName,
		PeerPlayerId:  peerPlayerId,
		IceCandidates: iceCandidates,
		Offer:         offer,
	}
}

func (httpHandler *HttpHandler) generateP2pAnswerReceivedServerEvent(peerPlayerId uuid.UUID, iceCandidates []any, offer any) p2pAnswerReceivedServerEvent {
	return p2pAnswerReceivedServerEvent{
		Name:          p2pAnswerReceivedServerEventName,
		PeerPlayerId:  peerPlayerId,
		IceCandidates: iceCandidates,
		Answer:        offer,
	}
}

func (httpHandler *HttpHandler) generateErroredServerEvent(err error) erroredServerEvent {
	return erroredServerEvent{
		Name:    erroredServerEventName,
		Message: err.Error(),
	}
}
