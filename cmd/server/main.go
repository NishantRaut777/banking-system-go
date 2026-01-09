package main

import (
	_ "github.com/NishantRaut777/banking-system-go/docs" // swagger docs
	"github.com/NishantRaut777/banking-system-go/internal/server"
)

// @title Banking API
// @version 1.0
// @description Banking system REST API built with Go & Gin

// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	server.Start()
}
