package client

import (
	"fmt"
	"main/database"
	"main/ent"
	"net"
	"strconv"
	"time"

	"github.com/rs/zerolog/log"
)

func TcpClient(t *ent.Target) {
	contextLogger := log.With().Str("target_name", t.Name).Str("type", t.Type.String()).Logger()

	start := time.Now()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(*t.Hostname, strconv.Itoa(int(*t.Port))), time.Duration(time.Duration(t.TimeoutSeconds).Seconds()))
	duration := time.Since(start).Milliseconds()
	state := 1
	status := ""
	if err != nil {
		contextLogger.Debug().Err(err).Msg("target error.")
		status = err.Error()
	} else if conn == nil {
		contextLogger.Debug().Msg("connection is nil.")
	} else {
		conn.Close()
		state = 0
	}
	contextLogger.Debug().Msg(fmt.Sprintf("target response: %d", state))
	queryInsert := `INSERT INTO uptime_history (time, target_id, response_time_ms, response_status, response_state) VALUES (NOW(), $1, $2, $3, $4);`
	_, dbErr := database.DB.Exec(
		queryInsert, t.ID,
		duration,
		status,
		state,
	)
	if dbErr != nil {
		contextLogger.Error().Err(dbErr).Msg("couldn't write time series data.")
	}
}
