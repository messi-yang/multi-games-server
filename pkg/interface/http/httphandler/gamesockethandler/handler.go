package gamesockethandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappsrv"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/messaging/redispubsub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/unitofwork/pguow"
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

type HttpHandler struct{}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}

func (httpHandler *HttpHandler) GameConnection(c *gin.Context) {
	socketConn, err := websocketUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer socketConn.Close()

	worldIdDto, err := uuid.Parse(c.Request.URL.Query().Get("id"))
	if err != nil {
		return
	}
	playerIdDto := uuid.New()

	redisChannelPublisher := redispubsub.NewChannelPublisher()

	sendMessageLocker := &sync.RWMutex{}
	sendMessage := func(jsonObj any) {
		sendMessageLocker.Lock()
		defer sendMessageLocker.Unlock()

		messageJsonInBytes := jsonutil.Marshal(jsonObj)
		compressedMessage, _ := gziputil.Gzip(messageJsonInBytes)

		err = socketConn.WriteMessage(2, compressedMessage)
		if err != nil {
			fmt.Println("send error")
		}
	}

	closeConnChan := make(chan bool, 100)
	closeConnectionOnError := func(err error) {
		sendMessage(errorHappenedResponse{
			Type:    errorHappenedResponseType,
			Message: err.Error(),
		})
		closeConnChan <- true
	}

	publishPlayersUpdatedEvent := func() {
		err = redisChannelPublisher.Publish(
			newPlayersUpdatedMessageChannel(worldIdDto),
			PlayersUpdatedMessage{},
		)
		if err != nil {
			closeConnectionOnError(err)
		}
	}

	publishUnitsUpdatedEvent := func() {
		err = redisChannelPublisher.Publish(
			NewUnitsUpdatedMessageChannel(worldIdDto),
			UnitsUpdatedMessage{},
		)
		if err != nil {
			closeConnectionOnError(err)
		}
	}

	doGetNearbyPlayersQuery := func() {
		pgUow := pguow.NewDummyUow()

		gameAppService := provideGameAppService(pgUow)
		myPlayerDto, otherPlayerDtos, err := gameAppService.GetNearbyPlayers(
			gameappsrv.GetNearbyPlayersQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			},
		)
		if err != nil {
			closeConnectionOnError(err)
			return
		}

		sendMessage(playersUpdatedResponse{
			Type:         playersUpdatedResponseType,
			MyPlayer:     myPlayerDto,
			OtherPlayers: otherPlayerDtos,
		})
	}

	doGetNearbyUnitsQuery := func() {
		pgUow := pguow.NewDummyUow()

		gameAppService := provideGameAppService(pgUow)
		unitDtos, err := gameAppService.GetNearbyUnits(
			gameappsrv.GetNearbyUnitsQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			},
		)
		if err != nil {
			closeConnectionOnError(err)
			return
		}

		sendMessage(unitsUpdatedResponse{
			Type:  unitsUpdatedResponseType,
			Units: unitDtos,
		})
	}

	doEnterWorldCommand := func() {
		pgUow := pguow.NewUow()

		gameAppService := provideGameAppService(pgUow)
		if err = gameAppService.EnterWorld(gameappsrv.EnterWorldCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		}); err != nil {
			pgUow.Rollback()
			closeConnectionOnError(err)
			return
		}
		pgUow.Commit()

		sendMessage(gameJoinedResponse{
			Type: gameJoinedResponseType,
		})
	}

	moveSteps := 0
	doMoveCommand := func(directionDto int8) {
		pgUow := pguow.NewUow()

		gameAppService := provideGameAppService(pgUow)
		if err := gameAppService.Move(gameappsrv.MoveCommand{
			WorldId:   worldIdDto,
			PlayerId:  playerIdDto,
			Direction: directionDto,
		}); err != nil {
			pgUow.Rollback()
			closeConnectionOnError(err)
			return
		}
		pgUow.Commit()
		moveSteps += 1
		if moveSteps%10 == 0 {
			doGetNearbyUnitsQuery()
		}
	}

	doChangeHeldItemCommand := func(itemIdDto uuid.UUID) {
		pgUow := pguow.NewUow()

		gameAppService := provideGameAppService(pgUow)
		if err = gameAppService.ChangeHeldItem(gameappsrv.ChangeHeldItemCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
			ItemId:   itemIdDto,
		}); err != nil {
			pgUow.Rollback()
			closeConnectionOnError(err)
		}
		pgUow.Commit()
	}

	doPlaceItemCommand := func() {
		pgUow := pguow.NewUow()

		gameAppService := provideGameAppService(pgUow)
		if err = gameAppService.PlaceItem(gameappsrv.PlaceItemCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		}); err != nil {
			pgUow.Rollback()
			closeConnectionOnError(err)
		}
		pgUow.Commit()
	}

	doRemoveItemCommand := func() {
		pgUow := pguow.NewUow()

		gameAppService := provideGameAppService(pgUow)
		if err = gameAppService.RemoveItem(gameappsrv.RemoveItemCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		}); err != nil {
			pgUow.Rollback()
			closeConnectionOnError(err)
		}
		pgUow.Commit()
	}

	doLeaveWorldCommand := func() {
		pgUow := pguow.NewUow()

		gameAppService := provideGameAppService(pgUow)
		if err = gameAppService.LeaveWorld(gameappsrv.LeaveWorldCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		}); err != nil {
			pgUow.Rollback()
			closeConnectionOnError(err)
		}
		pgUow.Commit()
	}

	playersUpdatedMessageUnsubscriber := redispubsub.NewChannelSubscriber[PlayersUpdatedMessage]().Subscribe(
		newPlayersUpdatedMessageChannel(worldIdDto),
		func(message PlayersUpdatedMessage) {
			doGetNearbyPlayersQuery()
		},
	)
	defer playersUpdatedMessageUnsubscriber()

	unitsUpdatedMessageTypeUnsubscriber := redispubsub.NewChannelSubscriber[UnitsUpdatedMessage]().Subscribe(
		NewUnitsUpdatedMessageChannel(worldIdDto),
		func(message UnitsUpdatedMessage) {
			doGetNearbyUnitsQuery()
		},
	)
	defer unitsUpdatedMessageTypeUnsubscriber()

	go func() {
		doEnterWorldCommand()
		doGetNearbyUnitsQuery()
		publishPlayersUpdatedEvent()
	}()

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
				doMoveCommand(requestDto.Direction)
				publishPlayersUpdatedEvent()
			case changeHeldItemRequestType:
				requestDto, err := jsonutil.Unmarshal[changeHeldItemRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				doChangeHeldItemCommand(requestDto.ItemId)
				publishPlayersUpdatedEvent()
			case placeItemRequestType:
				if _, err := jsonutil.Unmarshal[placeItemRequest](message); err != nil {
					closeConnectionOnError(err)
					return
				}
				doPlaceItemCommand()
				publishUnitsUpdatedEvent()
			case removeItemRequestType:
				_, err := jsonutil.Unmarshal[removeItemRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				doRemoveItemCommand()
				publishUnitsUpdatedEvent()
			default:
			}
		}
	}()

	<-closeConnChan

	doLeaveWorldCommand()
	publishPlayersUpdatedEvent()
}
