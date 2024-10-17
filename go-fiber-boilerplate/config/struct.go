package config

import (
	"time"

	"github.com/rs/zerolog"
)

type (
	Config struct {
		Env string
		API struct {
			Name  string
			Route string
			Port  int
		}
		Context struct {
			Timeout int
		}
		Database struct {
			Postgres PostgresConfig
			Mongo    MongoConfig
			Redis    RedisConfig
		}
		Encryption EncryptionConfig
		Service    string
		Webhook    string
		ApiKey     ApiKeyConfig
		Log        LogConfig
		Lang       LangConfig
	}

	PostgresConfig struct {
		Host                  string
		Port                  int
		Name                  string
		User                  string
		Password              string
		Sslmode               string
		MaxOpenConnections    int
		MaxIdleConnections    int
		MaxConnectionLifetime time.Duration
	}

	MongoConfig struct {
		Host     string
		Port     int
		Username string
		Password string
		Name     string
		SRV      bool
	}

	RedisConfig struct {
		Host            string
		Port            int
		Auth            string
		Db              int
		DbAuth          int
		ExpiredDuration time.Duration
		DialTimeout     time.Duration
	}

	EncryptionConfig struct {
		AccessTokenPassphrase  string
		RefreshTokenPassphrase string
		AdminTokenPassphrase   string
	}

	ApiKeyConfig struct {
		Pastebin string
	}

	LogConfig struct {
		Level          zerolog.Level
		Path           string
		Filename       string
		FileRotate     bool
		MaxSize        int
		MaxBackup      int
		MaxAge         int
		ConsoleEnabled bool
		FileEnabled    bool
	}

	LangConfig struct {
		DefaultLanguage   string
		PathLanguage      string
		AcceptedLanguages []string
	}
)
