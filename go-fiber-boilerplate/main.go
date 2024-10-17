package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/contrib/fiberi18n/v2"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/database"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/middleware"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

func init() {
	config.Init()
	conf := config.Get()
	logger.InitLogger(conf)
	database.PgInit(conf)
	database.MongoInit(conf)

}

func main() {
	conf := config.Get()
	rawLogs := logger.GetWithoutCaller(conf.API.Name)
	logs := rawLogs.With().Caller().Logger()
	ctx := context.Background()
	pg := database.GetPgConnection()
	mongo := database.GetMongoConnection()
	redis := database.RedisInit(ctx, conf.Database.Redis.Db)
	redisAuth := database.RedisInit(ctx, conf.Database.Redis.DbAuth)

	defer func() {
		err := pg.Close()
		if err != nil {
			logs.Err(err).Msg("Error close Postgres connection")
		}

		err = mongo.Client().Disconnect(ctx)
		if err != nil {
			logs.Err(err).Msg("Error close MongoDB connection")
		}

		err = redis.Close()
		if err != nil {
			logs.Err(err).Msg("Error close redis connection")
		}

		err = redisAuth.Close()
		if err != nil {
			logs.Err(err).Msg("Error close redis connection")
		}
	}()

	app := fiber.New(fiber.Config{
		AppName:      conf.API.Name,
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(fiberzerolog.New(config.FiberZerolog(rawLogs, logs)))
	app.Use(recover.New(config.FiberRecover()))
	app.Use(fiberi18n.New(config.I18n()))
	app.Use(requestid.New())
	app.Use(compress.New())
	app.Use(helmet.New())
	prom := middleware.NewPrometheusMiddleware()
	prom.RecordMetrics()
	app.Use(prom.MonitoringMiddleware())
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	repository := internal.NewRepositoryManager(pg, mongo, redis)
	service := internal.NewServiceManager(repository)
	handler := internal.NewHandlerManager(service)
	middleware := internal.NewMiddlewareManager(repository, redisAuth)
	route := internal.NewRouterManager(app, handler, middleware)

	// Initialize list routes
	route.Init()

	baseUrl := fmt.Sprintf(":%d", conf.API.Port)
	go func() {
		err := app.Listen(baseUrl)
		if err != nil {
			logs.Fatal().Err(err).Msg("Error app serve")
		}
	}()

	handleShutdown(logs)
}

func handleShutdown(logs zerolog.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logs.Warn().Msg("Shutting down server...")

	logs.Info().Msg("Exiting server")
}
