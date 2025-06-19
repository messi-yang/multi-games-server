package usecase

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/domaineventhandler/memdomaineventhandler"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/commandmodel"
	"github.com/dum-dum-genius/zossi-server/pkg/context/game/infrastructure/persistence/redisrepo"
)

type SaveCommandUseCase struct {
	commandRepo commandmodel.CommandRepo
}

func NewSaveCommandUseCase(commandRepo commandmodel.CommandRepo) SaveCommandUseCase {
	return SaveCommandUseCase{
		commandRepo: commandRepo,
	}
}

func ProvideSaveCommandUseCase(uow pguow.Uow) SaveCommandUseCase {
	domainEventDispatcher := memdomaineventhandler.NewDispatcher(uow)
	commandRepo := redisrepo.NewCommandRepo(domainEventDispatcher)
	return NewSaveCommandUseCase(commandRepo)
}

func (useCase *SaveCommandUseCase) Execute(commandDto dto.CommandDto) error {
	command := commandmodel.CreateCommand(
		commandmodel.NewCommandId(commandDto.Id),
		commandDto.Timestamp,
		commandDto.Name,
		commandDto.Payload,
	)

	if err := useCase.commandRepo.Add(command); err != nil {
		return err
	}

	return nil
}
