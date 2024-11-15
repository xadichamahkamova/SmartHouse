package redis

import (
	"context"
	"fmt"
	"command-service/config"

	"github.com/redis/go-redis/v9"
)
type Redis struct{
	Client redis.Client
}
func NewClient(cfg config.Config) (*Redis, error){

	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		Password: "",
		DB: 0,
	})
	if cmd := rdb.Ping(context.Background()); cmd != nil {
		return nil, cmd.Err()
	}
	return &Redis{
		Client: *rdb,
	}, nil
}