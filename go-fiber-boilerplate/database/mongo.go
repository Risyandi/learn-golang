package database

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

var (
	mongoConn *mongo.Database
	mongoLog  zerolog.Logger
)

func MongoInit(conf *config.Config) {
	mongoLog = logger.Get("mongodb")

	mongoLog.Info().Msg("Connecting to MongoDB server...")
	var URI string
	if conf.Database.Mongo.SRV {
		URI = fmt.Sprintf("mongodb+srv://%s:%s@%s", conf.Database.Mongo.Username, conf.Database.Mongo.Password, conf.Database.Mongo.Host)
	} else {
		URI = fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Database.Mongo.Username, conf.Database.Mongo.Password, conf.Database.Mongo.Host, conf.Database.Mongo.Port)
	}

	clientOptions := options.Client().ApplyURI(URI)
	if conf.Env == "production" {
		clientOptions.SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
	}

	if conf.Log.Level == zerolog.DebugLevel {
		cmdMonitor := &event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				mongoLog.Debug().Str("event", evt.CommandName).Msgf("%+v", evt)
			},
		}
		clientOptions.SetMonitor(cmdMonitor)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		mongoLog.Fatal().Err(err).Msg("MongoDB connection problem")
	}

	mongoConn = client.Database(conf.Database.Mongo.Name)

	if !MongoIsConnected() {
		mongoLog.Fatal().Str("connection", URI).Msg("MongoDB not connected")
	}

	mongoLog.Info().Msg("MongoDB connected")

}

func GetMongoConnection() *mongo.Database {
	return mongoConn
}

func MongoIsConnected() bool {
	if mongoConn == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := mongoConn.Client().Ping(ctx, nil)
	if err != nil {
		mongoLog.Err(err).Msg("MongoDB health check problem")
		return false
	}

	return true
}
