package messages

import "github.com/DumDumGeniuss/game-of-liberty-computer/messages/game"

func InitializeMessages() {
	game.CreateGameRelatedTopics()
	game.ListTopics()
}
