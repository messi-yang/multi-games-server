package gamecomputer

import "github.com/dum-dum-genius/game-of-liberty-computer/gamecomputer/jobs/tickunitmapjob"

func StartJobs() {
	gameRoomJob := tickunitmapjob.GetJob()
	gameRoomJob.Start()
}
