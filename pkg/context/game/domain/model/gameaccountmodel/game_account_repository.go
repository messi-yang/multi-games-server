package gameaccountmodel

import (
	"github.com/dum-dum-genius/zossi-server/pkg/context/global/domain/model/globalcommonmodel"
)

type GameAccountRepo interface {
	Add(GameAccount) error
	Update(GameAccount) error
	Get(GameAccountId) (GameAccount, error)
	GetAll() ([]GameAccount, error)
	GetGameAccountByUserId(globalcommonmodel.UserId) (gameAccount *GameAccount, err error)
	GetGameAccountOfUser(globalcommonmodel.UserId) (gameAccount GameAccount, err error)
}
