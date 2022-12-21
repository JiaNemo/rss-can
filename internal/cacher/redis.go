package cacher

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/soulteary/RSS-Can/internal/define"
	"github.com/soulteary/RSS-Can/internal/logger"
)

var instanceRedis *redis.Client
var ctx = context.Background()

const REDIS_KEY_NOT_EXIST = redis.Nil

func init() {
	if define.REDIS {
		connect(true)
	}
}

func connect(init bool) *redis.Client {
	if !init {
		err := instanceRedis.Ping(ctx).Err()
		if err == nil {
			return instanceRedis
		}
	}

	if !init {
		instanceRedis.Close()
	}

	instanceRedis = redis.NewClient(&redis.Options{
		Addr:     define.REDIS_SERVER,
		Password: define.REDIS_PASS,
		DB:       define.REDIS_DB,
		PoolSize: 100,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			logger.Instance.Info("Restore the connection to Redis.")
			return nil
		},
		MaxRetries: 3,
	})

	return instanceRedis
}

func Disconnect() (err error) {
	return instanceRedis.Close()
}

func UpdateDataToRedis(key string, value string) (err error) {
	rdb := connect(false)
	err = rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetDataFromRedis(key string) (result string, err error) {
	rdb := connect(false)
	data, err := rdb.Get(ctx, key).Result()
	if err == REDIS_KEY_NOT_EXIST {
		return "", nil
	} else if err != nil {
		return "", err
	} else {
		return data, nil
	}
}

func DelDataByKeyFromRedis(key string) (err error) {
	rdb := connect(false)
	_, err = rdb.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func SetDataExpireByKeyFromRedis(key string, expire time.Duration) (err error) {
	rdb := connect(false)
	_, err = rdb.Expire(ctx, key, expire).Result()
	if err != nil {
		return err
	}
	return nil
}
