package worldjourneyhandler

import "github.com/google/uuid"

type clientEventName string

const (
	pingClientEventName             clientEventName = "PING"
	commandRequestedClientEventName clientEventName = "COMMAND_REQUESTED"
	p2pOfferSentClientEventName     clientEventName = "P2P_OFFER_SENT"
	p2pAnswerSentClientEventName    clientEventName = "P2P_ANSWER_SENT"
)

type clientEvent struct {
	Name clientEventName `json:"name"`
}

type commandRequestedClientEvent[T any] struct {
	Name    clientEventName `json:"name"`
	Command T               `json:"command"`
}

type p2pOfferSentClientEvent struct {
	Name clientEventName `json:"name"`
	// Player that the client wanted to build connection with
	PeerPlayerId  uuid.UUID `json:"peerPlayerId"`
	IceCandidates []any     `json:"iceCandidates"`
	Offer         any       `json:"offer"`
}

type p2pAnswerSentClientEvent struct {
	Name clientEventName `json:"name"`
	// Player that the client wanted to build connection with
	PeerPlayerId  uuid.UUID `json:"peerPlayerId"`
	IceCandidates []any     `json:"iceCandidates"`
	Answer        any       `json:"answer"`
}
