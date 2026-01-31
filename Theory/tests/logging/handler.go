package main

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{}

	correlationID := uuid.New().String()
	log.AddHook(NewMyLogstashHook())
	log.WithFields(logrus.Fields{
		"correlation_id": correlationID,
		"event":          "request_started",
	}).Info("Starting Request")

	// business logic

	log.WithFields(logrus.Fields{
		"correlation_id": correlationID,
		"event":          "request_finished",
	}).Info("Finished Request")
}

type MyLogstashHook struct {
	logrus.Hook
}

func NewMyLogstashHook() logrus.Hook {
	return &MyLogstashHook{}
}

func (h *MyLogstashHook) Fire(entry *logrus.Entry) error {
	return nil
}

func (h *MyLogstashHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
