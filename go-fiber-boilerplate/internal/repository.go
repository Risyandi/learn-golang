package internal

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/database"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/internal/repository"
)

type RepositoryManager struct {
	auth    repository.AuthRepository
	session repository.SessionRepository
	product repository.ProductRepository
}

func NewRepositoryManager(pg *sql.DB, mongo *mongo.Database, redis database.RedisInstance) RepositoryManager {

	return RepositoryManager{
		auth:    repository.NewAuthRepository(pg),
		session: repository.NewSessionRepository(mongo),
		product: repository.NewProductRepository(pg),
	}
}
