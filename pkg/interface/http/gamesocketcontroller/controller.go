package gamesocketcontroller

import (
	"errors"
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
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/infrastructure/messaging/redispubsub"
	"github.com/dum-dum-genius/game-of-liberty-computer/pkg/interface/http/httpdto"
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
		players, err := gameAppService.GetNearbyPlayers(
			gameappservice.GetNearbyPlayersQuery{
				WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
				PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
			},
		)
		if err != nil {
			return err
		}
		myPlayer, myPlayrFound := lo.Find(players, func(player playermodel.PlayerAgg) bool {
			return player.GetId().Uuid() == playerIdDto
		})
		if !myPlayrFound {
			return errors.New("my player not found in players")
		}
		otherPlayers := lo.Filter(players, func(player playermodel.PlayerAgg, _ int) bool {
			return player.GetId().Uuid() != playerIdDto
		})
		sendMessage(playersUpdatedResponseDto{
			Type:     playersUpdatedResponseDtoType,
			MyPlayer: httpdto.NewPlayerAggDto(myPlayer),
			OtherPlayers: lo.Map(otherPlayers, func(player playermodel.PlayerAgg, _ int) httpdto.PlayerAggDto {
				return httpdto.NewPlayerAggDto(player)
			}),
		})
		return nil
	}

	doGetNearbyUnitsQuery := func() error {
		visionBound, units, err := gameAppService.GetNearbyUnits(
			gameappservice.GetNearbyUnitsQuery{
				WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
				PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
			},
		)
		if err != nil {
			return err
		}
		sendMessage(unitsUpdatedResponseDto{
			Type:        unitsUpdatedResponseDtoType,
			VisionBound: httpdto.NewBoundVoDto(visionBound),
			Units: lo.Map(units, func(unit unitmodel.UnitAgg, _ int) httpdto.UnitVoDto {
				return httpdto.NewUnitVoDto(unit)
			}),
		})
		return nil
	}

	doAddPlayerCommand := func() error {
		err := gameAppService.AddPlayer(gameappservice.AddPlayerCommand{
			WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
			PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
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

	if err = doAddPlayerCommand(); err != nil {
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

				isVisionBoundUpdated, err := gameAppService.MovePlayer(gameappservice.MovePlayerCommand{
					WorldId:   worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId:  playermodel.NewPlayerIdVo(playerIdDto),
					Direction: commonmodel.NewDirectionVo(requestDto.Direction),
				})
				if err != nil {
					return
				}
				if isVisionBoundUpdated {
					if err = doGetNearbyUnitsQuery(); err != nil {
						return
					}
				}
				if err = publishPlayersUpdatedEvent(); err != nil {
					return
				}
			case placeItemRequestDtoType:
				requestDto, err := jsonutil.Unmarshal[placeItemRequestDto](message)
				if err != nil {
					return
				}

				if err = gameAppService.PlaceItem(gameappservice.PlaceItemCommand{
					WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
					ItemId:   itemmodel.NewItemIdVo(requestDto.ItemId),
				}); err != nil {
					return
				}
				if err = publishUnitsUpdatedEvent(); err != nil {
					return
				}
			case destroyItemRequestDtoType:
				_, err := jsonutil.Unmarshal[destroyItemRequestDto](message)
				if err != nil {
					return
				}

				if err = gameAppService.DestroyItem(gameappservice.DestroyItemCommand{
					WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
					PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
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

	if err = gameAppService.RemovePlayer(gameappservice.RemovePlayerCommand{
		WorldId:  worldmodel.NewWorldIdVo(worldIdDto),
		PlayerId: playermodel.NewPlayerIdVo(playerIdDto),
	}); err != nil {
		fmt.Println(err)
		return
	}
	if err = publishPlayersUpdatedEvent(); err != nil {
		fmt.Println(err)
		return
	}
}
