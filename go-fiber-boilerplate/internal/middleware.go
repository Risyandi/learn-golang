package internal

import (
	"gitlab.com/sugaanaluam/gofiber-boilerplate/database"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/middleware"
)

type MiddlewareManager struct {
	auth    middleware.AuthMiddleware
	session middleware.SessionMiddleware
}

func NewMiddlewareManager(repo RepositoryManager, redis database.RedisInstance) MiddlewareManager {
	return MiddlewareManager{
		auth:    middleware.NewAuthMiddleware(repo.auth, redis), // Example middleware module
		session: middleware.NewSessionMiddleware(repo.session),
	}
}
