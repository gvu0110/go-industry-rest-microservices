package log_logrus

import (
	"fmt"
	"go-industry-rest-microservices/src/api/config"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)

func init() {
	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		//Log.Formatter = &logrus.JSONFormatter{}
		Log.Formatter = &logrus.TextFormatter{}
	}
}

// Info function
func Info(message string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Info(message)
}

// Error function
func Error(message string, err error, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}
	message = fmt.Sprintf("%s - ERROR - %v", message, err)
	Log.WithFields(parseFields(tags...)).Error(message)
}

// Debug function
func Debug(message string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}
	Log.WithFields(parseFields(tags...)).Debug(message)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		elements := strings.Split(tag, ":")
		result[strings.TrimSpace(elements[0])] = strings.TrimSpace(elements[1])
	}
	return result
}
