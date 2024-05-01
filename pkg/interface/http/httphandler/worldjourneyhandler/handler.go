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
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/embedunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/fenceunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/itemappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/linkunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/playerappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/portalunitappsrv"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/application/service/staticunitappsrv"
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
	sendErroredServerEvent := func(err error) {
		sendMessage(erroredServerEvent{
			Name:    erroredServerEventName,
			Message: err.Error(),
		})
	}

	var worldIdDto uuid.UUID
	var playerIdDto uuid.UUID
	authorizedUserIdDto := httpsession.GetAuthorizedUserId(c)

	if worldIdDto, err = uuid.Parse(c.Request.URL.Query().Get("id")); err != nil {
		sendErroredServerEvent(err)
		closeConnection()
		return
	}

	worldServerMessageUnusbscriber := httpHandler.redisServerMessageMediator.Receive(
		newWorldServerMessageChannel(worldIdDto),
		func(serverMessageBytes []byte) {
			serverEvent, err := jsonutil.Unmarshal[any](serverMessageBytes)
			if err != nil {
				return
			}
			sendMessage(serverEvent)
		},
	)
	defer worldServerMessageUnusbscriber()

	broadcastPlayerJoinedServerEvent := func(playerDto world_dto.PlayerDto) {
		httpHandler.redisServerMessageMediator.Send(
			newWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(playerJoinedServerEvent{
				Name:   playerJoinedServerEventName,
				Player: playerDto,
			}),
		)
	}

	broadcastPlayerLeftServerEvent := func(playerIdDto uuid.UUID) {
		httpHandler.redisServerMessageMediator.Send(
			newWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(playerLeftServerEvent{
				Name:     playerLeftServerEventName,
				PlayerId: playerIdDto,
			}),
		)
	}

	broadcastCommandSucceededServerEvent := func(command any) {
		httpHandler.redisServerMessageMediator.Send(
			newWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(commandSucceededServerEvent{
				Name:    commandSucceededServerEventName,
				Command: command,
			}),
		)
	}

	broadcastCommandFailedServerEvent := func(commandId uuid.UUID, err error) {
		httpHandler.redisServerMessageMediator.Send(
			newWorldServerMessageChannel(worldIdDto),
			jsonutil.Marshal(commandFailedServerEvent{
				Name:         commandFailedServerEventName,
				CommandId:    commandId,
				ErrorMessage: err.Error(),
			}),
		)
	}

	playerIdDto, err = httpHandler.enterWorld(worldIdDto, authorizedUserIdDto)
	if err != nil {
		sendErroredServerEvent(err)
		closeConnection()
		return
	}
	safelyLeaveWorldInAllCases := func() {
		if err = httpHandler.executeLeaveWorldCommand(worldIdDto, playerIdDto); err != nil {
			fmt.Println(err)
		}
		broadcastCommandSucceededServerEvent(removePlayerCommand{
			Id:        uuid.New(),
			Timestamp: time.Now().UnixMilli(),
			Name:      removePlayerCommandName,
			PlayerId:  playerIdDto,
		})
		broadcastPlayerLeftServerEvent(playerIdDto)
	}
	defer safelyLeaveWorldInAllCases()

	if err = httpHandler.sendWorldEnteredResponse(worldIdDto, playerIdDto, sendMessage); err != nil {
		sendErroredServerEvent(err)
		closeConnection()
		return
	}

	playerDto, err := httpHandler.queryPlayer(worldIdDto, playerIdDto)
	if err != nil {
		sendErroredServerEvent(err)
		closeConnection()
		return
	}
	broadcastCommandSucceededServerEvent(addPlayerCommand{
		Id:        uuid.New(),
		Timestamp: time.Now().UnixMilli(),
		Name:      addPlayerCommandName,
		Player:    playerDto,
	})
	broadcastPlayerJoinedServerEvent(playerDto)

	go func() {
		for {
			_, message, err := socketConn.ReadMessage()
			if err != nil {
				sendErroredServerEvent(err)
				closeConnection()
				return
			}

			clientEvent, err := jsonutil.Unmarshal[clientEvent](message)
			if err != nil {
				sendErroredServerEvent(err)
				return
			}

			switch clientEvent.Name {
			case pingClientEventName:
				continue
			case commandRequestedClientEventName:
				genericCommandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[command]](message)
				if err != nil {
					sendErroredServerEvent(err)
					return
				}
				switch genericCommandRequestedClientEvent.Command.Name {
				case changePlayerActionCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[changePlayerActionCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if commandDto.PlayerId != playerIdDto {
						sendErroredServerEvent(ErrCommandIsNotExecutedByOwnPlayer)
						return
					}
					if err = httpHandler.executeChangePlayerActionCommand(
						worldIdDto, commandDto.PlayerId, commandDto.Action,
					); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case sendPlayerIntoPortalCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[sendPlayerIntoPortalCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if commandDto.PlayerId != playerIdDto {
						sendErroredServerEvent(ErrCommandIsNotExecutedByOwnPlayer)
						return
					}
					if err = httpHandler.executeSendPlayerIntoPortalCommand(worldIdDto, commandDto.PlayerId, commandDto.UnitId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case changePlayerHeldItemCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[changePlayerHeldItemCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if commandDto.PlayerId != playerIdDto {
						sendErroredServerEvent(ErrCommandIsNotExecutedByOwnPlayer)
						return
					}
					if err = httpHandler.executeChangePlayerHeldItemCommand(worldIdDto, commandDto.PlayerId, commandDto.ItemId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case createStaticUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createStaticUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateStaticUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
					); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case createFenceUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createFenceUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateFenceUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
					); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case createPortalUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createPortalUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreatePortalUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
					); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case createLinkUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createLinkUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateLinkUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
						commandDto.Label,
						commandDto.Url,
					); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case createEmbedUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[createEmbedUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeCreateEmbedUnitCommand(
						commandDto.UnitId,
						worldIdDto,
						commandDto.ItemId,
						commandDto.Position,
						commandDto.Direction,
						commandDto.Label,
						commandDto.EmbedCode,
					); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case rotateUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[rotateUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRotateUnitCommand(commandDto.UnitId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case removeStaticUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeStaticUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveStaticUnitCommand(commandDto.UnitId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case removeFenceUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeFenceUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveFenceUnitCommand(commandDto.UniId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case removePortalUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removePortalUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemovePortalUnitCommand(commandDto.UnitId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case removeLinkUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeLinkUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveLinkUnitCommand(commandDto.UnitId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				case removeEmbedUnitCommandName:
					commandRequestedClientEvent, err := jsonutil.Unmarshal[commandRequestedClientEvent[removeEmbedUnitCommand]](message)
					if err != nil {
						sendErroredServerEvent(err)
						return
					}
					commandDto := commandRequestedClientEvent.Command

					if err = httpHandler.executeRemoveEmbedUnitCommand(commandDto.UnitId); err != nil {
						broadcastCommandFailedServerEvent(commandDto.Id, err)
						break
					}
					broadcastCommandSucceededServerEvent(commandDto)
				default:
				}
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
	unitIdDto uuid.UUID,
) error {
	playerAppService := world_provide_dependency.ProvidePlayerAppService()
	return playerAppService.SendPlayerIntoPortal(playerappsrv.SendPlayerIntoPortalCommand{
		WorldId:  worldIdDto,
		PlayerId: playerIdDto,
		UnitId:   unitIdDto,
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

	staticUnitAppService := world_provide_dependency.ProvideStaticUnitAppService(uow)
	if err := staticUnitAppService.CreateStaticUnit(staticunitappsrv.CreateStaticUnitCommand{
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

	staticUnitAppService := world_provide_dependency.ProvideStaticUnitAppService(uow)
	if err := staticUnitAppService.RemoveStaticUnit(staticunitappsrv.RemoveStaticUnitCommand{
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

	fenceUnitAppService := world_provide_dependency.ProvideFenceUnitAppService(uow)
	if err := fenceUnitAppService.CreateFenceUnit(fenceunitappsrv.CreateFenceUnitCommand{
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

	fenceUnitAppService := world_provide_dependency.ProvideFenceUnitAppService(uow)
	if err := fenceUnitAppService.RemoveFenceUnit(fenceunitappsrv.RemoveFenceUnitCommand{
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

	portalUnitAppService := world_provide_dependency.ProvidePortalUnitAppService(uow)
	if err := portalUnitAppService.CreatePortalUnit(portalunitappsrv.CreatePortalUnitCommand{
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

	portalUnitAppService := world_provide_dependency.ProvidePortalUnitAppService(uow)
	if err := portalUnitAppService.RemovePortalUnit(portalunitappsrv.RemovePortalUnitCommand{
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
	label *string,
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
		Label:     label,
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

func (httpHandler *HttpHandler) executeCreateEmbedUnitCommand(
	idDto uuid.UUID,
	worldIdDto uuid.UUID,
	itemIdDto uuid.UUID,
	positionDto world_dto.PositionDto,
	directionDto int8,
	label *string,
	embedCode string,
) error {
	uow := pguow.NewUow()

	embedUnitAppService := world_provide_dependency.ProvideEmbedUnitAppService(uow)
	if err := embedUnitAppService.CreateEmbedUnit(embedunitappsrv.CreateEmbedUnitCommand{
		Id:        idDto,
		WorldId:   worldIdDto,
		ItemId:    itemIdDto,
		Position:  positionDto,
		Direction: directionDto,
		Label:     label,
		EmbedCode: embedCode,
	}); err != nil {
		uow.RevertChanges()
		return err
	}
	uow.SaveChanges()
	return nil
}

func (httpHandler *HttpHandler) executeRemoveEmbedUnitCommand(idDto uuid.UUID) error {
	uow := pguow.NewUow()

	embedUnitAppService := world_provide_dependency.ProvideEmbedUnitAppService(uow)
	if err := embedUnitAppService.RemoveEmbedUnit(embedunitappsrv.RemoveEmbedUnitCommand{
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

func (httpHandler *HttpHandler) enterWorld(worldIdDto uuid.UUID, userIdDto *uuid.UUID) (playerIdDto uuid.UUID, err error) {
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

func (httpHandler *HttpHandler) sendWorldEnteredResponse(worldIdDto uuid.UUID, playerIdDto uuid.UUID, sendMessage func(any)) error {
	uow := pguow.NewDummyUow()

	unitAppService := world_provide_dependency.ProvideUnitAppService(uow)
	worldAppService := world_provide_dependency.ProvideWorldAppService(uow)
	playerAppService := world_provide_dependency.ProvidePlayerAppService()

	worldDto, err := worldAppService.GetWorld(worldappsrv.GetWorldQuery{
		WorldId: worldIdDto,
	})
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

	sendMessage(worldEnteredServerEvent{
		Name:       worldEnteredServerEventName,
		World:      viewmodel.WorldViewModel(worldDto),
		Units:      unitDtos,
		MyPlayerId: playerIdDto,
		Players:    playerDtos,
	})
	return nil
}
