package gamecomputer

import "github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/jobs/tickunitmapjob"

func Start() {
	gameRoomJob := tickunitmapjob.GetJob()
	gameRoomJob.Start()
}
