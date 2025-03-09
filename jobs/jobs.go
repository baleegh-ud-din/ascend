package jobs

import "github.com/robfig/cron/v3"

var c = cron.New()

func StartJobs() {

	c.AddFunc("@every 1m", func() {})

	c.Start()
}

func StopJobs() {
	c.Stop()
}
