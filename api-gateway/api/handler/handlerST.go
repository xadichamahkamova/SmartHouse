package handler

import (
	c "api-gateway/pkg/command-service"
	d "api-gateway/pkg/device-service"
	u "api-gateway/pkg/user-service"
	"context"
)

type HandlerST struct {
	User    u.UserService
	Device  d.DeviceService
	Command c.CommandService
}

var ctx = context.Background()