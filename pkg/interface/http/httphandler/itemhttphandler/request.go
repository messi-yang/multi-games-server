package itemhttphandler

import "github.com/google/uuid"

type getItemsOfIdsRequestBody struct {
	ItemIds []uuid.UUID `json:"itemIds"`
}
