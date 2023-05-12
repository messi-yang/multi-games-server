package gamesockethandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/gziputil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/jsonutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
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

func (httpHandler *HttpHandler) GameConnection(c *gin.Context) {
	socketConn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusBadRequest, "failed to upgrade http to socket")
		return
	}
	defer socketConn.Close()

	sendMessageLocker := &sync.RWMutex{}
	sendMessage := func(jsonObj any) {
		sendMessageLocker.Lock()
		defer sendMessageLocker.Unlock()

		messageJsonInBytes := jsonutil.Marshal(jsonObj)
		compressedMessage, _ := gziputil.Gzip(messageJsonInBytes)

		err = socketConn.WriteMessage(2, compressedMessage)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	var closeConnFlag sync.WaitGroup
	closeConnFlag.Add(1)
	closeConnection := func() {
		closeConnFlag.Done()
	}
	closeConnectionOnError := func(err error) {
		sendMessage(errorHappenedResponse{
			Type:    errorHappenedResponseType,
			Message: err.Error(),
		})
		closeConnection()
	}

	var worldIdDto uuid.UUID
	var playerIdDto uuid.UUID

	if worldIdDto, err = uuid.Parse(c.Request.URL.Query().Get("id")); err != nil {
		closeConnectionOnError(err)
		return
	}

	if playerIdDto, err = httpHandler.executeEnterWorldCommand(worldIdDto); err != nil {
		closeConnectionOnError(err)
		return
	}
	defer func() {
		if err = httpHandler.executeLeaveWorldCommand(worldIdDto, playerIdDto); err != nil {
			fmt.Println(err)
		}
	}()

	if err = httpHandler.executeGetNearbyUnitsQuery(worldIdDto, playerIdDto, sendMessage); err != nil {
		closeConnectionOnError(err)
		return
	}

	if err = httpHandler.executeGetNearbyPlayersQuery(worldIdDto, playerIdDto, sendMessage); err != nil {
		closeConnectionOnError(err)
		return
	}

	moveSteps := 0
	worldServerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		gameappsrv.NewWorldServerMessageChannel(worldIdDto),
		func(serverMessageBytes []byte) {
			serverMessage, err := jsonutil.Unmarshal[gameappsrv.ServerMessage](serverMessageBytes)
			if err != nil {
				return
			}

			switch serverMessage.Name {
			case gameappsrv.UnitCreated:
				httpHandler.executeGetNearbyUnitsQuery(worldIdDto, playerIdDto, sendMessage)
			case gameappsrv.UnitDeleted:
				httpHandler.executeGetNearbyUnitsQuery(worldIdDto, playerIdDto, sendMessage)
			case gameappsrv.PlayerJoined:
				httpHandler.executeGetNearbyPlayersQuery(worldIdDto, playerIdDto, sendMessage)
			case gameappsrv.PlayerMoved:
				httpHandler.executeGetNearbyPlayersQuery(worldIdDto, playerIdDto, sendMessage)
				moveSteps += 1
				if moveSteps%10 == 0 {
					httpHandler.executeGetNearbyUnitsQuery(worldIdDto, playerIdDto, sendMessage)
				}
			case gameappsrv.PlayerLeft:
				httpHandler.executeGetNearbyPlayersQuery(worldIdDto, playerIdDto, sendMessage)
			default:
			}
		},
	)
	defer worldServerMessageUnusbscriber()

	go func() {
		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				closeConnectionOnError(err)
				return
			}

			message, err := gziputil.Ungzip(compressedMessage)
			if err != nil {
				closeConnectionOnError(err)
				return
			}

			genericRequest, err := jsonutil.Unmarshal[genericRequest](message)
			if err != nil {
				closeConnectionOnError(err)
				return
			}

			switch genericRequest.Type {
			case pingRequestType:
				continue
			case moveRequestType:
				requestDto, err := jsonutil.Unmarshal[moveRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeMoveCommand(worldIdDto, playerIdDto, requestDto.Direction); err != nil {
					closeConnectionOnError(err)
				}
			case changeHeldItemRequestType:
				requestDto, err := jsonutil.Unmarshal[changeHeldItemRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeChangeHeldItemCommand(worldIdDto, playerIdDto, requestDto.ItemId); err != nil {
					closeConnectionOnError(err)
				}
			case placeItemRequestType:
				if _, err := jsonutil.Unmarshal[placeItemRequest](message); err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executePlaceItemCommand(worldIdDto, playerIdDto); err != nil {
					closeConnectionOnError(err)
				}
			case removeItemRequestType:
				_, err := jsonutil.Unmarshal[removeItemRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRemoveItemCommand(worldIdDto, playerIdDto); err != nil {
					closeConnectionOnError(err)
				}
			default:
			}
		}
	}()

	closeConnFlag.Wait()
}

func (httpHandler *HttpHandler) executeMoveCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, directionDto int8) error {
	pgUow := pguow.NewUow()

	gameAppService := provideGameAppService(pgUow)
	if err := gameAppService.Move(gameappsrv.MoveCommand{
		WorldId:   worldIdDto,
		PlayerId:  playerIdDto,
		Direction: directionDto,
	}); err != nil {
		pgUow.RevertChanges()
		return err
	}
	pgUow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeChangeHeldItemCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, itemIdDto uuid.UUID) error {
	pgUow := pguow.NewUow()

	gameAppService := provideGameAppService(pgUow)
	if err := gameAppService.ChangeHeldItem(gameappsrv.ChangeHeldItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		ItemId:   itemIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return err
	}
	pgUow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executePlaceItemCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	pgUow := pguow.NewUow()

	gameAppService := provideGameAppService(pgUow)
	if err := gameAppService.PlaceItem(gameappsrv.PlaceItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return err
	}
	pgUow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveItemCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	pgUow := pguow.NewUow()

	gameAppService := provideGameAppService(pgUow)
	if err := gameAppService.RemoveItem(gameappsrv.RemoveItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return err
	}
	pgUow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeEnterWorldCommand(worldIdDto uuid.UUID) (playerIdDto uuid.UUID, err error) {
	pgUow := pguow.NewUow()

	gameAppService := provideGameAppService(pgUow)
	if playerIdDto, err = gameAppService.EnterWorld(gameappsrv.EnterWorldCommand{
		WorldId: worldIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return playerIdDto, err
	}
	pgUow.SaveChanges()
	return playerIdDto, nil
}

func (httpHandler *HttpHandler) executeLeaveWorldCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	pgUow := pguow.NewUow()

	gameAppService := provideGameAppService(pgUow)
	if err := gameAppService.LeaveWorld(gameappsrv.LeaveWorldCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return err
	}
	pgUow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeGetNearbyUnitsQuery(worldIdDto uuid.UUID, playerIdDto uuid.UUID, sendMessage func(any)) error {
	pgUow := pguow.NewDummyUow()

	gameAppService := provideGameAppService(pgUow)
	unitDtos, err := gameAppService.GetNearbyUnits(
		gameappsrv.GetNearbyUnitsQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return err
	}

	sendMessage(unitsUpdatedResponse{
		Type:  unitsUpdatedResponseType,
		Units: unitDtos,
	})
	return nil
}

func (httpHandler *HttpHandler) executeGetNearbyPlayersQuery(worldIdDto uuid.UUID, playerIdDto uuid.UUID, sendMessage func(any)) error {
	pgUow := pguow.NewDummyUow()

	gameAppService := provideGameAppService(pgUow)
	myPlayerDto, otherPlayerDtos, err := gameAppService.GetNearbyPlayers(
		gameappsrv.GetNearbyPlayersQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return err
	}

	sendMessage(playersUpdatedResponse{
		Type:         playersUpdatedResponseType,
		MyPlayer:     myPlayerDto,
		OtherPlayers: otherPlayerDtos,
	})
	return nil
}
