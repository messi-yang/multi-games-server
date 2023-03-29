package gamesocketcontroller

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/application/service/gameappservice"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/gziputil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/common/util/jsonutil"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/commonmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/itemmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/playermodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/unitmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/domain/model/worldmodel"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httpdto"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/messaging/redisinteventsubscriber"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/samber/lo"
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
			players, err := gameAppService.FindNearbyPlayers(
				gameappservice.FindNearbyPlayersQuery{
					WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
				},
			)
			if err != nil {
				disconnect()
				return
			}
			sendMessage(playersUpdatedResponseDto{
				Type: playersUpdatedResponseDtoType,
				Players: lo.Map(players, func(player playermodel.PlayerAgg, _ int) httpdto.PlayerAggDto {
					return httpdto.NewPlayerAggDto(player)
				}),
			})
		},
	)
	defer playersUpdatedIntEventUnsubscriber()

	unitsUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameappservice.UnitsUpdatedIntEvent]().Subscribe(
		gameappservice.NewUnitsUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.UnitsUpdatedIntEvent) {
			visionBound, units, err := gameAppService.QueryUnits(
				gameappservice.QueryUnitsQuery{
					WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
				},
			)
			if err != nil {
				disconnect()
				return
			}
			sendMessage(unitsUpdatedResponseDto{
				Type:        unitsUpdatedResponseDtoType,
				VisionBound: httpdto.NewBoundVoDto(visionBound),
				Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) httpdto.UnitVoDto {
					return httpdto.NewUnitVoDto(unit)
				}),
			})
		},
	)
	defer unitsUpdatedIntEventTypeUnsubscriber()

	visionBoundUpdatedIntEventTypeUnsubscriber := redisinteventsubscriber.New[gameappservice.VisionBoundUpdatedIntEvent]().Subscribe(
		gameappservice.NewVisionBoundUpdatedIntEventChannel(worldIdDto, playerIdDto),
		func(intEvent gameappservice.VisionBoundUpdatedIntEvent) {
			visionBound, units, err := gameAppService.QueryUnits(
				gameappservice.QueryUnitsQuery{
					WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
				},
			)
			if err != nil {
				disconnect()
				return
			}
			sendMessage(unitsUpdatedResponseDto{
				Type:        unitsUpdatedResponseDtoType,
				VisionBound: httpdto.NewBoundVoDto(visionBound),
				Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) httpdto.UnitVoDto {
					return httpdto.NewUnitVoDto(unit)
				}),
			})
		},
	)
	defer visionBoundUpdatedIntEventTypeUnsubscriber()

	{
		items, players, visionBound, units, err := gameAppService.AddPlayer(gameappservice.AddPlayerCommand{
			WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
			PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
		})
		if err != nil {
			disconnect()
			return
		}

		sendMessage(gameJoinedResponseDto{
			Type: gameJoinedResponseDtoType,
			Items: lo.Map(items, func(item itemmodel.ItemAgg, _ int) httpdto.ItemAggDto {
				return httpdto.NewItemAggDto(item)
			}),
			PlayerId: playerIdDto,
			Players: lo.Map(players, func(p playermodel.PlayerAgg, _ int) httpdto.PlayerAggDto {
				return httpdto.NewPlayerAggDto(p)
			}),
			VisionBound: httpdto.NewBoundVoDto(visionBound),
			Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) httpdto.UnitVoDto {
				return httpdto.NewUnitVoDto(unit)
			}),
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
					WorldId:   worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId:  playermodel.NewPlayerIdVo(playerIdDto),
					Direction: commonmodel.NewDirectionVo(requestDto.Direction),
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
					WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
					ItemId:   itemmodel.NewItemIdVo(requestDto.ItemId),
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
					WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
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
			WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
			PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
