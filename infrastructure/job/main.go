package job

import "github.com/DumDumGeniuss/game-of-liberty-computer/infrastructure/job/gameupdatejob"

func StartJobs() {
	gameRoomJob := gameupdatejob.GetGameUpdateJob()
	gameRoomJob.Start()
}
