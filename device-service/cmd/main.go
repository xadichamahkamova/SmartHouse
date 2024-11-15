package main

import (
	"device-service/config"
	"device-service/consumer"
	userS "device-service/internal/service"
	"device-service/logger"
	mongo "device-service/pkg/mongosh"
	gRPC "device-service/pkg/service"
	"log"
	"sync"
	"time"
)

func main() {

	logger.InitLog()
	logger.Info("Starting device-service...")

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

	service := userS.NewDeviceService(mongoDB)
	logger.Info("Device-service working clean")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		consumer.StartConsuming()
	}()
	logger.Info("Starting Consuming...")
	
	ready := gRPC.NewService(service)
	go func() {
		defer wg.Done()
		if err := ready.RUN(*cfg); err != nil {
			log.Fatal(err)
		}
	}()

	time.Sleep(5 * time.Second)

	wg.Wait()
}
