package govideo

import (
	"fmt"
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/sickyoon/govideo/govideo/models"
)

// RedisClient -
type RedisClient struct {
	*redis.Pool
	secret     string
	userExpiry int
}

// NewRedisClient -
func NewRedisClient(config *models.Config) (*RedisClient, error) {
	secret, err := GenerateKey()
	if err != nil {
		return nil, err
	}
	redisClient := &RedisClient{
		Pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", config.Cache.URI)
				if err != nil {
					return nil, err
				}
				if config.Cache.Password != "" {
					if _, err := c.Do("AUTH", config.Cache.Password); err != nil {
						c.Close()
						return nil, err
					}
				}
				if config.Cache.Database != "" {
					if _, err := c.Do("SELECT", config.Cache.Database); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, nil
			},
		},
		secret:     secret,
		userExpiry: config.App.UserExpiry,
	}
	err = redisClient.Pool.Get().Err()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %v", err)
	}
	return redisClient, nil
}

// SetAuthCache sets user data in redis cache
func (rc *RedisClient) SetAuthCache(userID string, data []byte) ([]byte, error) {
	conn := rc.Get()
	defer conn.Close()
	key := []byte(rc.secret + ":user:" + userID)
	_, err := conn.Do("SETEX", key, strconv.Itoa(rc.userExpiry), data)
	return key, err
}

// GetAuthCache gets user data from redis cache
func (rc *RedisClient) GetAuthCache(key []byte) ([]byte, error) {
	conn := rc.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

// ClearAuthCache clears user data from redis cache
func (rc *RedisClient) ClearAuthCache(key []byte) error {
	conn := rc.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	return err
}

// GetEncodedPath returns encoded path from real file path
func (rc *RedisClient) GetEncodedPath(path, ipAddr string) (string, error) {
	conn := rc.Get()
	defer conn.Close()
	encodedPath, err := GenerateKey()
	if err != nil {
		return "", err
	}
	key := []byte(rc.secret + ":encoded:" + ipAddr + ":" + encodedPath)
	_, err = conn.Do("SETEX", key, strconv.Itoa(rc.userExpiry), path)
	return encodedPath, err
}

// GetRealPath returns real file path from encoded path
func (rc *RedisClient) GetRealPath(encodedPath, ipAddr string) ([]byte, error) {
	conn := rc.Get()
	defer conn.Close()
	key := []byte(rc.secret + ":encoded:" + ipAddr + ":" + encodedPath)
	return redis.Bytes(conn.Do("GET", key))
}
