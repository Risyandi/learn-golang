package repository

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/database/query/session"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/exception"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
)

type (
	SessionRepository interface {
		NewTransaction(ctx context.Context) (SessionRepositoryTx, error)
		GetActiveSession(ctx context.Context, userID string, day time.Time) (string, error)
	}
	sessionRepository struct {
		db        *mongo.Database
		exception exception.Exception
		log       zerolog.Logger
	}
	SessionRepositoryTx interface {
		Commit() error
		InsertSession(sess *session.UserSessionEntity) error
		InsertLog(col string, logs interface{}) error
		DisableExpiredSession(userID string, day time.Time) (int64, error)
	}
	sessionRepositoryTx struct {
		ctx     context.Context
		db      *mongo.Database
		session mongo.Session
	}
)

func NewSessionRepository(mongo *mongo.Database) SessionRepository {
	return &sessionRepository{
		db:        mongo,
		exception: exception.NewException("session-repository"),
		log:       logger.Get("session-repository"),
	}
}

func (r *sessionRepository) NewTransaction(ctx context.Context) (SessionRepositoryTx, error) {

	session, err := r.db.Client().StartSession()
	if err != nil {
		return nil, err
	}

	err = session.StartTransaction()
	if err != nil {
		return nil, err
	}

	return &sessionRepositoryTx{ctx, r.db, session}, nil
}

func (r *sessionRepositoryTx) Commit() error {
	if err := r.session.CommitTransaction(r.ctx); err != nil {
		return err
	}
	r.session.EndSession(r.ctx)
	return nil
}

func (r *sessionRepository) GetActiveSession(ctx context.Context, userID string, day time.Time) (string, error) {
	var res session.UserSessionEntity
	coll := r.db.Collection("user_sessions")

	opts := options.FindOne().SetProjection(bson.M{"session_id": 1})
	err := coll.FindOne(
		ctx,
		bson.M{
			"user_id":   userID,
			"is_active": true,
			"created_at": bson.M{
				"$gte": day,
			},
		},
		opts).Decode(&res)
	if err != nil {
		return res.SessionID, err
	}
	return res.SessionID, nil
}

func (r sessionRepositoryTx) DisableExpiredSession(userID string, day time.Time) (int64, error) {
	coll := r.db.Collection("user_sessions")

	result, err := coll.UpdateMany(
		r.ctx,
		bson.M{
			"user_id":   userID,
			"is_active": true,
			"created_at": bson.M{
				"$lte": day.Unix(),
			},
		},
		bson.M{
			"$set": bson.M{
				"is_active": false,
			}},
	)
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil

}

func (r *sessionRepositoryTx) InsertSession(sess *session.UserSessionEntity) error {

	coll := r.db.Collection("user_sessions")

	_, err := coll.InsertOne(r.ctx, sess)
	if err != nil {
		return err
	}
	return nil
}

func (r *sessionRepositoryTx) InsertLog(col string, logs interface{}) error {
	coll := r.db.Collection(col)

	_, err := coll.InsertOne(r.ctx, logs)
	if err != nil {
		return err
	}
	return nil
}
