package main

import (
	"github.com/diegobbrito/gopportunities/config"
	"github.com/diegobbrito/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main") // Initialize the logger
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}
	router.Initialize() // Initialize the router
}
