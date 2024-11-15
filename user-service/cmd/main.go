package main

import (
	"log"
	"user-service/config"
	userS "user-service/internal/service"
	"user-service/logger"
	mongo "user-service/pkg/mongosh"
	rds "user-service/pkg/redis"
	gRPC "user-service/pkg/service"
)

func main() {
	logger.InitLog()
	logger.Info("Starting user-service...")

	cfg, err := config.Load(".")
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Configuration loaded")

	mongoDB, err := mongo.NewConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("MongoDB connected")

	redis, err := rds.NewClient(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Redis connected")

	service := userS.NewUserService(mongoDB, redis)
	logger.Info("User-service working clean")

	ready := gRPC.NewService(service)
	if err := ready.RUN(*cfg); err != nil {
		log.Fatal(err)
	}
}
