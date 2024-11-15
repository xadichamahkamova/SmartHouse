package main

import (
	"api-gateway/api"
	"api-gateway/config"
	_ "api-gateway/docs"
	"api-gateway/logger"
	c "api-gateway/pkg/command-service"
	d "api-gateway/pkg/device-service"
	u "api-gateway/pkg/user-service"
	"fmt"
	"log"
)

func main() {

	logger.InitLog()
	logger.Info("Api-gateway starting...")
	
	cfg, err := config.Load(".") 
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Configuration loaded")

	user, err := u.DialWithUserService(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Dial with user-service is ok")

	device, err := d.DialWithDeviceService(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Dial with device-service is ok")

	command, err := c.DialWithCommandService(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	logger.Info("Dial with command-service is ok")

	r := api.NewGin(user, device, command)
	addr := fmt.Sprintf(":%s", cfg.ApiGatewayPort)
	r.Run(addr)
}
