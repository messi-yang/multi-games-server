package dto

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/game/domain/model/messagemodel"
	"github.com/google/uuid"
)

type MessageDto struct {
	Id         uuid.UUID `json:"id"`
	PlayerName *string   `json:"playerName"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
}

func NewMessageDto(
	message messagemodel.MessageModel,
) MessageDto {
	return MessageDto{
		Id:         message.GetId(),
		PlayerName: message.GetPlayerName(),
		Content:    message.GetContent(),
		CreatedAt:  message.GetCreatedAt(),
	}
}

func ParseMessageDto(messageDto MessageDto) (message messagemodel.MessageModel) {
	return messagemodel.LoadMessage(
		messageDto.Id,
		messageDto.PlayerName,
		messageDto.Content,
		messageDto.CreatedAt,
	)
}
