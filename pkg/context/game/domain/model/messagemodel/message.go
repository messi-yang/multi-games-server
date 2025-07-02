package messagemodel

import (
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/domain"
	"github.com/google/uuid"
)

type MessageModel struct {
	Id                   uuid.UUID
	PlayerName           *string
	Content              string
	CreatedAt            time.Time
	domainEventCollector *domain.DomainEventCollector
}

// Interface Implementation Check
var _ domain.DomainEventDispatchableAggregate = (*MessageModel)(nil)

func LoadMessage(
	id uuid.UUID,
	playerName *string,
	content string,
	createdAt time.Time,
) MessageModel {
	return MessageModel{
		Id:                   uuid.New(),
		PlayerName:           playerName,
		Content:              content,
		CreatedAt:            time.Now(),
		domainEventCollector: domain.NewDomainEventCollector(),
	}
}

func (message *MessageModel) PopDomainEvents() []domain.DomainEvent {
	return message.domainEventCollector.PopAll()
}

func (message *MessageModel) GetId() uuid.UUID {
	return message.Id
}

func (message *MessageModel) GetPlayerName() *string {
	return message.PlayerName
}

func (message *MessageModel) GetContent() string {
	return message.Content
}

func (message *MessageModel) GetCreatedAt() time.Time {
	return message.CreatedAt
}
