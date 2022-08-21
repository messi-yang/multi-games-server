package job

import "github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/job/tickunitmapjob"

func StartJobs() {
	gameRoomJob := tickunitmapjob.GetTickUnitMapJob()
	gameRoomJob.Start()
}
