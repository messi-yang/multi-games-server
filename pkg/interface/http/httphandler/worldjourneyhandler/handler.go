package worldjourneyhandler

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	iam_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/playerappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/unitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/viewmodel"
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
				serverMessage, err := jsonutil.Unmarshal[unitCreatedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendUnitCreatedResponse(serverMessage.Unit, sendMessage)
			case unitDeletedServerMessageName:
				serverMessage, err := jsonutil.Unmarshal[unitDeletedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendUnitDeletedResponse(serverMessage.Position, sendMessage)
			case playerJoinedServerMessageName:
				serverMessage, err := jsonutil.Unmarshal[playerJoinedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendPlayerJoinedResponse(serverMessage.Player, sendMessage)
			case playerMovedServerMessageName:
				serverMessage, err := jsonutil.Unmarshal[playerMovedServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendPlayerMovedResponse(serverMessage.Player, sendMessage)
			case playerLeftServerMessageName:
				serverMessage, err := jsonutil.Unmarshal[playerLeftServerMessage](serverMessageBytes)
				if err != nil {
					return
				}
				httpHandler.sendPlayerLeftResponse(serverMessage.PlayerId, sendMessage)
			default:
			}
		},
	)
	defer worldServerMessageUnusbscriber()

	playerIdDto, err = httpHandler.executeEnterWorldCommand(worldIdDto)
	if err != nil {
		closeConnectionOnError(err)
		return
	}
	safelyLeaveWorldInAllCases := func() {
		if err = httpHandler.executeLeaveWorldCommand(worldIdDto, playerIdDto); err != nil {
			fmt.Println(err)
		}
		if err = httpHandler.broadcastPlayerLeftServerMessage(worldIdDto, playerIdDto); err != nil {
			closeConnectionOnError(err)
			return
		}
	}
	defer safelyLeaveWorldInAllCases()

	if err = httpHandler.sendWorldEnteredResponse(worldIdDto, playerIdDto, sendMessage); err != nil {
		closeConnectionOnError(err)
		return
	}

	if err = httpHandler.broadcastPlayerJoinedServerMessage(worldIdDto, playerIdDto); err != nil {
		closeConnectionOnError(err)
		return
	}

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
				if err = httpHandler.broadcastPlayerMovedServerMessage(worldIdDto, playerIdDto); err != nil {
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
				if err = httpHandler.broadcastPlayerMovedServerMessage(worldIdDto, playerIdDto); err != nil {
					sendError(err)
				}
			case createUnitRequestType:
				requestDto, err := jsonutil.Unmarshal[createUnitRequest](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeCreateUnitCommand(
					worldIdDto,
					requestDto.ItemId,
					requestDto.Position,
					requestDto.Direction,
				); err != nil {
					sendError(err)
				}
				if err = httpHandler.broadcastUnitCreatedServerMessage(worldIdDto, requestDto.Position); err != nil {
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
				if err = httpHandler.broadcastUnitDeletedServerMessage(worldIdDto, requestDto.Position); err != nil {
					sendError(err)
				}
			default:
			}
		}
	}()

	closeConnFlag.Wait()
}

func (httpHandler *HttpHandler) broadcastUnitCreatedServerMessage(worldIdDto uuid.UUID, positionDto dto.PositionDto) error {
	uow := pguow.NewDummyUow()

	unitAppService := providedependency.ProvideUnitAppService(uow)
	unitDto, err := unitAppService.GetUnit(unitappsrv.GetUnitQuery{
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

func (httpHandler *HttpHandler) broadcastUnitDeletedServerMessage(worldIdDto uuid.UUID, positionDto dto.PositionDto) error {
	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(newunitDeletedServerMessage(worldIdDto, positionDto)),
	)

	return nil
}

func (httpHandler *HttpHandler) broadcastPlayerJoinedServerMessage(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	playerAppService := providedependency.ProvidePlayerAppService()
	playerDto, err := playerAppService.GetPlayer(playerappsrv.GetPlayerQuery{
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

func (httpHandler *HttpHandler) broadcastPlayerMovedServerMessage(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	playerAppService := providedependency.ProvidePlayerAppService()
	playerDto, err := playerAppService.GetPlayer(playerappsrv.GetPlayerQuery{
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

func (httpHandler *HttpHandler) broadcastPlayerLeftServerMessage(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(newplayerLeftServerMessage(playerIdDto)),
	)
	return nil
}

func (httpHandler *HttpHandler) executeMoveCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, directionDto int8) error {
	playerAppService := providedependency.ProvidePlayerAppService()
	if err := playerAppService.Move(playerappsrv.MoveCommand{
		WorldId:   worldIdDto,
		PlayerId:  playerIdDto,
		Direction: directionDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) executeChangeHeldItemCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, itemIdDto uuid.UUID) error {
	playerAppService := providedependency.ProvidePlayerAppService()
	if err := playerAppService.ChangeHeldItem(playerappsrv.ChangeHeldItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		ItemId:   itemIdDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) executeCreateUnitCommand(
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	unitAppService := providedependency.ProvideUnitAppService(uow)
	if err := unitAppService.CreateUnit(unitappsrv.CreateUnitCommand{
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

	unitAppService := providedependency.ProvideUnitAppService(uow)
	if err := unitAppService.RemoveUnit(unitappsrv.RemoveUnitCommand{
		WorldId:  worldIdDto,
		Position: positionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeEnterWorldCommand(worldIdDto uuid.UUID) (playerIdDto uuid.UUID, err error) {
	uow := pguow.NewUow()

	playerAppService := providedependency.ProvidePlayerAppService()
	itemAppService := providedependency.ProvideItemAppService(uow)

	itemDtos, err := itemAppService.QueryItems(itemappsrv.QueryItemsQuery{})
	if err != nil {
		uow.RevertChanges()
		return playerIdDto, err
	}

	if playerIdDto, err = playerAppService.EnterWorld(playerappsrv.EnterWorldCommand{
		WorldId:          worldIdDto,
		PlayerHeldItemId: itemDtos[0].Id,
	}); err != nil {
		uow.RevertChanges()
		return playerIdDto, err
	}
	uow.SaveChanges()

	return playerIdDto, nil
}

func (httpHandler *HttpHandler) executeLeaveWorldCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	playerAppService := providedependency.ProvidePlayerAppService()
	if err := playerAppService.LeaveWorld(playerappsrv.LeaveWorldCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) sendWorldEnteredResponse(worldIdDto uuid.UUID, playerIdDto uuid.UUID, sendMessage func(any)) error {
	uow := pguow.NewDummyUow()

	unitAppService := providedependency.ProvideUnitAppService(uow)
	worldAppService := providedependency.ProvideWorldAppService(uow)
	playerAppService := providedependency.ProvidePlayerAppService()
	userAppService := iam_provide_dependency.ProvideUserAppService(uow)

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		return err
	}

	userDto, err := userAppService.GetUser(userappsrv.GetUserQuery{UserId: worldDto.UserId})

	unitDtos, err := unitAppService.GetUnits(
		unitappsrv.GetUnitsQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return err
	}

	playerDtos, err := playerAppService.GetPlayers(
		playerappsrv.GetPlayersQuery{
			WorldId:  worldIdDto,
			PlayerId: playerIdDto,
		},
	)
	if err != nil {
		return err
	}

	sendMessage(worldEnteredResponse{
		Type:       worldEnteredResponseType,
		World:      viewmodel.WorldViewModel{WorldDto: worldDto, UserDto: userDto},
		Units:      unitDtos,
		MyPlayerId: playerIdDto,
		Players:    playerDtos,
	})
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
