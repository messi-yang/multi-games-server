package roomservicehandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/google/uuid"
)

type serverEventName string

const (
	roomJoinedServerEventName        serverEventName = "ROOM_JOINED"
	playerJoinedServerEventName      serverEventName = "PLAYER_JOINED"
	playerLeftServerEventName        serverEventName = "PLAYER_LEFT"
	commandReceivedServerEventName   serverEventName = "COMMAND_RECEIVED"
	commandFailedServerEventName     serverEventName = "COMMAND_FAILED"
	p2pOfferReceivedServerEventName  serverEventName = "P2P_OFFER_RECEIVED"
	p2pAnswerReceivedServerEventName serverEventName = "P2P_ANSWER_RECEIVED"
	erroredServerEventName           serverEventName = "ERRORED"
)

type serverEvent struct {
	Name serverEventName `json:"name"`
}

type roomJoinedServerEvent struct {
	Name       serverEventName  `json:"name"`
	Game       dto.GameDto      `json:"game"`
	Room       dto.RoomDto      `json:"room"`
	Commands   []dto.CommandDto `json:"commands"`
	MyPlayerId uuid.UUID        `json:"myPlayerId"`
	Players    []dto.PlayerDto  `json:"players"`
}

type playerJoinedServerEvent struct {
	Name   serverEventName `json:"name"`
	Player dto.PlayerDto   `json:"player"`
}

type playerLeftServerEvent struct {
	Name     serverEventName `json:"name"`
	PlayerId uuid.UUID       `json:"playerId"`
}

type commandReceivedServerEvent struct {
	Name    serverEventName `json:"name"`
	Command any             `json:"command"`
}

type commandFailedServerEvent struct {
	Name      serverEventName `json:"name"`
	CommandId uuid.UUID       `json:"commandId"`
}

type p2pOfferReceivedServerEvent struct {
	Name serverEventName `json:"name"`
	// Player that wanted to build connection with you
	PeerPlayerId  uuid.UUID `json:"peerPlayerId"`
	IceCandidates []any     `json:"iceCandidates"`
	Offer         any       `json:"offer"`
}

type p2pAnswerReceivedServerEvent struct {
	Name serverEventName `json:"name"`
	// Player that wanted to build connection with you
	PeerPlayerId  uuid.UUID `json:"peerPlayerId"`
	IceCandidates []any     `json:"iceCandidates"`
	Answer        any       `json:"answer"`
}

type erroredServerEvent struct {
	Name    serverEventName `json:"name"`
	Message string          `json:"message"`
}
