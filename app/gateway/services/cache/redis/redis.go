package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/sptGabriel/starwars/app/gateway/services/cache"
)

const defaultExpireTime = 44000

var _ cache.Cache = Redis{}

type Redis struct {
	pool       *redis.Pool
	expireTime int
}

func New(pool *redis.Pool, expireTime int) Redis {
	redis := Redis{
		pool:       pool,
		expireTime: expireTime,
	}

	if redis.expireTime == 0 {
		redis.expireTime = defaultExpireTime
	}

	return redis
}

func (r Redis) Get(key string) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()

	value, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	return value, nil
}

func (r Redis) Save(key string, value []byte) error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value, "EX", r.expireTime, "NX")
	if err != nil {
		return err
	}

	return nil
}
