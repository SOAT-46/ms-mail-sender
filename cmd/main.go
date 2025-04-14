package main

import (
  logger "github.com/sirupsen/logrus"
  "os"
  "os/signal"
  "syscall"
)

func main() {
  logger.Infof("Starting ms-mail-sender")

  signalChan := make(chan os.Signal, 1)
  signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
  defer handlePanic()

  <-signalChan

  logger.Infof("Shutting down ms-mail-sender")
}

func handlePanic() {
  if r := recover(); r != nil {
    logger.WithField("panic", r).Error("Panic occurred in the application.")
  }
}
