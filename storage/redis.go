package storage

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
)

type redisDB struct {
	pool *redis.Pool
}

//NewRedis make a new Redis database
func NewRedis() Database {
	var p = newPool()
	return &redisDB{pool: p}
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func (r *redisDB) getid() string {
	c := r.pool.Get()
	defer c.Close()
	value, _ := redis.Int(c.Do("INCR", "IDCOUNT"))
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(value)))
}

//Get is used to return the key from Redis
func (r *redisDB) Get(key string) (status, error) {
	c := r.pool.Get()
	defer c.Close()
	count, err := redis.Int(c.Do("GET", fmt.Sprintf("%q_COUNT", key)))
	if err != nil {
		count = 0
	}

	value, err := redis.String(c.Do("GET", key))
	if err != nil {
		return status{}, ErrNotFound
	}

	return status{key, value, count}, nil
}

//Set is used to set the url into Redis
func (r *redisDB) Set(url string) (status, error) {
	c := r.pool.Get()
	defer c.Close()

	id := r.getid()

	c.Do("SET", id, url)
	return status{id, url, 0}, nil
}

// Incr is used to Increment the count in Redis
func (r *redisDB) Incr(key string) {
	c := r.pool.Get()
	defer c.Close()
	c.Do("INCR", fmt.Sprintf("%q_COUNT", key))
}
