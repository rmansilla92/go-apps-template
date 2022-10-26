package api

import (
	"fmt"
	"github.com/rmansilla92/go-apps-template/internal/controller"
	"github.com/rmansilla92/go-apps-template/internal/service"
)

func Start(port string) {
	if err := run(port); err != nil {
		fmt.Println("error running server", err) //TODO: Change to logger
	}
}

func run(port string) error {
	services := service.NewServices()

	controllers := controller.NewControllers(services)

	router := routes(controllers)

	return router.Run(":" + port)
}
