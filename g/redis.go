package g

import (
	// "github.com/gomodule/redigo/redis"
	"github.com/go-redis/redis"
	// "github.com/garyburd/redigo/redis"
)

// var RedisConnPool *redis.Pool
var RedisClient *redis.Client

func InitRedisConnPool() {
	redisConfig := Config().Redis
	redisSentinelConfig := Config().RedisSentinel

	if redisConfig.EnableSentinel {
		RedisClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    redisSentinelConfig.MasterName,
			SentinelAddrs: redisSentinelConfig.SentinelAddrs,
			DB:            redisSentinelConfig.Db,
		})
	} else {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     redisConfig.Addr,
			Password: redisConfig.Passwd, // no password set
			DB:       0,                  // use default DB
		})
	}

	// RedisClient = redis.NewFailoverClient(&redis.FailoverOptions{
	// 	MasterName:    "master1",
	// 	SentinelAddrs: []string{"172.20.207.60:26379", "172.20.207.114:26379"},
	// 	DB:            63,
	// })
	// RedisClient.Ping()
	// RedisConnPool = &redis.Pool{
	// 	MaxIdle:     redisConfig.MaxIdle,
	// 	IdleTimeout: 240 * time.Second,
	// 	Dial: func() (redis.Conn, error) {
	// 		option := redis.DialPassword(redisConfig.Passwd)
	// 		c, err := redis.Dial("tcp", redisConfig.Addr, option)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		return c, err
	// 	},
	// 	TestOnBorrow: PingRedis,
	// }

}

// func PingRedis(c redis.Conn, t time.Time) error {
// 	_, err := c.Do("ping")
// 	if err != nil {
// 		log.Println("[ERROR] ping redis fail", err)
// 	}
// 	return err
// }
