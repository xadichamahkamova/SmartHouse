package handler

import (
	pb "api-gateway/genproto/userpb"
	"api-gateway/models"

	"github.com/gin-gonic/gin"
)

// @Router /users/register [post]
// @Summary CREATE USER
// @Description This method creates users
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) Register(c *gin.Context) {
	req := pb.User{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}

	resp, err := h.User.Client.Register(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /users/login [post]
// @Summary GET USER
// @Description This method gets user
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param user body models.LoginRequest true "LoginRequest"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) Login(c *gin.Context) {
	req := pb.LoginRequest{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}
	resp, err := h.User.Client.Login(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /users/profile [get]
// @Summary GET USERS
// @Description This method gets users
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Success 200 {object} models.ListResponse 
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) ListOfUser(c *gin.Context) {
	resp, err := h.User.Client.ListOfUser(ctx, &pb.ListRequest{})
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /users [put]
// @Summary UPDATE USER
// @Description This method updates user
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 200 {object} models.UniverResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) UpdateUser(c *gin.Context) {
	req := pb.User{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, models.Message{Err: err.Error()})
		return
	}

	resp, err := h.User.Client.UpdateUser(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}

// @Router /users/{id} [delete]
// @Summary DELETE USER
// @Description This method deletes user
// @Security BearerAuth
// @Tags USER
// @Accept json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} models.UniverResponse
// @Failure 400 {object} models.Message
// @Failure 500 {object} models.Message
func (h *HandlerST) DeleteUser(c *gin.Context) {
	req := pb.DeleteRequest{}
	req.Id = c.Param("id")
	resp, err := h.User.Client.DeleteUser(ctx, &req)
	if err != nil {
		c.JSON(500, models.Message{Err: err.Error()})
		return
	}
	c.JSON(200, resp)
}
