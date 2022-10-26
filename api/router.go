package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rmansilla92/go-apps-template/configs"
	"github.com/rmansilla92/go-apps-template/internal/controller"
	"os"
)

func routes(controllers controller.Controllers) *gin.Engine {
	router := gin.New()

	configs.Scope = os.Getenv("SCOPE")

	mapRoutes(router, controllers)

	return router
}

func mapRoutes(router *gin.Engine, controllers controller.Controllers) {
	router.GET("/ping", controllers.Ping)
}
