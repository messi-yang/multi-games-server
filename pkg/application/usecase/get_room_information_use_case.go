package usecase

import (
	"errors"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commandmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/gamemodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/playermodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/roommodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/pgrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/redisrepo"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type GetRoomInformationUseCase struct {
	roomRepo    roommodel.RoomRepo
	gameRepo    gamemodel.GameRepo
	commandRepo commandmodel.CommandRepo
	playerRepo  playermodel.PlayerRepo
}

func NewGetRoomInformationUseCase(roomRepo roommodel.RoomRepo, gameRepo gamemodel.GameRepo, commandRepo commandmodel.CommandRepo, playerRepo playermodel.PlayerRepo) GetRoomInformationUseCase {
	return GetRoomInformationUseCase{roomRepo, gameRepo, commandRepo, playerRepo}
}

func ProvideGetRoomInformationUseCase(uow pguow.Uow) GetRoomInformationUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	roomRepo := pgrepo.NewRoomRepo(uow, domainEventDispatcher)
	gameRepo := pgrepo.NewGameRepo(uow, domainEventDispatcher)
	commandRepo := redisrepo.NewCommandRepo(domainEventDispatcher)
	playerRepo := redisrepo.NewPlayerRepo(domainEventDispatcher)

	return NewGetRoomInformationUseCase(roomRepo, gameRepo, commandRepo, playerRepo)
}

func (useCase *GetRoomInformationUseCase) Execute(roomIdDto uuid.UUID) (
	roomDto dto.RoomDto, gameDto dto.GameDto, commandDtos []dto.CommandDto, playerDtos []dto.PlayerDto, err error) {
	roomId := globalcommonmodel.NewRoomId(roomIdDto)

	room, err := useCase.roomRepo.Get(roomId)
	if err != nil {
		return roomDto, gameDto, commandDtos, playerDtos, err
	}
	roomDto = dto.NewRoomDto(room)

	currentGameId := room.GetCurrentGameId()
	if currentGameId == nil {
		return roomDto, gameDto, commandDtos, playerDtos, errors.New("this room has no current game")
	}

	game, err := useCase.gameRepo.Get(*currentGameId)
	if err != nil {
		return roomDto, gameDto, commandDtos, playerDtos, err
	}
	gameDto = dto.NewGameDto(game)

	commands, err := useCase.commandRepo.GetCommandsOfGame(game.GetId())
	if err != nil {
		return roomDto, gameDto, commandDtos, playerDtos, err
	}
	commandDtos = lo.Map(commands, func(_command commandmodel.Command, _ int) dto.CommandDto {
		return dto.NewCommandDto(_command)
	})

	players, err := useCase.playerRepo.GetPlayersOfRoom(roomId)
	if err != nil {
		return roomDto, gameDto, commandDtos, playerDtos, err
	}
	playerDtos = lo.Map(players, func(_player playermodel.Player, _ int) dto.PlayerDto {
		return dto.NewPlayerDto(_player)
	})

	return roomDto, gameDto, commandDtos, playerDtos, nil
}
