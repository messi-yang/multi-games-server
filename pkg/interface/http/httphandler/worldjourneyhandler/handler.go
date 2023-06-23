package worldjourneyhandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/messaging/redis/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/infrastructure/persistence/postgres/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldjourneyappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
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

func (httpHandler *HttpHandler) StartJourney(c *gin.Context) {
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

	if playerIdDto, err = httpHandler.executeEnterWorldCommand(worldIdDto, sendMessage); err != nil {
		closeConnectionOnError(err)
		return
	}
	defer func() {
		if err = httpHandler.executeLeaveWorldCommand(worldIdDto, playerIdDto); err != nil {
			fmt.Println(err)
		}
	}()

	worldServerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		worldjourneyappsrv.NewWorldServerMessageChannel(worldIdDto),
		func(serverMessageBytes []byte) {
			serverMessage, err := jsonutil.Unmarshal[worldjourneyappsrv.ServerMessage](serverMessageBytes)
			if err != nil {
				return
			}

			switch serverMessage.Name {
			case worldjourneyappsrv.UnitCreated:
				unitCreatedServerMessage, err := jsonutil.Unmarshal[worldjourneyappsrv.UnitCreatedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendUnitCreatedResponse(unitCreatedServerMessage.Unit, sendMessage)
			case worldjourneyappsrv.UnitDeleted:
				unitDeletedServerMessage, err := jsonutil.Unmarshal[worldjourneyappsrv.UnitDeletedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendUnitDeletedResponse(unitDeletedServerMessage.Position, sendMessage)
			case worldjourneyappsrv.PlayerJoined:
				playerJoinedServerMessage, err := jsonutil.Unmarshal[worldjourneyappsrv.PlayerJoinedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendPlayerJoinedResponse(playerJoinedServerMessage.Player, sendMessage)
			case worldjourneyappsrv.PlayerMoved:
				playerMovedServerMessage, err := jsonutil.Unmarshal[worldjourneyappsrv.PlayerMovedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendPlayerMovedResponse(playerMovedServerMessage.Player, sendMessage)
			case worldjourneyappsrv.PlayerLeft:
				playerLeftServerMessage, err := jsonutil.Unmarshal[worldjourneyappsrv.PlayerLeftServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendPlayerLeftResponse(playerLeftServerMessage.PlayerId, sendMessage)
			default:
			}
		},
	)
	defer worldServerMessageUnusbscriber()

	go func() {
		for {
			_, message, err := socketConn.ReadMessage()
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

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(pgUow)
	if err := worldJourneyAppService.Move(worldjourneyappsrv.MoveCommand{
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

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(pgUow)
	if err := worldJourneyAppService.ChangeHeldItem(worldjourneyappsrv.ChangeHeldItemCommand{
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

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(pgUow)
	if err := worldJourneyAppService.PlaceItem(worldjourneyappsrv.PlaceItemCommand{
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

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(pgUow)
	if err := worldJourneyAppService.RemoveItem(worldjourneyappsrv.RemoveItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return err
	}
	pgUow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeEnterWorldCommand(worldIdDto uuid.UUID, sendMessage func(any)) (playerIdDto uuid.UUID, err error) {
	pgUow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(pgUow)
	worldAppService := providedependency.ProvideWorldAppService(pgUow)

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		return playerIdDto, err
	}

	if playerIdDto, err = worldJourneyAppService.EnterWorld(worldjourneyappsrv.EnterWorldCommand{
		WorldId: worldIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return playerIdDto, err
	}

	unitDtos, err := worldJourneyAppService.GetUnits(
		worldjourneyappsrv.GetUnitsQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return playerIdDto, err
	}

	playerDtos, err := worldJourneyAppService.GetPlayers(
		worldjourneyappsrv.GetPlayersQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return playerIdDto, err
	}

	pgUow.SaveChanges()

	sendMessage(worldsEnteredResponse{
		Type:       worldEnteredResponseType,
		World:      worldDto,
		Units:      unitDtos,
		MyPlayerId: playerIdDto,
		Players:    playerDtos,
	})
	return playerIdDto, nil
}

func (httpHandler *HttpHandler) executeLeaveWorldCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	pgUow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(pgUow)
	if err := worldJourneyAppService.LeaveWorld(worldjourneyappsrv.LeaveWorldCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		pgUow.RevertChanges()
		return err
	}
	pgUow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) sendUnitCreatedResponse(unitDto dto.UnitDto, sendMessage func(any)) error {
	sendMessage(unitCreatedResponse{
		Type: unitCreatedResponseType,
		Unit: unitDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendUnitDeletedResponse(positionDto dto.PositionDto, sendMessage func(any)) error {
	sendMessage(unitDeletedResponse{
		Type:     unitDeletedResponseType,
		Position: positionDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendPlayerJoinedResponse(playerDto dto.PlayerDto, sendMessage func(any)) error {
	sendMessage(playerJoinedResponse{
		Type:   playerJoinedResponseType,
		Player: playerDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendPlayerLeftResponse(playerIdDto uuid.UUID, sendMessage func(any)) error {
	sendMessage(playerLeftResponse{
		Type:     playerLeftResponseType,
		PlayerId: playerIdDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendPlayerMovedResponse(playerDto dto.PlayerDto, sendMessage func(any)) error {
	sendMessage(playerMovedResponse{
		Type:   playerMovedResponseType,
		Player: playerDto,
	})
	return nil
}
