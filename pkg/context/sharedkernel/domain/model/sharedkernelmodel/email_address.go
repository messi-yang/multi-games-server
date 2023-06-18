package sharedkernelmodel

import (
	"fmt"
	"regexp"

	"github.com/dum-dum-genius/zossi-server/pkg/context/sharedkernel/domain"
)

type EmailAddress struct {
	value string
}

// Interface Implementation Check
var _ domain.ValueObject[EmailAddress] = (*EmailAddress)(nil)

func NewEmailAddress(value string) (emailAddress EmailAddress, err error) {
	pattern := "(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])"
	match, err := regexp.MatchString(pattern, value)
	if err != nil {
		return emailAddress, err
	}
	if !match {
		return emailAddress, fmt.Errorf("email is not valid")
	}
	return EmailAddress{
		value: value,
	}, nil
}

func (emailAddress EmailAddress) IsEqual(otherEmailAddress EmailAddress) bool {
	return emailAddress.String() == otherEmailAddress.String()
}

func (emailAddress EmailAddress) String() string {
	return emailAddress.value
}
