package api

import (
	"api-gateway/api/handler"
	"api-gateway/api/handler/middleware"
	c "api-gateway/pkg/command-service"
	d "api-gateway/pkg/device-service"
	u "api-gateway/pkg/user-service"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @tite Api-gateway service
// @version 1.0
// @description Api-gateway service
// @host localhost:9000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authourization
func NewGin(user *u.UserService, device *d.DeviceService, command *c.CommandService) *gin.Engine {

	r := gin.Default()

	handler := handler.HandlerST{
		User:    *user,
		Device:  *device,
		Command: *command,
	}
	
	r.Use(middleware.Logger())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
		userGroup.GET("/profile", handler.ListOfUser)
		userGroup.PUT("/", handler.UpdateUser)
		userGroup.DELETE("/:id", handler.DeleteUser)

	}

	r.POST("/devices", handler.CreateDevice)
	r.GET("/devices/:id", handler.GetDevice)
	r.GET("/devices", handler.ListOfDevice)
	r.PUT("/devices", handler.UpdateDevice)
	r.DELETE("/devices/:id", handler.DeleteDevice)

	r.POST("/control", handler.CreateCommand)
	r.GET("/command/:id", handler.GetCommand)
	r.GET("/commands", handler.ListOfCommand)
	r.PUT("/command", handler.UpdateCommand)
	r.DELETE("/command/:id", handler.DeleteCommand)

	return r
}
