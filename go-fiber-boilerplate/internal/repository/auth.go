package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/database/query/auth"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

type AuthRepository interface {
	GetUserStatusRepository(ctx context.Context, userID uuid.UUID) (auth.UserStatusEntity, error)
}

type authRepository struct {
	db        *sql.DB
	exception exception.Exception
	log       zerolog.Logger
}

func NewAuthRepository(pg *sql.DB) AuthRepository {
	return &authRepository{
		db:        pg,
		exception: exception.NewException("auth-repository"),
		log:       logger.Get("auth-repository"),
	}
}

func (r authRepository) GetUserStatusRepository(ctx context.Context, id uuid.UUID) (auth.UserStatusEntity, error) {
	queryStr := auth.CheckUserStatusSQL

	var res auth.UserStatusEntity

	err := r.db.QueryRowContext(ctx, queryStr, id).Scan(
		&res.UserID,
		&res.IsActive,
		&res.IsSuspended,
		&res.IsKyc,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			r.exception.ErrorWithoutNoSqlResult(err)
			return res, err
		}
		return res, err
	}

	return res, nil
}
