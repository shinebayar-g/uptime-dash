package client

import (
	"fmt"
	"main/database"
	"main/ent"
	"net"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
)

func isTimeoutError(err error) bool {
	e, ok := err.(net.Error)
	return ok && e.Timeout()
}

func HttpClient(t *ent.Target) {
	contextLogger := log.With().Str("target_name", t.Name).Str("type", t.Type.String()).Logger()

	client := resty.New()
	client.SetTimeout(time.Duration(t.TimeoutSeconds) * time.Second)
	res, err := client.R().EnableTrace().Get(*t.URL)

	lat := res.Request.TraceInfo().ResponseTime.Milliseconds()
	status := res.Status()
	if err != nil {
		contextLogger.Debug().Err(err).Msg("http client error.")
		status = err.Error()
		if isTimeoutError(err) {
			lat = int64(t.TimeoutSeconds) * time.Second.Milliseconds()
		}
	}

	state := 1
	if err == nil && res.StatusCode() >= 200 && res.StatusCode() < 300 {
		state = 0
	}

	contextLogger.Debug().Msg(fmt.Sprintf("target response: %s", res.Status()))
	queryInsert := `INSERT INTO uptime_history (time, target_id, response_time_ms, response_status, response_state) VALUES (NOW(), $1, $2, $3, $4);`
	_, dbErr := database.DB.Exec(
		queryInsert, t.ID,
		lat,
		status,
		state,
	)
	if dbErr != nil {
		contextLogger.Error().Err(dbErr).Msg("couldn't write time series data.")
	}
}
