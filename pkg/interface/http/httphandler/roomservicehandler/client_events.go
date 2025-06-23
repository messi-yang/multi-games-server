package roomservicehandler

import (
	"github.com/dum-dum-genius/zossi-server/pkg/application/dto"
	"github.com/google/uuid"
)

type clientEventName string

const (
	pingClientEventName                  clientEventName = "PING"
	startGameRequestedClientEventName    clientEventName = "START_GAME_REQUESTED"
	setupNewGameRequestedClientEventName clientEventName = "SETUP_NEW_GAME_REQUESTED"
	commandSentClientEventName           clientEventName = "COMMAND_SENT"
	commandExecutedClientEventName       clientEventName = "COMMAND_EXECUTED"
	p2pOfferSentClientEventName          clientEventName = "P2P_OFFER_SENT"
	p2pAnswerSentClientEventName         clientEventName = "P2P_ANSWER_SENT"
	p2pConnectedClientEventName          clientEventName = "P2P_CONNECTED"
)

type clientEvent struct {
	Name clientEventName `json:"name"`
}

type startGameRequestedClientEvent struct {
	Name      clientEventName        `json:"name"`
	GameId    uuid.UUID              `json:"gameId"`
	GameState map[string]interface{} `json:"gameState"`
}

type setupNewGameRequestedClientEvent struct {
	Name     clientEventName `json:"name"`
	GameName string          `json:"gameName"`
}

type commandSentClientEvent struct {
	Name         clientEventName `json:"name"`
	PeerPlayerId uuid.UUID       `json:"peerPlayerId"`
	Command      dto.CommandDto  `json:"command"`
}

type commandExecutedClientEvent struct {
	Name    clientEventName `json:"name"`
	Command dto.CommandDto  `json:"command"`
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

type p2pConnectedClientEvent struct {
	Name         clientEventName `json:"name"`
	PeerPlayerId uuid.UUID       `json:"peerPlayerId"`
}
