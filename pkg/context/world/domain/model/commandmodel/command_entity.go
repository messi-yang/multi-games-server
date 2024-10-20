package commandmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/util/jsonutil"
)

type CommandEntity struct {
	id        CommandId
	timestamp int64
	name      CommandName
	payload   any
}

// Interface Implementation Check
var _ domain.Entity[CommandId] = (*CommandEntity)(nil)

func NewCommandEntity(
	id CommandId,
	timestamp int64,
	name CommandName,
	payload any,
) CommandEntity {
	return CommandEntity{
		id:        id,
		timestamp: timestamp,
		name:      name,
		payload:   payload,
	}
}

func LoadCommandEntity(
	id CommandId,
	timestamp int64,
	name CommandName,
	payload any,
) CommandEntity {
	return CommandEntity{
		id:        id,
		timestamp: timestamp,
		name:      name,
		payload:   payload,
	}
}

func (command *CommandEntity) GetId() CommandId {
	return command.id
}

func (command *CommandEntity) GetTimestamp() int64 {
	return command.timestamp
}

func (command *CommandEntity) GetCommandName() CommandName {
	return command.name
}

func (command *CommandEntity) GetPayload() any {
	return command.name
}

func (command *CommandEntity) GetChangePlayerActionCommandPayload() (ChangePlayerActionCommandPayloadJson, error) {
	return jsonutil.Unmarshal[ChangePlayerActionCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetSendPlayerIntoPortalCommandPayload() (SendPlayerIntoPortalCommandPayloadJson, error) {
	return jsonutil.Unmarshal[SendPlayerIntoPortalCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetChangePlayerHeldItemCommandPayload() (ChangePlayerHeldItemCommandPayloadJson, error) {
	return jsonutil.Unmarshal[ChangePlayerHeldItemCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetCreateStaticUnitCommandPayload() (CreateStaticUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[CreateStaticUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetRemoveStaticUnitCommandPayload() (RemoveStaticUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[RemoveStaticUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetCreateFenceUnitCommandPayload() (CreateFenceUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[CreateFenceUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetRemoveFenceUnitCommandPayload() (RemoveFenceUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[RemoveFenceUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetCreatePortalUnitCommandPayload() (CreatePortalUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[CreatePortalUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetRemovePortalUnitCommandPayload() (RemovePortalUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[RemovePortalUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetCreateLinkUnitCommandPayload() (CreateLinkUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[CreateLinkUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetRemoveLinkUnitCommandPayload() (RemoveLinkUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[RemoveLinkUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetCreateEmbedUnitCommandPayload() (CreateEmbedUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[CreateEmbedUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetRemoveEmbedUnitCommandPayload() (RemoveEmbedUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[RemoveEmbedUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
func (command *CommandEntity) GetRotateUnitCommandPayload() (RotateUnitCommandPayloadJson, error) {
	return jsonutil.Unmarshal[RotateUnitCommandPayloadJson](jsonutil.Marshal(command.payload))
}
