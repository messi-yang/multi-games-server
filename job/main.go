package job

import "github.com/dum-dum-genius/game-of-liberty-computer/job/tickunitmapjob"

func StartJobs() {
	gameRoomJob := tickunitmapjob.GetJob()
	gameRoomJob.Start()
}
