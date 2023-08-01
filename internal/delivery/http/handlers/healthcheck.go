package handlers

import "github.com/gin-gonic/gin"

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Check(c *gin.Context) {
	c.JSON(200, "UP")
}
