package main

import (
	"os"
	"os/signal"
	"syscall"

	logger "github.com/sirupsen/logrus"
)

func main() {
	logger.Infof("Starting ms-mail-sender")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	app := injectApps()
	defer handlePanic()
	RunModules(app)
	<-signalChan

	logger.Infof("Shutting down ms-mail-sender")
}

func handlePanic() {
	if r := recover(); r != nil {
		logger.WithField("panic", r).Error("Panic occurred in the application.")
	}
}
