package gamesocketcontroller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/gziputil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/jsonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/messaging/redisinteventsubscriber"
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

	playersUpdatedIntEventUnsubscriber := redisinteventsubscriber.New[gameappservice.PlayersUpdatedIntEvent]().Subscribe(
		gameappservice.NewPlayersUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.PlayersUpdatedIntEvent) {
			playerDtos, err := gameAppService.GetPlayersAroundPlayer(
				gameappservice.GetPlayersAroundPlayerQuery{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				},
			)
			if err != nil {
				disconnect()
				return
			}
			sendMessage(playersUpdatedResponseDto{
				Type:    playersUpdatedResponseDtoType,
				Players: playerDtos,
			})
		},
	)
	defer playersUpdatedIntEventUnsubscriber()

	unitsUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameappservice.UnitsUpdatedIntEvent]().Subscribe(
		gameappservice.NewUnitsUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.UnitsUpdatedIntEvent) {
			visionBoundDto, unitDtos, err := gameAppService.GetUnitsVisibleByPlayer(
				gameappservice.GetUnitsVisibleByPlayerQuery{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				},
			)
			if err != nil {
				disconnect()
				return
			}
			sendMessage(unitsUpdatedResponseDto{
				Type:        unitsUpdatedResponseDtoType,
				VisionBound: visionBoundDto,
				Units:       unitDtos,
			})
		},
	)
	defer unitsUpdatedIntEventTypeUnsubscriber()

	visionBoundUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameappservice.VisionBoundUpdatedIntEvent]().Subscribe(
		gameappservice.NewVisionBoundUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.VisionBoundUpdatedIntEvent) {
			visionBoundDto, unitDtos, err := gameAppService.GetUnitsVisibleByPlayer(
				gameappservice.GetUnitsVisibleByPlayerQuery{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				},
			)
			if err != nil {
				disconnect()
				return
			}
			sendMessage(unitsUpdatedResponseDto{
				Type:        unitsUpdatedResponseDtoType,
				VisionBound: visionBoundDto,
				Units:       unitDtos,
			})
		},
	)
	defer visionBoundUpdatedIntEventTypeUnsubscriber()

	{
		itemDtos, playerDtos, visionBoundDto, unitDtos, err := gameAppService.AddPlayer(
			gameappservice.AddPlayerCommand{
				WorldId:  worldIdDto,
				PlayerId: playerIdDto,
			},
		)
		if err != nil {
			disconnect()
			return
		}
		sendMessage(gameJoinedResponseDto{
			Type:        gameJoinedResponseDtoType,
			Items:       itemDtos,
			PlayerId:    playerIdDto,
			Players:     playerDtos,
			VisionBound: visionBoundDto,
			Units:       unitDtos,
		})
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

				err = gameAppService.MovePlayer(gameappservice.MovePlayerCommand{
					WorldId:   worldIdDto,
					PlayerId:  playerIdDto,
					Direction: requestDto.Direction,
				})
				if err != nil {
					return
				}
			case placeItemRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[placeItemRequestDto](message)
				if err != nil {
					return
				}

				err = gameAppService.PlaceItem(gameappservice.PlaceItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
					ItemId:   requestDto.ItemId,
				})
				if err != nil {
					return
				}
			case destroyItemRequestDtoType:
				_, err := jsonutil.Unmarshal[destroyItemRequestDto](message)
				if err != nil {
					return
				}

				err = gameAppService.DestroyItem(gameappservice.DestroyItemCommand{
					WorldId:  worldIdDto,
					PlayerId: playerIdDto,
				})
				if err != nil {
					return
				}
			default:
			}
		}
	}()

	for {
		<-closeConnChan

		err = gameAppService.RemovePlayer(gameappservice.RemovePlayerCommand{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
