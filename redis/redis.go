package redis

import (
	"fmt"
	redis_cli "github.com/go-redis/redis/v7"
)

type redis struct {
	config Config
	conn   *redis_cli.Client
}

func New(config Config) *redis {
	return &redis{
		config: config,
	}
}

func (r *redis) Init() error {
	dns := fmt.Sprintf("%s:%s", r.config.Host, r.config.Port)

	conn := redis_cli.NewClient(&redis_cli.Options{
		Addr: dns,
	})

	if _, err := conn.Ping().Result(); err != nil {
		panic(err)
	}

	r.conn = conn
	return nil
}

func (r *redis) Conn() *redis_cli.Client {
	return r.conn
}
