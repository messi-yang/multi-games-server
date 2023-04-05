package gamesocketcontroller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/gziputil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/jsonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/messaging/redispubsub"
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
	disconnect := func() {
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
			disconnect()
		}
	}

	publishPlayersUpdatedEvent := func() error {
		return redisChannelPublisher.Publish(
			newPlayersUpdatedMessageChannel(worldIdDto),
			PlayersUpdatedMessage{},
		)
	}

	publishUnitsUpdatedEvent := func() error {
		return redisChannelPublisher.Publish(
			NewUnitsUpdatedMessageChannel(worldIdDto),
			UnitsUpdatedMessage{},
		)
	}

	doGetNearbyPlayersQuery := func() error {
		myPlayerDto, otherPlayerDtos, err := gameAppService.GetNearbyPlayers(
			gameappservice.GetNearbyPlayersQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			},
		)
		if err != nil {
			return err
		}
		sendMessage(playersUpdatedResponseDto{
			Type:         playersUpdatedResponseDtoType,
			MyPlayer:     myPlayerDto,
			OtherPlayers: otherPlayerDtos,
		})
		return nil
	}

	doGetNearbyUnitsQuery := func() error {
		visionBoundDto, unitDtos, err := gameAppService.GetNearbyUnits(
			gameappservice.GetNearbyUnitsQuery{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			},
		)
		if err != nil {
			return err
		}
		sendMessage(unitsUpdatedResponseDto{
			Type:        unitsUpdatedResponseDtoType,
			VisionBound: visionBoundDto,
			Units:       unitDtos,
		})
		return nil
	}

	doEnterWorldCommand := func() error {
		err := gameAppService.EnterWorld(gameappservice.EnterWorldCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		})
		if err != nil {
			return err
		}

		sendMessage(gameJoinedResponseDto{
			Type: gameJoinedResponseDtoType,
		})
		return nil
	}

	playersUpdatedMessageUnsubscriber := redispubsub.NewChannelSubscriber[PlayersUpdatedMessage]().Subscribe(
		newPlayersUpdatedMessageChannel(worldIdDto),
		func(message PlayersUpdatedMessage) {
			if err = doGetNearbyPlayersQuery(); err != nil {
				disconnect()
			}
		},
	)
	defer playersUpdatedMessageUnsubscriber()

	unitsUpdatedMessageTypeUnsubscriber := redispubsub.NewChannelSubscriber[UnitsUpdatedMessage]().Subscribe(
		NewUnitsUpdatedMessageChannel(worldIdDto),
		func(message UnitsUpdatedMessage) {
			if err = doGetNearbyUnitsQuery(); err != nil {
				disconnect()
			}
		},
	)
	defer unitsUpdatedMessageTypeUnsubscriber()

	if err = doEnterWorldCommand(); err != nil {
		disconnect()
		return
	}
	if err = publishPlayersUpdatedEvent(); err != nil {
		disconnect()
		return
	}
	if err = doGetNearbyUnitsQuery(); err != nil {
		disconnect()
		return
	}

	go func() {
		defer disconnect()

		for {
			_, compressedMessage, err := socketConn.ReadMessage()
			if err != nil {
				return
			}

			message, err := gziputil.Ungzip(compressedMessage)
			if err != nil {
				return
			}

			genericRequestDto, err := jsonutil.Unmarshal[genericRequestDto](message)
			if err != nil {
				return
			}

			switch genericRequestDto.Type {
			case pingRequestDtoType:
				continue
			case moveRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[moveRequestDto](message)
				if err != nil {
					return
				}

				playerDto, err := gameAppService.GetPlayer(gameappservice.GetPlayerQuery{PlayerId: playerIdDto})
				if err != nil {
					return
				}
				if err := gameAppService.Move(gameappservice.MoveCommand{
					WorldId:   worldIdDto,
					PlayerId:  playerIdDto,
					Direction: requestDto.Direction,
				}); err != nil {
					return
				}
				updatedPlayerDto, err := gameAppService.GetPlayer(gameappservice.GetPlayerQuery{PlayerId: playerIdDto})
				if err != nil {
					return
				}
				playerVisionBoundUpdated := playerDto.VisionBound != updatedPlayerDto.VisionBound
				if playerVisionBoundUpdated {
					if err = doGetNearbyUnitsQuery(); err != nil {
						return
					}
				}
				if err = publishPlayersUpdatedEvent(); err != nil {
					return
				}
			case changeHeldItemRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[changeHeldItemRequestDto](message)
				if err != nil {
					return
				}

				if err = gameAppService.ChangeHeldItem(gameappservice.ChangeHeldItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
					ItemId:   requestDto.ItemId,
				}); err != nil {
					return
				}
				if err = publishPlayersUpdatedEvent(); err != nil {
					return
				}
			case placeItemRequestDtoType:
				if _, err := jsonutil.Unmarshal[placeItemRequestDto](message); err != nil {
					return
				}

				if err = gameAppService.PlaceItem(gameappservice.PlaceItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				}); err != nil {
					return
				}
				if err = publishUnitsUpdatedEvent(); err != nil {
					return
				}
			case removeItemRequestDtoType:
				_, err := jsonutil.Unmarshal[removeItemRequestDto](message)
				if err != nil {
					return
				}

				if err = gameAppService.RemoveItem(gameappservice.RemoveItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				}); err != nil {
					return
				}
				if err = publishUnitsUpdatedEvent(); err != nil {
					return
				}
			default:
			}
		}
	}()

	<-closeConnChan

	if err = gameAppService.LeaveWorld(gameappservice.LeaveWorldCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		fmt.Println(err)
		return
	}
	if err = publishPlayersUpdatedEvent(); err != nil {
		fmt.Println(err)
		return
	}
}
