package scheduler

import (
	"time"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

var (
	Scheduler *gocron.Scheduler
)

func SetupScheduler() {
	Scheduler = gocron.NewScheduler(time.UTC)
	Scheduler.TagsUnique()
	log.Info().Msg("Starting task scheduler.")
	go syncScheduler()
	Scheduler.StartAsync()
}
