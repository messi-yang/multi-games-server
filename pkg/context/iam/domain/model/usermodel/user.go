package usermodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type User struct {
	id           globalcommonmodel.UserId
	emailAddress globalcommonmodel.EmailAddress
	username     globalcommonmodel.Username
	friendlyName FriendlyName
	createdAt    time.Time
	updatedAt    time.Time
}

// Interface Implementation Check
var _ domain.Aggregate[globalcommonmodel.UserId] = (*User)(nil)

func NewUser(
	id globalcommonmodel.UserId,
	emailAddress globalcommonmodel.EmailAddress,
	username globalcommonmodel.Username,
	friendlyName FriendlyName,
) User {
	return User{
		id:           id,
		emailAddress: emailAddress,
		username:     username,
		friendlyName: friendlyName,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}
}

func LoadUser(
	id globalcommonmodel.UserId,
	emailAddress globalcommonmodel.EmailAddress,
	username globalcommonmodel.Username,
	friendlyName FriendlyName,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	newUser := User{
		id:           id,
		emailAddress: emailAddress,
		username:     username,
		friendlyName: friendlyName,
		createdAt:    createdAt,
		updatedAt:    updatedAt,
	}
	return newUser
}

func (user *User) GetId() globalcommonmodel.UserId {
	return user.id
}

func (user *User) GetEmailAddress() globalcommonmodel.EmailAddress {
	return user.emailAddress
}

func (user *User) GetUsername() globalcommonmodel.Username {
	return user.username
}

func (user *User) UpdateUsername(username globalcommonmodel.Username) {
	user.username = username
}

func (user *User) GetFriendlyName() FriendlyName {
	return user.friendlyName
}

func (user *User) UpdateFriendlyName(friendlyName FriendlyName) {
	user.friendlyName = friendlyName
}

func (user *User) GetCreatedAt() time.Time {
	return user.createdAt
}

func (user *User) GetUpdatedAt() time.Time {
	return user.updatedAt
}
