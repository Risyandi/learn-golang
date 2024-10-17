package logger

import (
	"path"
	"sync"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
)

type LumberjackLog interface {
	Run() *lumberjack.Logger
}

type lumberjackLog struct {
	LogPath       string
	CompressLog   bool
	DailyRotate   bool
	lastLogDate   string
	MaxBackup     int
	MaxAge        int
	MaxSize       int
	SleepDuration time.Duration
	lumberjackLog *lumberjack.Logger
}

func newLumberjackLogger(conf *config.Config) LumberjackLog {
	return &lumberjackLog{
		LogPath:       path.Join(conf.Log.Path, conf.Log.Filename),
		DailyRotate:   conf.Log.FileRotate,
		CompressLog:   true,
		MaxBackup:     conf.Log.MaxBackup,
		MaxAge:        conf.Log.MaxSize,
		MaxSize:       conf.Log.MaxSize,
		SleepDuration: 5,
	}
}

func (l *lumberjackLog) Run() *lumberjack.Logger {
	l.lumberjackLog = &lumberjack.Logger{
		Filename:   l.LogPath,
		Compress:   l.CompressLog,
		LocalTime:  true,
		MaxBackups: l.MaxBackup, // Max number of old log files to retain
		MaxAge:     l.MaxAge,    // Max number of days to retain old log files

	}

	l.lastLogDate = time.Now().Format("02012006")

	if l.DailyRotate {
		wg := &sync.WaitGroup{}

		wg.Add(1)
		go func() {
			l.rotate()
		}()
	}

	return l.lumberjackLog
}

func (l *lumberjackLog) rotate() {
	for {
		if l.lumberjackLog == nil {
			continue
		}

		now := time.Now().Format("02012006")
		if l.lastLogDate != now {
			l.lastLogDate = now
			l.lumberjackLog.Rotate()
		}

		time.Sleep(l.SleepDuration)
	}
}
