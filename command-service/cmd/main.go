package main

import (
	"command-service/config"
	userS "command-service/internal/service"
	"command-service/logger"
	mongo "command-service/pkg/mongosh"
	gRPC "command-service/pkg/service"
	p "command-service/producer"
	"log"
	"time"
)

func main() {

	logger.InitLog()
	logger.Info("Starting command-service...")

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

	pro, err := p.DialProducer()
	if err != nil {
		logger.Error("RabbitMQ connection FAILED", err)
		log.Fatal(err)
	}
	
	logger.Info("Dial with Producer is succesfully")
	defer pro.Connection.Close()
	defer pro.Channel.Close()

	service := userS.NewCommandService(mongoDB, pro)
	logger.Info("Command-Service working clean")

	ready := gRPC.NewService(service)
	if err := ready.RUN(*cfg); err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
}
