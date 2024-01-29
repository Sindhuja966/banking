package main

import (
	"github.com/Sindhuja966/banking/app"
	"github.com/Sindhuja966/banking/logger"
)

func main() {

	logger.Info("Starting out application")
	app.Start()
}
