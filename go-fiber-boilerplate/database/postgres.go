package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // postgres driver
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

var (
	pgConn *sql.DB
	pgLog  zerolog.Logger
)

func PgInit(conf *config.Config) {
	pgLog = logger.Get("postgres")
	pgLog.Info().Msg("Connecting to Postgres server...")

	sslrootcert := ""
	if conf.Env == "production" {
		sslrootcert = "sslrootcert=./rds-combined-ca-bundle.pem"
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s %s sslmode=%s",
		conf.Database.Postgres.Host,
		conf.Database.Postgres.Port,
		conf.Database.Postgres.User,
		conf.Database.Postgres.Password,
		conf.Database.Postgres.Name,
		sslrootcert,
		conf.Database.Postgres.Sslmode,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		pgLog.Fatal().Err(err).Msg("Postgresql connection problem")
	}

	db.SetConnMaxLifetime(conf.Database.Postgres.MaxConnectionLifetime)
	db.SetMaxIdleConns(conf.Database.Postgres.MaxIdleConnections)
	db.SetMaxOpenConns(conf.Database.Postgres.MaxOpenConnections)
	pgConn = new(sql.DB)
	pgConn = db

	if !PgIsConnected() {
		pgLog.Fatal().Str("connection", connString).Msg("Postgress not connected")
	}

	pgLog.Info().Msg("Postgres connected")

}

func GetPgConnection() *sql.DB {
	return pgConn
}

func PgIsConnected() bool {
	if pgConn == nil {
		return false
	}

	err := pgConn.Ping()
	if err != nil {
		pgLog.Err(err).Msg("Postgres health check problem")
		return false
	}

	return true
}
