package playermodel

import "fmt"

type PlayerActionNameEnum string

const (
	PlayerActionNameEnumStand    PlayerActionNameEnum = "stand"
	PlayerActionNameEnumWalk     PlayerActionNameEnum = "walk"
	PlayerActionNameEnumTeleport PlayerActionNameEnum = "teleport"
)

func ParsePlayerActionNameEnum(value string) (PlayerActionNameEnum, error) {
	switch value {
	case string(PlayerActionNameEnumStand):
		return PlayerActionNameEnumStand, nil
	case string(PlayerActionNameEnumWalk):
		return PlayerActionNameEnumWalk, nil
	case string(PlayerActionNameEnumTeleport):
		return PlayerActionNameEnumTeleport, nil
	default:
		return "", fmt.Errorf("invalid player name: %s", value)
	}
}
