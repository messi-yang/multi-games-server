package worldjourneyhandler

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/messaging/redisservermessagemediator"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/iam/application/service/userappsrv"
	iam_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/iam/infrastructure/providedependency"
	world_dto "github.com/dum-dum-genius/zossi-server/pkg/context/world/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/playerappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/unitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/worldappsrv"
	world_provide_dependency "github.com/dum-dum-genius/zossi-server/pkg/context/world/infrastructure/providedependency"
	"github.com/dum-dum-genius/zossi-server/pkg/interface/http/httpsession"
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
		sendMessage(displayerErrorCommand{
			Id:        uuid.New(),
			Timestamp: time.Now().UnixMilli(),
			Name:      displayerErrorCommandName,
			Message:   err.Error(),
		})
	}
	closeConnectionOnError := func(err error) {
		sendError(err)
		closeConnection()
	}

	var worldIdDto uuid.UUID
	var playerIdDto uuid.UUID
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)

	if worldIdDto, err = uuid.Parse(c.Request.URL.Query().Get("id")); err != nil {
		closeConnectionOnError(err)
		return
	}

	worldServerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newWorldServerMessageChannel(worldIdDto),
		func(serverMessageBytes []byte) {
			command, err := jsonutil.Unmarshal[command](serverMessageBytes)
			if err != nil {
				return
			}

			switch command.Name {
			case createStaticUnitCommandName:
				command, err := jsonutil.Unmarshal[createStaticUnitCommand](serverMessageBytes)
				if err != nil {
					return
				}
				sendMessage(createStaticUnitCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					ItemId:    command.ItemId,
					Position:  command.Position,
					Direction: command.Direction,
				})
			case createPortalUnitCommandName:
				command, err := jsonutil.Unmarshal[createPortalUnitCommand](serverMessageBytes)
				if err != nil {
					return
				}
				sendMessage(createPortalUnitCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					ItemId:    command.ItemId,
					Position:  command.Position,
					Direction: command.Direction,
				})
			case rotateUnitCommandName:
				command, err := jsonutil.Unmarshal[rotateUnitCommand](serverMessageBytes)
				if err != nil {
					return
				}
				sendMessage(rotateUnitCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					Position:  command.Position,
				})
			case removeUnitCommandName:
				command, err := jsonutil.Unmarshal[removeUnitCommand](serverMessageBytes)
				if err != nil {
					return
				}
				sendMessage(removeUnitCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					Position:  command.Position,
				})
			case addPlayerCommandName:
				command, err := jsonutil.Unmarshal[addPlayerCommand](serverMessageBytes)
				if err != nil {
					return
				}
				if command.Player.Id == playerIdDto {
					return
				}
				sendMessage(addPlayerCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					Player:    command.Player,
				})
			case movePlayerCommandName:
				command, err := jsonutil.Unmarshal[movePlayerCommand](serverMessageBytes)
				if err != nil {
					return
				}
				sendMessage(movePlayerCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					PlayerId:  command.PlayerId,
					Position:  command.Position,
					Direction: command.Direction,
				})
			case changePlayerHeldItemCommandName:
				command, err := jsonutil.Unmarshal[changePlayerHeldItemCommand](serverMessageBytes)
				if err != nil {
					return
				}
				sendMessage(changePlayerHeldItemCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					PlayerId:  command.PlayerId,
					ItemId:    command.ItemId,
				})
			case removePlayerCommandName:
				command, err := jsonutil.Unmarshal[removePlayerCommand](serverMessageBytes)
				if err != nil {
					return
				}
				sendMessage(removePlayerCommand{
					Id:        command.Id,
					Timestamp: command.Timestamp,
					Name:      command.Name,
					PlayerId:  command.PlayerId,
				})
			default:
			}
		},
	)
	defer worldServerMessageUnusbscriber()

	playerIdDto, err = httpHandler.executeEnterWorldCommand(worldIdDto, authorizedUserIdDto)
	if err != nil {
		closeConnectionOnError(err)
		return
	}
	safelyLeaveWorldInAllCases := func() {
		if err = httpHandler.executeLeaveWorldCommand(worldIdDto, playerIdDto); err != nil {
			fmt.Println(err)
		}
		httpHandler.redisServerMessageMediator.Send(
			newWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(removePlayerCommand{
				Id:        uuid.New(),
				Timestamp: time.Now().UnixMilli(),
				Name:      removePlayerCommandName,
				PlayerId:  playerIdDto,
			}),
		)
	}
	defer safelyLeaveWorldInAllCases()

	if err = httpHandler.sendAddWorldCommandResponse(worldIdDto, playerIdDto, sendMessage); err != nil {
		closeConnectionOnError(err)
		return
	}

	playerDto, err := httpHandler.queryPlayer(worldIdDto, playerIdDto)
	httpHandler.redisServerMessageMediator.Send(
		newWorldServerMessageChannel(worldIdDto),
		jsonutil.Marshal(addPlayerCommand{
			Id:        uuid.New(),
			Timestamp: time.Now().UnixMilli(),
			Name:      addPlayerCommandName,
			Player:    playerDto,
		}),
	)

	go func() {
		for {
			_, message, err := socketConn.ReadMessage()
			if err != nil {
				closeConnectionOnError(err)
				return
			}

			command, err := jsonutil.Unmarshal[command](message)
			if err != nil {
				closeConnectionOnError(err)
				return
			}

			switch command.Name {
			case pingCommandName:
				continue
			case movePlayerCommandName:
				commandDto, err := jsonutil.Unmarshal[movePlayerCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if commandDto.PlayerId != playerIdDto {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeMovePlayerCommand(worldIdDto, commandDto.PlayerId, commandDto.Position, commandDto.Direction); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(movePlayerCommand{
						Id:        commandDto.Id,
						Timestamp: commandDto.Timestamp,
						Name:      commandDto.Name,
						PlayerId:  commandDto.PlayerId,
						Position:  commandDto.Position,
						Direction: commandDto.Direction,
					}),
				)
			case changePlayerHeldItemCommandName:
				commandDto, err := jsonutil.Unmarshal[changePlayerHeldItemCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if commandDto.PlayerId != playerIdDto {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeChangePlayerHeldItemCommand(worldIdDto, commandDto.PlayerId, commandDto.ItemId); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(changePlayerHeldItemCommand{
						Id:        commandDto.Id,
						Timestamp: commandDto.Timestamp,
						Name:      commandDto.Name,
						PlayerId:  commandDto.PlayerId,
						ItemId:    commandDto.ItemId,
					}),
				)
			case createStaticUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[createStaticUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeCreateStaticUnitCommand(
					worldIdDto,
					commandDto.ItemId,
					commandDto.Position,
					commandDto.Direction,
				); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(createStaticUnitCommand{
						Id:        commandDto.Id,
						Timestamp: commandDto.Timestamp,
						Name:      commandDto.Name,
						ItemId:    commandDto.ItemId,
						Position:  commandDto.Position,
						Direction: commandDto.Direction,
					}),
				)
			case createPortalUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[createPortalUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeCreatePortalUnitCommand(
					worldIdDto,
					commandDto.ItemId,
					commandDto.Position,
					commandDto.Direction,
				); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(createPortalUnitCommand{
						Id:        commandDto.Id,
						Timestamp: commandDto.Timestamp,
						Name:      commandDto.Name,
						ItemId:    commandDto.ItemId,
						Position:  commandDto.Position,
						Direction: commandDto.Direction,
					}),
				)
			case rotateUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[rotateUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRotateUnitCommand(worldIdDto, commandDto.Position); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(rotateUnitCommand{
						Id:        commandDto.Id,
						Timestamp: commandDto.Timestamp,
						Name:      commandDto.Name,
						Position:  commandDto.Position,
					}),
				)
			case removeUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[removeUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRemoveUnitCommand(worldIdDto, commandDto.Position); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(removeUnitCommand{
						Id:        commandDto.Id,
						Timestamp: commandDto.Timestamp,
						Name:      commandDto.Name,
						Position:  commandDto.Position,
					}),
				)
			default:
			}
		}
	}()

	closeConnFlag.Wait()
}

func (httpHandler *HttpHandler) queryPlayer(worldIdDto uuid.UUID, playerIdDto uuid.UUID) (playerDto world_dto.PlayerDto, err error) {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	return playerAppService.GetPlayer(playerappsrv.GetPlayerQuery{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	})
}

func (httpHandler *HttpHandler) executeMovePlayerCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, positionDto world_dto.PositionDto, directionDto int8) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	if err := playerAppService.MovePlayer(playerappsrv.MovePlayerCommand{
		WorldId:   worldIdDto,
		PlayerId:  playerIdDto,
		Position:  positionDto,
		Direction: directionDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) executeChangePlayerHeldItemCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID, itemIdDto uuid.UUID) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	if err := playerAppService.ChangePlayerHeldItem(playerappsrv.ChangePlayerHeldItemCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		ItemId:   itemIdDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) executeCreateStaticUnitCommand(
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.CreateStaticUnit(unitappsrv.CreateStaticUnitCommand{
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

func (httpHandler *HttpHandler) executeCreatePortalUnitCommand(
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.CreatePortalUnit(unitappsrv.CreatePortalUnitCommand{
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

func (httpHandler *HttpHandler) executeRotateUnitCommand(worldIdDto uuid.UUID, positionDto world_dto.PositionDto) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.RotateUnit(unitappsrv.RotateUnitCommand{
		WorldId:  worldIdDto,
		Position: positionDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveUnitCommand(worldIdDto uuid.UUID, positionDto world_dto.PositionDto) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
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

func (httpHandler *HttpHandler) executeEnterWorldCommand(worldIdDto uuid.UUID, userIdDto *uuid.UUID) (playerIdDto uuid.UUID, err error) {
	uow := pguow.NewUow()

	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	itemAppService := world_provide_dependency.ProvideItemAppService(uow)
	userAppService := iam_provide_dependency.ProvideUserAppService(uow)

	itemDtos, err := itemAppService.QueryItems(itemappsrv.QueryItemsQuery{})
	if err != nil {
		uow.RevertChanges()
		return playerIdDto, err
	}

	playerName := "Guest"
	if userIdDto != nil {
		user, err := userAppService.GetUser(userappsrv.GetUserQuery{
			UserId: *userIdDto,
		})
		if err != nil {
			uow.RevertChanges()
			return playerIdDto, err
		}
		playerName = user.FriendlyName
	}

	if playerIdDto, err = playerAppService.EnterWorld(playerappsrv.EnterWorldCommand{
		WorldId:          worldIdDto,
		PlayerName:       playerName,
		PlayerHeldItemId: itemDtos[0].Id,
	}); err != nil {
		uow.RevertChanges()
		return playerIdDto, err
	}

	uow.SaveChanges()

	return playerIdDto, nil
}

func (httpHandler *HttpHandler) executeLeaveWorldCommand(worldIdDto uuid.UUID, playerIdDto uuid.UUID) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	if err := playerAppService.LeaveWorld(playerappsrv.LeaveWorldCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) sendAddWorldCommandResponse(worldIdDto uuid.UUID, playerIdDto uuid.UUID, sendMessage func(any)) error {
	uow := pguow.NewDummyUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	worldAppService := world_provide_dependency.ProvideWorldAppService(uow)
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	userAppService := iam_provide_dependency.ProvideUserAppService(uow)

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{
		WorldId: worldIdDto,
	})
	if err != nil {
		return err
	}

	userDto, err := userAppService.GetUser(userappsrv.GetUserQuery{UserId: worldDto.UserId})
	if err != nil {
		return err
	}

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

	sendMessage(enterWorldCommand{
		Id:         uuid.New(),
		Timestamp:  time.Now().UnixMilli(),
		Name:       enterWorldCommandName,
		World:      viewmodel.WorldViewModel{WorldDto: worldDto, UserDto: userDto},
		Units:      unitDtos,
		MyPlayerId: playerIdDto,
		Players:    playerDtos,
	})
	return nil
}
