package main

import (
	"github.com/AJackTi/banking-lib/logger"
	"github.com/AJackTi/banking/app"
)

func main() {
	logger.Info("Starting our application...")
	app.Start()
}
