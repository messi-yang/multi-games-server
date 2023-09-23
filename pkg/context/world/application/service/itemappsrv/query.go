package itemappsrv

import "github.com/google/uuid"

type QueryItemsQuery struct {
}

type GetItemsOfIdsQuery struct {
	ItemIds []uuid.UUID
}
