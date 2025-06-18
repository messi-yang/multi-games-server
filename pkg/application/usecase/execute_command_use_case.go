package usecase

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/persistence/pguow"
	"github.com/dum-dum-genius/zossi-server/pkg/context/world/domain/model/commandmodel"
	"github.com/google/uuid"
)

type ExecuteCommandUseCase struct{}

func NewExecuteCommandUseCase() ExecuteCommandUseCase {
	return ExecuteCommandUseCase{}
}

func ProvideExecuteCommandUseCase(uow pguow.Uow) ExecuteCommandUseCase {
	return NewExecuteCommandUseCase()
}

func (useCase *ExecuteCommandUseCase) Execute(worldIdDto uuid.UUID, commandDto dto.CommandDto) error {
	command := commandmodel.CreateCommand(
		commandmodel.NewCommandId(commandDto.Id),
		commandDto.Timestamp,
		commandDto.Name,
		commandDto.Payload,
	)
	fmt.Println(command)

	// TODO - execute command

	return nil
}
