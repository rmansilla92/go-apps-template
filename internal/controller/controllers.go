package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rmansilla92/go-apps-template/internal/service"
)

type Controllers interface {
	Ping(c *gin.Context)
}

type controllers struct {
	service service.Services
}

func NewControllers(services service.Services) Controllers {
	return &controllers{
		service: services,
	}
}
