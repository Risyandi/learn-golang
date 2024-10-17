package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
)

var log zerolog.Logger

func InitLogger(conf *config.Config) {
	var writers []io.Writer

	if conf.Log.ConsoleEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true})
	}

	// WARNING!!! The Consecuent of this function is CPU usage will be over load.
	if conf.Log.FileEnabled {
		logFile := newLumberjackLogger(conf).Run()
		writers = append(writers, logFile)
	}
	mw := io.MultiWriter(writers...)

	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.TimeFieldFormat = time.RFC3339

	log = zerolog.New(mw).
		Level(conf.Log.Level).
		With().
		Timestamp().
		Logger()

}

func UpdateLogLevel(level zerolog.Level) {
	log = log.Level(level)
}

func Get(types string) zerolog.Logger {
	return log.With().Str("type", types).Caller().Logger()
}

func GetWithoutCaller(types string) *zerolog.Logger {
	logg := log.With().Str("type", types).Logger()
	return &logg
}
