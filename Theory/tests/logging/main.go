package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	log.SetLevel(logrus.DebugLevel)

	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("Theory/tests/logging/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	log.Info("Successfully logged in")

	log.WithFields(logrus.Fields{
		"username": "Tim",
		"id":       17,
	}).Info("Logged in user")

	errStr := "Error with server"
	log.WithFields(logrus.Fields{
		"event": "error_with_server",
		"error": errStr,
	}).Error("Error connecting with server")
}
