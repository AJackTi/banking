package main

import (
	"github.com/AJackTi/banking/app"
	"github.com/AJackTi/banking/logger"
)

func main() {
	logger.Info("Starting our application...")
	app.Start()
}
