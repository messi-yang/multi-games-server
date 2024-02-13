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
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/linkunitappsrv"
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

var (
	ErrCommandIsNotExecutedByOwnPlayer = fmt.Errorf("the command is not executed by its own player")
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
		sendMessage(erroredEvent{
			Name:    erroredEventName,
			Message: err.Error(),
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
			command, err := jsonutil.Unmarshal[any](serverMessageBytes)
			if err != nil {
				return
			}
			sendMessage(commandSucceededEvent{
				Name:    commandSucceededEventName,
				Command: command,
			})
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

	if err = httpHandler.sendEnterWorldCommandResponse(worldIdDto, playerIdDto, sendMessage); err != nil {
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
			case changePlayerActionCommandName:
				commandDto, err := jsonutil.Unmarshal[changePlayerActionCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if commandDto.PlayerId != playerIdDto {
					closeConnectionOnError(ErrCommandIsNotExecutedByOwnPlayer)
					return
				}
				if err = httpHandler.executeChangePlayerActionCommand(
					worldIdDto, commandDto.PlayerId, commandDto.Action,
				); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case sendPlayerIntoPortalCommandName:
				commandDto, err := jsonutil.Unmarshal[sendPlayerIntoPortalCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if commandDto.PlayerId != playerIdDto {
					closeConnectionOnError(ErrCommandIsNotExecutedByOwnPlayer)
					return
				}
				if err = httpHandler.executeSendPlayerIntoPortalCommand(worldIdDto, commandDto.PlayerId, commandDto.Position); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case changePlayerHeldItemCommandName:
				commandDto, err := jsonutil.Unmarshal[changePlayerHeldItemCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if commandDto.PlayerId != playerIdDto {
					closeConnectionOnError(ErrCommandIsNotExecutedByOwnPlayer)
					return
				}
				if err = httpHandler.executeChangePlayerHeldItemCommand(worldIdDto, commandDto.PlayerId, commandDto.ItemId); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case createStaticUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[createStaticUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeCreateStaticUnitCommand(
					commandDto.UnitId,
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
					jsonutil.Marshal(commandDto),
				)
			case createFenceUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[createFenceUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeCreateFenceUnitCommand(
					commandDto.UnitId,
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
					jsonutil.Marshal(commandDto),
				)
			case createPortalUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[createPortalUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeCreatePortalUnitCommand(
					commandDto.UnitId,
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
					jsonutil.Marshal(commandDto),
				)
			case createLinkUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[createLinkUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeCreateLinkUnitCommand(
					commandDto.UnitId,
					worldIdDto,
					commandDto.ItemId,
					commandDto.Position,
					commandDto.Direction,
					commandDto.Url,
				); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case rotateUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[rotateUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRotateUnitCommand(commandDto.UnitId); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case removeStaticUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[removeStaticUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRemoveStaticUnitCommand(commandDto.UnitId); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case removeFenceUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[removeFenceUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRemoveFenceUnitCommand(commandDto.UniId); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case removePortalUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[removePortalUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRemovePortalUnitCommand(commandDto.UnitId); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
				)
			case removeLinkUnitCommandName:
				commandDto, err := jsonutil.Unmarshal[removeLinkUnitCommand](message)
				if err != nil {
					closeConnectionOnError(err)
					return
				}
				if err = httpHandler.executeRemoveLinkUnitCommand(commandDto.UnitId); err != nil {
					sendError(err)
					break
				}
				httpHandler.redisServerMessageMediator.Send(
					newWorldServerMessageChannel(worldIdDto),
					jsonutil.Marshal(commandDto),
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

func (httpHandler *HttpHandler) executeChangePlayerActionCommand(
	worldIdDto uuid.UUID, playerIdDto uuid.UUID, actionDto world_dto.PlayerActionDto,
) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	if err := playerAppService.ChangePlayerAction(playerappsrv.ChangePlayerActionCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		Action:   actionDto,
	}); err != nil {
		return err
	}
	return nil
}

func (httpHandler *HttpHandler) executeSendPlayerIntoPortalCommand(
	worldIdDto uuid.UUID,
	playerIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	return playerAppService.SendPlayerIntoPortal(playerappsrv.SendPlayerIntoPortalCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		Position: positionDto,
	})
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
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.CreateStaticUnit(unitappsrv.CreateStaticUnitCommand{
		Id:        idDto,
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

func (httpHandler *HttpHandler) executeRemoveStaticUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.RemoveStaticUnit(unitappsrv.RemoveStaticUnitCommand{
		Id: idDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeCreateFenceUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.CreateFenceUnit(unitappsrv.CreateFenceUnitCommand{
		Id:        idDto,
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

func (httpHandler *HttpHandler) executeRemoveFenceUnitCommand(unidIdDto uuid.UUID) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.RemoveFenceUnit(unitappsrv.RemoveFenceUnitCommand{
		Id: unidIdDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeCreatePortalUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.CreatePortalUnit(unitappsrv.CreatePortalUnitCommand{
		Id:        idDto,
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

func (httpHandler *HttpHandler) executeRemovePortalUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.RemovePortalUnit(unitappsrv.RemovePortalUnitCommand{
		Id: idDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeCreateLinkUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
	url string,
) error {
	uow := pguow.NewUow()

	linkUnitAppService := world_provide_dependency.ProvideLinkUnitAppService(uow)
	if err := linkUnitAppService.CreateLinkUnit(linkunitappsrv.CreateLinkUnitCommand{
		Id:        idDto,
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
		Url:       url,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveLinkUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	linkUnitAppService := world_provide_dependency.ProvideLinkUnitAppService(uow)
	if err := linkUnitAppService.RemoveLinkUnit(linkunitappsrv.RemoveLinkUnitCommand{
		Id: idDto,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRotateUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	if err := unitAppService.RotateUnit(unitappsrv.RotateUnitCommand{
		Id: idDto,
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

func (httpHandler *HttpHandler) sendEnterWorldCommandResponse(worldIdDto uuid.UUID, playerIdDto uuid.UUID, sendMessage func(any)) error {
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

	sendMessage(worldEnteredEvent{
		Name:       worldEnteredEventName,
		World:      viewmodel.WorldViewModel{WorldDto: worldDto, UserDto: userDto},
		Units:      unitDtos,
		MyPlayerId: playerIdDto,
		Players:    playerDtos,
	})
	return nil
}
