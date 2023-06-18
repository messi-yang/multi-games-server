package sharedkernelmodel

import (
	"fmt"
	"math/rand"
	"regexp"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

type Username struct {
	value string
}

// Interface Implementation Check
var _ domain.ValueObject[Username] = (*Username)(nil)

func NewUsername(value string) (username Username, err error) {
	match, _ := regexp.MatchString("^[a-z_]+$", value)
	if !match {
		return username, fmt.Errorf("email is not valid")
	}
	match, _ = regexp.MatchString("_{2,}", value)
	if match {
		return username, fmt.Errorf("cannot have two underscores in a row")
	}
	match, _ = regexp.MatchString("^_|_$", value)
	if match {
		return username, fmt.Errorf("cannot have underscores at the start or the end")
	}
	match, _ = regexp.MatchString("^.{8,20}$", value)
	if !match {
		return username, fmt.Errorf("cannot only username with length between 8 to 20")
	}
	return Username{
		value: value,
	}, nil
}

func NewRandomUsername() Username {
	const availableCharacters = "abcdefghijklmnopqrstuvwxyz"
	randomUsernameBytes := make([]byte, 12)
	for i := range randomUsernameBytes {
		randomUsernameBytes[i] = availableCharacters[rand.Intn(len(availableCharacters))]
	}
	return Username{
		value: fmt.Sprintf("newuser_%s", string(randomUsernameBytes)),
	}
}

func (username Username) IsEqual(otherUsername Username) bool {
	return username.value == otherUsername.value
}

func (username Username) String() string {
	return username.value
}
