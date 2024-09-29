package worldjourneyhandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/google/uuid"
)

type clientEventName string

const (
	pingClientEventName             clientEventName = "PING"
	commandSentClientEventName      clientEventName = "COMMAND_SENT"
	commandRequestedClientEventName clientEventName = "COMMAND_REQUESTED"
	unitsFetchedClientEventName     clientEventName = "UNITS_FETCHED"
	p2pOfferSentClientEventName     clientEventName = "P2P_OFFER_SENT"
	p2pAnswerSentClientEventName    clientEventName = "P2P_ANSWER_SENT"
)

type clientEvent struct {
	Name clientEventName `json:"name"`
}

// This command is used when the P2P connection with the peer player fails
type commandSentClientEvent struct {
	Name         clientEventName `json:"name"`
	PeerPlayerId uuid.UUID       `json:"peerPlayerId"`
	// The command will possibly be client-only command, so we use type "any" here
	Command any `json:"command"`
}

type commandRequestedClientEvent[T any] struct {
	Name    clientEventName `json:"name"`
	Command T               `json:"command"`
}

type unitsFetchedClientEvent struct {
	Name   clientEventName `json:"name"`
	Blocks []dto.BlockDto  `json:"blocks"`
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
