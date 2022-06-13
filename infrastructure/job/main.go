package job

import "github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/job/gamecomputejob"

func StartJobs() {
	gameRoomJob := gamecomputejob.GetGameComputeJob()
	gameRoomJob.Start()
}
