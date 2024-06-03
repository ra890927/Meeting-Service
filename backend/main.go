package main

import (
	"os"
	"os/signal"
	"syscall"

	"meeting-center/src/clients"
	"meeting-center/src/router"
	"meeting-center/src/utils"
)

func init() {
	utils.InitConfig()
	clients.InitDB()
}

func main() {
	router.HttpServerRun()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
