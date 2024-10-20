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

type commandSentClientEvent struct {
	Name         clientEventName `json:"name"`
	PeerPlayerId uuid.UUID       `json:"peerPlayerId"`
	Command      dto.CommandDto  `json:"command"`
}

type commandRequestedClientEvent struct {
	Name    clientEventName `json:"name"`
	Command dto.CommandDto  `json:"command"`
}

type unitsFetchedClientEvent struct {
	Name     clientEventName  `json:"name"`
	BlockIds []dto.BlockIdDto `json:"blockIds"`
}

type p2pOfferSentClientEvent struct {
	Name          clientEventName `json:"name"`
	PeerPlayerId  uuid.UUID       `json:"peerPlayerId"`
	IceCandidates []any           `json:"iceCandidates"`
	Offer         any             `json:"offer"`
}

type p2pAnswerSentClientEvent struct {
	Name          clientEventName `json:"name"`
	PeerPlayerId  uuid.UUID       `json:"peerPlayerId"`
	IceCandidates []any           `json:"iceCandidates"`
	Answer        any             `json:"answer"`
}
