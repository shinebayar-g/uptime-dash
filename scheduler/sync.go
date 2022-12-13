package scheduler

import (
	"context"
	"fmt"
	"main/client"
	"main/database"
	"main/ent/target"
	"math/rand"
	"time"

	"github.com/rs/zerolog/log"
)

func syncScheduler() {
	log.Info().Msg("Syncing tasks from database.")
	entTargets, err := database.Client.Target.
		Query().All(context.Background())
	if err != nil {
		log.Err(err).Msg("couldn't query targets.")
	}

	registerTargets := 0
	for _, t := range entTargets {
		rand.Seed(time.Now().UnixNano())
		r := rand.Intn(5)
		time.Sleep(time.Duration(r) * time.Second)

		log.Info().Str("name", t.Name).Str("type", t.Type.String()).Uint32("intervalSeconds", t.IntervalSeconds).Msg("Registering target")

		var err error
		switch t.Type {
		case target.TypeTYPE_HTTP:
			_, err = Scheduler.Every(int(t.IntervalSeconds)).Seconds().Tag(t.ID.String()).Do(client.HttpClient, t)
		case target.TypeTYPE_TCP:
			_, err = Scheduler.Every(int(t.IntervalSeconds)).Seconds().Tag(t.ID.String()).Do(client.TcpClient, t)
		case target.TypeTYPE_PING:
		}

		if err != nil {
			log.Error().Err(err).Str("name", t.Name).Msg("couldn't register task.")
		} else {
			registerTargets += 1
		}
	}
	log.Info().Msg(fmt.Sprintf("Registered %d/%d target(s)", registerTargets, len(entTargets)))
}
