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
	sendError := func(err error) {
		sendMessage(errorHappenedResponse{
			Type:    errorHappenedResponseType,
			Message: err.Error(),
		})
	}
	closeConnectionOnError := func(err error) {
		sendError(err)
		closeConnection()
	}

	var worldIdDto uuid.UUID
	var playerIdDto uuid.UUID

	if worldIdDto, err = uuid.Parse(c.Request.URL.Query().Get("id")); err != nil {
		closeConnectionOnError(err)
		return
	}

	worldServerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newWorldServerMessageChannel(worldIdDto),
		func(serverMessageBytes []byte) {
			serverMessage, err := jsonutil.Unmarshal[ServerMessage](serverMessageBytes)
			if err != nil {
				return
			}

			switch serverMessage.Name {
			case unitCreatedServerMessageName:
				unitCreatedServerMessage, err := jsonutil.Unmarshal[unitCreatedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendunitCreatedServerMessageNameResponse(unitCreatedServerMessage.Unit, sendMessage)
			case unitDeletedServerMessageName:
				unitDeletedServerMessage, err := jsonutil.Unmarshal[unitDeletedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendunitDeletedServerMessageNameResponse(unitDeletedServerMessage.Position, sendMessage)
			case playerJoinedServerMessageName:
				playerJoinedServerMessage, err := jsonutil.Unmarshal[playerJoinedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendplayerJoinedServerMessageNameResponse(playerJoinedServerMessage.Player, sendMessage)
			case playerMovedServerMessageName:
				playerMovedServerMessage, err := jsonutil.Unmarshal[playerMovedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendplayerMovedServerMessageNameResponse(playerMovedServerMessage.Player, sendMessage)
			case playerLeftServerMessageName:
				playerLeftServerMessage, err := jsonutil.Unmarshal[playerLeftServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendplayerLeftServerMessageNameResponse(playerLeftServerMessage.PlayerId, sendMessage)
			default:
			}
		},
	)
	defer worldServerMessageUnusbscriber()

	if playerIdDto, err = httpHandler.executeEnterWorldCommand(worldIdDto, sendMessage); err != nil {
		closeConnectionOnError(err)
		return
	}
	if err = httpHandler.broadcastplayerJoinedServerMessage(worldIdDto, playerIdDto); err != nil {
		closeConnectionOnError(err)
		return
	}
	defer func() {
		if err = httpHandler.executeLeaveWorldCommand(worldIdDto, playerIdDto); err != nil {
			fmt.Println(err)
		}
		if err = httpHandler.broadcastplayerLeftServerMessage(worldIdDto, playerIdDto); err != nil {
			closeConnectionOnError(err)
			return
		}
	}()

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
					sendError(err)
				}
				if err = httpHandler.broadcastplayerMovedServerMessage(worldIdDto, playerIdDto); err != nil {
					sendError(err)
				}
			case changeHeldItemRequestType:
				requestDto, err := jsonutil.Unmarshal[changeHeldItemRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeChangeHeldItemCommand(worldIdDto, playerIdDto, requestDto.ItemId); err != nil {
					sendError(err)
				}
			case placeUnitRequestType:
				requestDto, err := jsonutil.Unmarshal[placeUnitRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executePlaceUnitCommand(
					worldIdDto,
					requestDto.ItemId,
					requestDto.Position,
					requestDto.Direction,
				); err != nil {
					sendError(err)
				}
				if err = httpHandler.broadcastunitCreatedServerMessage(worldIdDto, requestDto.Position); err != nil {
					sendError(err)
				}
			case removeUnitRequestType:
				requestDto, err := jsonutil.Unmarshal[removeUnitRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRemoveUnitCommand(worldIdDto, requestDto.Position); err != nil {
					sendError(err)
				}
				if err = httpHandler.broadcastunitDeletedServerMessage(worldIdDto, requestDto.Position); err != nil {
					sendError(err)
				}
			default:
			}
		}
	}()

	closeConnFlag.Wait()
}

func (httpHandler *HttpHandler) broadcastunitCreatedServerMessage(worldIdDto uuid.UUID, positionDto dto.PositionDto) error {
	uow := pguow.NewDummyUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	unitDto, err := worldJourneyAppService.GetUnit(worldjourneyappsrv.GetUnitQuery{
		WorldId:  worldIdDto,
		Position: positionDto,
	})
	if err != nil {
		return err
	}

	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(newunitCreatedServerMessage(unitDto)),
	)

	return nil
}

func (httpHandler *HttpHandler) broadcastunitDeletedServerMessage(worldIdDto uuid.UUID, positionDto dto.PositionDto) error {
	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(newunitDeletedServerMessage(worldIdDto, positionDto)),
	)

	return nil
}

func (httpHandler *HttpHandler) broadcastplayerJoinedServerMessage(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	uow := pguow.NewDummyUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	playerDto, err := worldJourneyAppService.GetPlayer(worldjourneyappsrv.GetPlayerQuery{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	})
	if err != nil {
		return err
	}
	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(newplayerJoinedServerMessage(playerDto)),
	)
	return nil
}

func (httpHandler *HttpHandler) broadcastplayerMovedServerMessage(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	uow := pguow.NewDummyUow()
	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	playerDto, err := worldJourneyAppService.GetPlayer(worldjourneyappsrv.GetPlayerQuery{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	})
	if err != nil {
		return err
	}
	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(newplayerMovedServerMessage(playerDto)),
	)
	return nil
}

func (httpHandler *HttpHandler) broadcastplayerLeftServerMessage(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(newplayerLeftServerMessage(playerIdDto)),
	)
	return nil
}

func (httpHandler *HttpHandler) executeMoveCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, directionDto int8) error {
	uow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	if err := worldJourneyAppService.Move(worldjourneyappsrv.MoveCommand{
		WorldId:   worldIdDto,
		PlayerId:  playerIdDto,
		Direction: directionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeChangeHeldItemCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, itemIdDto uuid.UUID) error {
	uow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	if err := worldJourneyAppService.ChangeHeldItem(worldjourneyappsrv.ChangeHeldItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		ItemId:   itemIdDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executePlaceUnitCommand(
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	if err := worldJourneyAppService.PlaceUnit(worldjourneyappsrv.PlaceUnitCommand{
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveUnitCommand(worldIdDto uuid.UUID, positionDto dto.PositionDto) error {
	uow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	if err := worldJourneyAppService.RemoveUnit(worldjourneyappsrv.RemoveUnitCommand{
		WorldId:  worldIdDto,
		Position: positionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeEnterWorldCommand(worldIdDto uuid.UUID, sendMessage func(any)) (playerIdDto uuid.UUID, err error) {
	uow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	worldAppService := providedependency.ProvideWorldAppService(uow)

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		return playerIdDto, err
	}

	if playerIdDto, err = worldJourneyAppService.EnterWorld(worldjourneyappsrv.EnterWorldCommand{
		WorldId: worldIdDto,
	}); err != nil {
		uow.RevertChanges()
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

	uow.SaveChanges()

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
	uow := pguow.NewUow()

	worldJourneyAppService := providedependency.ProvideWorldJourneyAppService(uow)
	if err := worldJourneyAppService.LeaveWorld(worldjourneyappsrv.LeaveWorldCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) sendunitCreatedServerMessageNameResponse(unitDto dto.UnitDto, sendMessage func(any)) error {
	sendMessage(unitCreatedResponse{
		Type: unitCreatedResponseType,
		Unit: unitDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendunitDeletedServerMessageNameResponse(positionDto dto.PositionDto, sendMessage func(any)) error {
	sendMessage(unitDeletedResponse{
		Type:     unitDeletedResponseType,
		Position: positionDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendplayerJoinedServerMessageNameResponse(playerDto dto.PlayerDto, sendMessage func(any)) error {
	sendMessage(playerJoinedResponse{
		Type:   playerJoinedResponseType,
		Player: playerDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendplayerLeftServerMessageNameResponse(playerIdDto uuid.UUID, sendMessage func(any)) error {
	sendMessage(playerLeftResponse{
		Type:     playerLeftResponseType,
		PlayerId: playerIdDto,
	})
	return nil
}

func (httpHandler *HttpHandler) sendplayerMovedServerMessageNameResponse(playerDto dto.PlayerDto, sendMessage func(any)) error {
	sendMessage(playerMovedResponse{
		Type:   playerMovedResponseType,
		Player: playerDto,
	})
	return nil
}
