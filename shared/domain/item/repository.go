package item

import (
	"github.com/google/uuid"
)

type ItemRepository interface {
	Add(item Item) error
	Get(itemId uuid.UUID) (item Item, err error)
	Update(itemId uuid.UUID, item Item) error
}
