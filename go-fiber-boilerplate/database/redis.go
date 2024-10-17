package database

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"gitlab.com/sugaanaluam/gofiber-boilerplate/config"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/logger"
	"gitlab.com/sugaanaluam/gofiber-boilerplate/pkg/utils"
)

var (
	redisConn *redis.Client
	redisLog  zerolog.Logger
)

type (
	redisInstance struct {
		conn *redis.Client
		conf *config.Config
		ctx  context.Context
	}
	RedisInstance interface {
		Ping() error
		Get(key string, dest interface{}) error
		GetString(key string) (string, error)
		Set(key string, val interface{}) error
		Setx(key string, val interface{}) error
		Setxc(key string, expire time.Duration, val interface{}) error
		HashSet(key string, val interface{}) error
		HashSetField(key string, field string, val interface{}) error
		HashGetAll(key string) (map[string]string, error)
		HashGet(key string, field string) (string, error)
		Del(key string) error
		DelKeysByPatern(patern string) error
		Keys(patern string) ([]string, error)
		Duration(key string) (*time.Duration, error)
		Close() error
	}
)

func RedisInit(ctx context.Context, db int) RedisInstance {
	redisLog = logger.Get("redis")
	conf := config.Get()

	redisLog.Info().Msg(fmt.Sprintf("Connect redis client DB[%d]...", db))
	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", conf.Database.Redis.Host, conf.Database.Redis.Port),
		Password:    conf.Database.Redis.Auth,
		DB:          db,
		DialTimeout: conf.Database.Redis.DialTimeout,
	})

	redisConn = new(redis.Client)
	redisConn = rdb

	if !RedisIsConnected() {
		redisLog.Fatal().Int("DB number", db).Msg("Redis not connected")
	}

	redisLog.Info().Msg(fmt.Sprintf("Redis connected DB[%d]...", db))
	return &redisInstance{
		conn: redisConn,
		conf: conf,
		ctx:  ctx,
	}
}

func RedisIsConnected() bool {
	ctx := context.Background()
	_, err := redisConn.Ping(ctx).Result()
	if err != nil {
		redisLog.Err(err).Msg("Redis health check problem")
		return false
	}
	return true
}

func (r *redisInstance) Ping() error {
	_, err := r.conn.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *redisInstance) Get(key string, dest interface{}) error {
	err := utils.MustBePointer(dest, "dest")
	if err != nil {
		return err
	}
	err = r.Ping()
	if err != nil {
		return err
	}
	val, err := r.conn.Get(r.ctx, key).Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(val, dest)
	if err != nil {
		return err
	}
	return nil
}

func (r *redisInstance) GetString(key string) (string, error) {
	err := r.Ping()
	if err != nil {
		return "", err
	}
	val, err := r.conn.Get(r.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisInstance) Set(key string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.Set(r.ctx, key, val, 0).Err()
}

func (r *redisInstance) Setx(key string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.Set(r.ctx, key, val, r.conf.Database.Redis.ExpiredDuration).Err()
}

func (r *redisInstance) Setxc(key string, expire time.Duration, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.Set(r.ctx, key, val, expire).Err()
}

func (r *redisInstance) HashSet(key string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.HSet(r.ctx, key, val).Err()
}

func (r *redisInstance) HashSetField(key string, field string, val interface{}) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.HSet(r.ctx, key, field, val).Err()
}

func (r *redisInstance) HashGetAll(key string) (map[string]string, error) {
	err := r.Ping()
	if err != nil {
		return nil, err
	}
	data, err := r.conn.HGetAll(r.ctx, key).Result()
	return data, err
}

func (r *redisInstance) HashGet(key string, field string) (string, error) {
	err := r.Ping()
	if err != nil {
		return "", err
	}
	data, err := r.conn.HGet(r.ctx, key, field).Result()
	return data, err
}

func (r *redisInstance) Del(key string) error {
	err := r.Ping()
	if err != nil {
		return err
	}
	return r.conn.Del(r.ctx, key).Err()
}

func (r *redisInstance) DelKeysByPatern(patern string) error {
	val, err := r.Keys(patern)
	if err != nil {
		return err
	}
	return r.conn.Del(r.ctx, val...).Err()
}

func (r *redisInstance) Keys(patern string) ([]string, error) {
	err := r.Ping()
	if err != nil {
		return nil, err
	}

	val, err := r.conn.Keys(r.ctx, patern).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *redisInstance) Duration(key string) (*time.Duration, error) {
	err := r.Ping()
	if err != nil {
		return nil, err
	}

	timeDuration, err := r.conn.TTL(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return &timeDuration, nil
}

func (r *redisInstance) Close() error {
	return r.conn.Close()
}
