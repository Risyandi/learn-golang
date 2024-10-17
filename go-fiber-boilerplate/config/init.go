package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var conf *Config

func Init() {
	// checking file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Viper to read environment variables
	viper.AutomaticEnv()
	conf = new(Config)

	conf.Env = viper.GetString("ENVIRONTMENT")
	conf.API.Name = viper.GetString("API_NAME")
	conf.API.Route = viper.GetString("API_ROUTE_NAME")
	conf.API.Port = viper.GetInt("API_PORT")
	conf.Context.Timeout = viper.GetInt("CONTEXT_TIMEOUT")
	conf.Database.Postgres.Host = viper.GetString("DB_PG_HOST")
	conf.Database.Postgres.Port = viper.GetInt("DB_PG_PORT")
	conf.Database.Postgres.User = viper.GetString("DB_PG_USER")
	conf.Database.Postgres.Password = viper.GetString("DB_PG_PASSWORD")
	conf.Database.Postgres.Name = viper.GetString("DB_PG_NAME")
	conf.Database.Postgres.Sslmode = viper.GetString("DB_PG_SSLMODE")
	conf.Database.Postgres.MaxOpenConnections = viper.GetInt("DB_PG_MAX_OPEN_CONNECTION")
	conf.Database.Postgres.MaxIdleConnections = viper.GetInt("DB_PG_MAX_IDLE_CONNECTION")
	conf.Database.Postgres.MaxConnectionLifetime = viper.GetDuration("DB_PG_MAX_CONNECTION_LIFETIME")
	conf.Database.Mongo.Host = viper.GetString("DB_MONGO_HOST")
	conf.Database.Mongo.Port = viper.GetInt("DB_MONGO_PORT")
	conf.Database.Mongo.Username = viper.GetString("DB_MONGO_USER")
	conf.Database.Mongo.Password = viper.GetString("DB_MONGO_PASSWORD")
	conf.Database.Mongo.Name = viper.GetString("DB_MONGO_NAME")
	conf.Database.Mongo.SRV = viper.GetBool("DB_MONGO_SRV")
	conf.Database.Redis.Host = viper.GetString("REDIS_HOST")
	conf.Database.Redis.Port = viper.GetInt("REDIS_PORT")
	conf.Database.Redis.Auth = viper.GetString("REDIS_AUTH")
	conf.Database.Redis.Db = viper.GetInt("REDIS_DB")
	conf.Database.Redis.DbAuth = viper.GetInt("REDIS_DB_AUTH")
	conf.Database.Redis.ExpiredDuration = viper.GetDuration("REDIS_EXPIRED_DURATION")
	conf.Database.Redis.DialTimeout = viper.GetDuration("REDIS_DIAL_TIMEOUT")

	conf.Log.Level, err = zerolog.ParseLevel(viper.GetString("LOG_LEVEL"))
	if err != nil {
		conf.Log.Level = zerolog.InfoLevel
		log.Fatalf("Error parsing log level: %v", err)
	}
	conf.Log.Path = viper.GetString("LOG_PATH")
	conf.Log.Filename = viper.GetString("LOG_FILENAME")
	conf.Log.MaxSize = viper.GetInt("LOG_MAX_SIZE_MB")
	conf.Log.MaxBackup = viper.GetInt("LOG_MAX_BACKUP")
	conf.Log.MaxAge = viper.GetInt("LOG_MAX_AGE")
	conf.Log.ConsoleEnabled = viper.GetBool("LOG_CONSOLE_ENABLED")
	conf.Log.FileEnabled = viper.GetBool("LOG_FILE_ENABLED")
	conf.Log.FileRotate = viper.GetBool("LOG_FILE_ROTATE")

	conf.Lang.DefaultLanguage = viper.GetString("DEFAULT_LANG")
	conf.Lang.PathLanguage = viper.GetString("PATH_LANG")
	viper.UnmarshalKey("ACCEPT_LANG", &conf.Lang.AcceptedLanguages)
	// conf.Lang.AcceptedLanguages = viper.GetStringSlice("ACCEPT_LANG")

	conf.Encryption.AccessTokenPassphrase = viper.GetString("ACCESS_TOKEN_PASSPHRASE")
	conf.Encryption.RefreshTokenPassphrase = viper.GetString("REFRESH_TOKEN_PASSPHRASE")
	conf.Encryption.AdminTokenPassphrase = viper.GetString("ADMIN_TOKEN_PASSPHRASE")
	conf.Service = viper.GetString("SERVICE_URL")
	conf.Webhook = viper.GetString("WEBHOOK_URL")
	conf.ApiKey.Pastebin = viper.GetString("API_KEY_PASTEBIN")

	log.Println("Config initialized successfully")
}

func Get() *Config {
	return conf
}
