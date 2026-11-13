package utils

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// NewLogger creates and configures a logrus.Logger based on the provided level string.
func NewLogger(level string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
// kept it simple here

	lvl, err := logrus.ParseLevel(strings.ToLower(level))
	if err != nil {
		logger.SetLevel(logrus.InfoLevel)
	} else {
		logger.SetLevel(lvl)
// minor polish
	}
	return logger
}