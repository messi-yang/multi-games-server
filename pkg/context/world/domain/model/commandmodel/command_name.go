package commandmodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

type commandNameValue string

const (
	changePlayerActionCommandName   commandNameValue = "CHANGE_PLAYER_ACTION"
	sendPlayerIntoPortalCommandName commandNameValue = "SEND_PLAYER_INTO_PORTAL"
	changePlayerHeldItemCommandName commandNameValue = "CHANGE_PLAYER_HELD_ITEM"
	createStaticUnitCommandName     commandNameValue = "CREATE_STATIC_UNIT"
	removeStaticUnitCommandName     commandNameValue = "REMOVE_STATIC_UNIT"
	createFenceUnitCommandName      commandNameValue = "CREATE_FENCE_UNIT"
	removeFenceUnitCommandName      commandNameValue = "REMOVE_FENCE_UNIT"
	createPortalUnitCommandName     commandNameValue = "CREATE_PORTAL_UNIT"
	removePortalUnitCommandName     commandNameValue = "REMOVE_PORTAL_UNIT"
	createLinkUnitCommandName       commandNameValue = "CREATE_LINK_UNIT"
	removeLinkUnitCommandName       commandNameValue = "REMOVE_LINK_UNIT"
	createEmbedUnitCommandName      commandNameValue = "CREATE_EMBED_UNIT"
	removeEmbedUnitCommandName      commandNameValue = "REMOVE_EMBED_UNIT"
	createColorUnitCommandName      commandNameValue = "CREATE_COLOR_UNIT"
	removeColorUnitCommandName      commandNameValue = "REMOVE_COLOR_UNIT"
	rotateUnitCommandName           commandNameValue = "ROTATE_UNIT"
)

type CommandName struct {
	value commandNameValue
}

// Interface Implementation Check
var _ domain.ValueObject[CommandName] = (*CommandName)(nil)

func NewCommandName(value string) (CommandName, error) {
	switch value {
	case "CHANGE_PLAYER_ACTION":
		return CommandName{
			value: changePlayerActionCommandName,
		}, nil
	case "SEND_PLAYER_INTO_PORTAL":
		return CommandName{
			value: sendPlayerIntoPortalCommandName,
		}, nil
	case "CHANGE_PLAYER_HELD_ITEM":
		return CommandName{
			value: changePlayerHeldItemCommandName,
		}, nil
	case "CREATE_STATIC_UNIT":
		return CommandName{
			value: createStaticUnitCommandName,
		}, nil
	case "REMOVE_STATIC_UNIT":
		return CommandName{
			value: removeStaticUnitCommandName,
		}, nil
	case "CREATE_FENCE_UNIT":
		return CommandName{
			value: createFenceUnitCommandName,
		}, nil
	case "REMOVE_FENCE_UNIT":
		return CommandName{
			value: removeFenceUnitCommandName,
		}, nil
	case "CREATE_PORTAL_UNIT":
		return CommandName{
			value: createPortalUnitCommandName,
		}, nil
	case "REMOVE_PORTAL_UNIT":
		return CommandName{
			value: removePortalUnitCommandName,
		}, nil
	case "CREATE_LINK_UNIT":
		return CommandName{
			value: createLinkUnitCommandName,
		}, nil
	case "REMOVE_LINK_UNIT":
		return CommandName{
			value: removeLinkUnitCommandName,
		}, nil
	case "CREATE_EMBED_UNIT":
		return CommandName{
			value: createEmbedUnitCommandName,
		}, nil
	case "REMOVE_EMBED_UNIT":
		return CommandName{
			value: removeEmbedUnitCommandName,
		}, nil
	case "CREATE_COLOR_UNIT":
		return CommandName{
			value: createColorUnitCommandName,
		}, nil
	case "REMOVE_COLOR_UNIT":
		return CommandName{
			value: removeColorUnitCommandName,
		}, nil
	case "ROTATE_UNIT":
		return CommandName{
			value: rotateUnitCommandName,
		}, nil
	default:
		return CommandName{}, fmt.Errorf("invalid CommandName: %s", value)
	}
}

func (commandName CommandName) IsChangePlayerActionCommandName() bool {
	return commandName.value == changePlayerActionCommandName
}

func (commandName CommandName) IsSendPlayerIntoPortalCommandName() bool {
	return commandName.value == sendPlayerIntoPortalCommandName
}

func (commandName CommandName) IsChangePlayerHeldItemCommandName() bool {
	return commandName.value == changePlayerHeldItemCommandName
}

func (commandName CommandName) IsCreateStaticUnitCommandName() bool {
	return commandName.value == createStaticUnitCommandName
}

func (commandName CommandName) IsRemoveStaticUnitCommandName() bool {
	return commandName.value == removeStaticUnitCommandName
}

func (commandName CommandName) IsCreateFenceUnitCommandName() bool {
	return commandName.value == createFenceUnitCommandName
}

func (commandName CommandName) IsRemoveFenceUnitCommandName() bool {
	return commandName.value == removeFenceUnitCommandName
}

func (commandName CommandName) IsCreatePortalUnitCommandName() bool {
	return commandName.value == createPortalUnitCommandName
}

func (commandName CommandName) IsRemovePortalUnitCommandName() bool {
	return commandName.value == removePortalUnitCommandName
}

func (commandName CommandName) IsCreateLinkUnitCommandName() bool {
	return commandName.value == createLinkUnitCommandName
}

func (commandName CommandName) IsRemoveLinkUnitCommandName() bool {
	return commandName.value == removeLinkUnitCommandName
}

func (commandName CommandName) IsCreateEmbedUnitCommandName() bool {
	return commandName.value == createEmbedUnitCommandName
}

func (commandName CommandName) IsRemoveEmbedUnitCommandName() bool {
	return commandName.value == removeEmbedUnitCommandName
}

func (commandName CommandName) IsCreateColorUnitCommandName() bool {
	return commandName.value == createColorUnitCommandName
}

func (commandName CommandName) IsRemoveColorUnitCommandName() bool {
	return commandName.value == removeColorUnitCommandName
}

func (commandName CommandName) IsRotateUnitCommandName() bool {
	return commandName.value == rotateUnitCommandName
}

func (commandName CommandName) IsEqual(otherCommandName CommandName) bool {
	return commandName.value == otherCommandName.value
}
