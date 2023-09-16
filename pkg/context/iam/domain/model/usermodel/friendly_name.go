package usermodel

import (
	"fmt"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
)

var (
	ErrFriendlyNameLengthNotValid = fmt.Errorf("length of friendly name should be between 1 and 20")
)

type FriendlyName struct {
	value string
}

// Interface Implementation Check
var _ domain.ValueObject[FriendlyName] = (*FriendlyName)(nil)

func NewFriendlyName(value string) (friendlyName FriendlyName, err error) {
	if len(value) < 1 || len(value) > 20 {
		return friendlyName, ErrFriendlyNameLengthNotValid
	}
	return FriendlyName{
		value: value,
	}, nil
}

func (friendlyName FriendlyName) IsEqual(friendlyNameB FriendlyName) bool {
	return friendlyName.value == friendlyNameB.value
}

func (friendlyName FriendlyName) String() string {
	return friendlyName.value
}
