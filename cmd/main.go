package main

import (
	"github.com/rmansilla92/go-apps-template/api"
	"os"
	"strings"
)

const (
	defaultPort string = "8080"
)

func main() {
	port := os.Getenv("PORT")

	if strings.EqualFold(port, "") {
		port = defaultPort
	}

	api.Start(port)
}
