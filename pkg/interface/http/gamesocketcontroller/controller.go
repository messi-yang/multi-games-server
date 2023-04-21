package gamesocketcontroller

import (
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/game/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/context/sharedkernel/infrastructure/redispubsub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/gziputil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/util/jsonutil"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func gameConnectionHandler(c *gin.Context) {
	socketConn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
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
	gameAppService, err := provideGameAppService()
	if err != nil {
		return
	}

	closeConnChan := make(chan bool)
	disconnectOnError := func(err error) {
		closeConnChan <- true
	}

	sendMessageLocker := &sync.RWMutex{}
	sendMessage := func(jsonObj any) {
		sendMessageLocker.Lock()
		defer sendMessageLocker.Unlock()

		messageJsonInBytes := jsonutil.Marshal(jsonObj)
		compressedMessage, _ := gziputil.Gzip(messageJsonInBytes)

		err = socketConn.WriteMessage(2, compressedMessage)
		if err != nil {
			disconnectOnError(err)
		}
	}

	publishPlayersUpdatedEvent := func() {
		err = redisChannelPublisher.Publish(
			newPlayersUpdatedMessageChannel(worldIdDto),
			PlayersUpdatedMessage{},
		)
		if err != nil {
			disconnectOnError(err)
		}
	}

	publishUnitsUpdatedEvent := func() {
		err = redisChannelPublisher.Publish(
			NewUnitsUpdatedMessageChannel(worldIdDto),
			UnitsUpdatedMessage{},
		)
		if err != nil {
			disconnectOnError(err)
		}
	}

	doGetNearbyPlayersQuery := func() {
		myPlayerDto, otherPlayerDtos, err := gameAppService.GetNearbyPlayers(
			gameappservice.GetNearbyPlayersQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			},
		)
		if err != nil {
			disconnectOnError(err)
			return
		}
		sendMessage(playersUpdatedResponseDto{
			Type:         playersUpdatedResponseDtoType,
			MyPlayer:     myPlayerDto,
			OtherPlayers: otherPlayerDtos,
		})
	}

	doGetNearbyUnitsQuery := func() {
		unitDtos, err := gameAppService.GetNearbyUnits(
			gameappservice.GetNearbyUnitsQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			},
		)
		if err != nil {
			disconnectOnError(err)
			return
		}
		sendMessage(unitsUpdatedResponseDto{
			Type:  unitsUpdatedResponseDtoType,
			Units: unitDtos,
		})
	}

	doEnterWorldCommand := func() {
		err := gameAppService.EnterWorld(gameappservice.EnterWorldCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		})
		if err != nil {
			disconnectOnError(err)
			return
		}

		sendMessage(gameJoinedResponseDto{
			Type: gameJoinedResponseDtoType,
		})
	}

	doMoveCommand := func(directionDto int8) {
		playerDto, err := gameAppService.GetPlayer(gameappservice.GetPlayerQuery{PlayerId: playerIdDto})
		if err != nil {
			disconnectOnError(err)
			return
		}
		if err := gameAppService.Move(gameappservice.MoveCommand{
			WorldId:   worldIdDto,
			PlayerId:  playerIdDto,
			Direction: directionDto,
		}); err != nil {
			disconnectOnError(err)
			return
		}
		updatedPlayerDto, err := gameAppService.GetPlayer(gameappservice.GetPlayerQuery{PlayerId: playerIdDto})
		if err != nil {
			disconnectOnError(err)
			return
		}
		playerVisionBoundUpdated := playerDto.VisionBound != updatedPlayerDto.VisionBound
		if playerVisionBoundUpdated {
			doGetNearbyUnitsQuery()
		}
	}

	doChangeHeldItemCommand := func(itemIdDto uuid.UUID) {
		if err = gameAppService.ChangeHeldItem(gameappservice.ChangeHeldItemCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
			ItemId:   itemIdDto,
		}); err != nil {
			disconnectOnError(err)
		}
	}

	doPlaceItemCommand := func() {
		if err = gameAppService.PlaceItem(gameappservice.PlaceItemCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		}); err != nil {
			disconnectOnError(err)
		}
	}

	doRemoveItemCommand := func() {
		if err = gameAppService.RemoveItem(gameappservice.RemoveItemCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		}); err != nil {
			disconnectOnError(err)
		}
	}

	doLeaveWorldCommand := func() {
		if err = gameAppService.LeaveWorld(gameappservice.LeaveWorldCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		}); err != nil {
			disconnectOnError(err)
		}
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

	doEnterWorldCommand()
	doGetNearbyUnitsQuery()
	publishPlayersUpdatedEvent()

	go func() {
		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				return
			}

			message, err := gziputil.Ungzip(compressedMessage)
			if err != nil {
				continue
			}

			genericRequestDto, err := jsonutil.Unmarshal[genericRequestDto](message)
			if err != nil {
				continue
			}

			switch genericRequestDto.Type {
			case pingRequestDtoType:
				continue
			case moveRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[moveRequestDto](message)
				if err != nil {
					disconnectOnError(err)
					return
				}
				doMoveCommand(requestDto.Direction)
				publishPlayersUpdatedEvent()
			case changeHeldItemRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[changeHeldItemRequestDto](message)
				if err != nil {
					disconnectOnError(err)
					return
				}
				doChangeHeldItemCommand(requestDto.ItemId)
				publishPlayersUpdatedEvent()
			case placeItemRequestDtoType:
				if _, err := jsonutil.Unmarshal[placeItemRequestDto](message); err != nil {
					disconnectOnError(err)
					return
				}
				doPlaceItemCommand()
				publishUnitsUpdatedEvent()
			case removeItemRequestDtoType:
				_, err := jsonutil.Unmarshal[removeItemRequestDto](message)
				if err != nil {
					disconnectOnError(err)
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
