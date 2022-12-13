package main

import (
	"main/database"
	"main/handlers"
	"main/logger"
	"main/scheduler"
)

func main() {
	logger.SetupLogger()
	database.SetupDatabase()
	scheduler.SetupScheduler()
	handlers.StartServer()
	database.Client.Close()
}
