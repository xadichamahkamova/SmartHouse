package handler

import (
	pb "api-gateway/genproto/devicepb"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @Router /devices [post]
// @Summary CREATE DEVICE
// @Description This method creates device
// @Security BearerAuth
// @Tags DEVICE
// @Accept json
// @Produce json
// @Param device body models.Device true "Device"
// @Success 200 {object} models.ResponseOfDevice
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) CreateDevice(c *gin.Context) {
	req := pb.Device{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}

	resp, err := h.Device.Client.CreateDevice(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /devices/{id} [get]
// @Summary GET DEVICE
// @Description This method gets device
// @Security BearerAuth
// @Tags DEVICE
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Device
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetDevice(c *gin.Context) {
	req := pb.SingleID{}
	req.DeviceId = c.Param("id")

	resp, err := h.Device.Client.GetDevice(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /devices [get]
// @Summary GET DEVICES
// @Description This method gets devices
// @Security BearerAuth
// @Tags DEVICE
// @Accept json
// @Produce json
// @Success 200 {object} models.ListResponseOfDevice
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) ListOfDevice(c *gin.Context) {
	resp, err := h.Device.Client.ListOfDevice(ctx, &pb.ListRequestOfDevice{})
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /devices [put]
// @Summary UPDATE DEVICE
// @Description This method updates device
// @Security BearerAuth
// @Tags DEVICE
// @Accept json
// @Produce json
// @Param device body models.Device true "Device"
// @Success 200 {object} models.ResponseOfDevice
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateDevice(c *gin.Context) {
	req := pb.Device{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}

	resp, err := h.Device.Client.UpdateDevice(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /devices/{id} [delete]
// @Summary DELETE DEVICE
// @Description This method deletes device
// @Security BearerAuth
// @Tags DEVICE
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.ResponseOfDevice
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteDevice(c *gin.Context) {
	req := pb.SingleID{}
	req.DeviceId = c.Param("id")

	resp, err := h.Device.Client.DeleteDevice(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}
