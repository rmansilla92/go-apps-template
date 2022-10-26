package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (controller *controllers) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
