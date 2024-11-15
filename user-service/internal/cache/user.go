package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	pb "user-service/genproto/userpb"
	tok "user-service/internal/cache/token"
	"user-service/logger"
	rds "user-service/pkg/redis"
)

type RedisMethod struct {
	RDS rds.Redis
}

func NewRedis(redis *rds.Redis) *RedisMethod {
	return &RedisMethod{*redis}
}

func (r *RedisMethod) HoldOnUser(user *pb.Response) error {

	mapUser, err := tok.ParseToken(user.Token.AccessToken)
	if err != nil {
		logger.Error("Failed to parse token", map[string]interface{}{
			"error": err,
		})
		return err
	}
	unique := mapUser["id"]

	timeout, ok := mapUser["exp"].(float64)
	if !ok {
		logger.Error("Invalid expiration time", map[string]interface{}{
			"error": err,
		})
		return errors.New("invalid expiration time")
	}

	expirationTime := time.Unix(int64(timeout), 0)
	duration := time.Until(expirationTime)
	redisDataName := fmt.Sprintf("%sUser", unique)

	logger.Info("Redis data name and duration", map[string]interface{}{
		"data_name": redisDataName,
		"duration":  duration,
	})

	ctx := context.Background()
	jsonToken, err := json.Marshal(user.Token)
	if err != nil {
		logger.Error("Failed to marshal token", map[string]interface{}{
			"error": err,
		})
		return err
	}

	err = r.RDS.Client.Set(ctx, redisDataName, jsonToken, duration).Err()
	if err != nil {
		logger.Error("Failed to set redis data", map[string]interface{}{
			"error": err,
		})
		return err
	}

	logger.Info("Set User to Redis is successfully")
	return nil
}
