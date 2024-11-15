package handler

import (
	pb "api-gateway/genproto/commandpb"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @Router /control [post]
// @Summary POST CONTROL
// @Description This method creates command
// @Security BearerAuth
// @Tags COMMAND
// @Accept json
// @Produce json
// @Param command body models.Command true "Command"
// @Success 200 {object} models.ResponseOfCommand
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) CreateCommand(c *gin.Context) {
	req := pb.Command{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}

	resp, err := h.Command.Client.CreateCommand(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /command/{id} [get]
// @Summary GET COMMAND
// @Description This method gets command
// @Security BearerAuth
// @Tags COMMAND
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.Command
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) GetCommand(c *gin.Context) {
	req := pb.SingleId{}
	req.CommandId = c.Param("id")

	resp, err := h.Command.Client.GetCommand(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /commands [get]
// @Summary GET COMMANDS
// @Description This method gets commands
// @Security BearerAuth
// @Tags COMMAND
// @Accept json
// @Produce json
// @Success 200 {object} models.ListResponseOfCommand
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) ListOfCommand(c *gin.Context) {
	resp, err := h.Command.Client.ListOfCommand(ctx, &pb.ListRequestOfCommand{})
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /command [put]
// @Summary UPDATE COMMAND
// @Description This method updates command
// @Security BearerAuth
// @Tags COMMAND
// @Accept json
// @Produce json
// @Param command body models.Command true "Command"
// @Success 200 {object} models.ResponseOfCommand
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateCommand(c *gin.Context) {
	req := pb.Command{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}

	resp, err := h.Command.Client.UpdateCommand(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /command/{id} [delete]
// @Summary DELETE COMMAND
// @Description This method deletes command
// @Security BearerAuth
// @Tags COMMAND
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.ResponseOfCommand
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteCommand(c *gin.Context) {
	req := pb.SingleId{}
	req.CommandId = c.Param("id")

	resp, err := h.Command.Client.DeleteCommand(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}
