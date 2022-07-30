package job

import "github.com/dum-dum-genius/game-of-liberty-computer/infrastructure/job/gamecomputejob"

func StartJobs() {
	gameRoomJob := gamecomputejob.GetGameComputeJob()
	gameRoomJob.Start()
}
