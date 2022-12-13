package database

import (
	"context"
	"database/sql"
	"fmt"
	"main/ent"
	"os"

	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rs/zerolog/log"
)

var (
	Client *ent.Client
	DB     *sql.DB
)

func open(dsn string) *ent.Client {
	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database.")
	}

	// Create an ent.Driver from `DB`.
	drv := entsql.OpenDB("postgres", DB)
	return ent.NewClient(ent.Driver(drv))
}

func migrateSchema() {
	if err := Client.Schema.Create(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate schema.")
	}
	log.Info().Msg("Successfully applied migrations.")
}

func createHyperTable() {
	queryCreateHypertable := `CREATE TABLE IF NOT EXISTS uptime_history (
		time TIMESTAMPTZ NOT NULL,
		target_id uuid NOT NULL,
		response_time_ms INT2 NOT NULL,
		response_status TEXT,
		response_state INT2 NOT NULL
		);
		SELECT create_hypertable('uptime_history','time', if_not_exists => TRUE);
		CREATE INDEX IF NOT EXISTS target_id_time_idx ON uptime_history (target_id, time DESC);
		`
	_, err := DB.Exec(queryCreateHypertable)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create hyper table.")
	}
	log.Info().Msg("Hyper table is created.")
}

func SetupDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, name)
	Client = open(dsn)
	log.Info().Msg("Connected to database.")
	migrateSchema()
	createHyperTable()
}
