package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// init initializes the logger with specific settings.
func init() {
	// create a new logger instance
	log = logrus.New()

	// set output to stdout
	log.SetOutput(os.Stdout)

	// set log formatter
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

// SetLevel set level to write log.
// Accepted value: panic, fatal, error, warn, warning, info, debug
func SetLevel(lvl string) error {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		return err
	}

	log.SetLevel(level)
	return nil
}

// Info logs a message with the INFO level.
func Info(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

// InfoF logs a formatted message with the INFO level.
func InfoF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Infof(format, args...)
}

// Debug logs a message with the DEBUG level.
func Debug(message string, fields logrus.Fields) {
	log.WithFields(fields).Debug(message)
}

// DebugF logs a formatted message with the DEBUG level.
func DebugF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Debugf(format, args...)
}

// Error logs a message with the ERROR level.
func Error(message string, fields logrus.Fields) {
	log.WithFields(fields).Error(message)
}

// ErrorF logs a formatted message with the ERROR level.
func ErrorF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Errorf(format, args...)
}

// Fatal logs a message with the FATAL level and exits the application.
func Fatal(message string, fields logrus.Fields) {
	log.WithFields(fields).Fatal(message)
}

// FatalF logs a formatted message with the FATAL level and exits the application.
func FatalF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Fatalf(format, args...)
}

// Panic logs a message with the PANIC level and then panics.
func Panic(message string, fields logrus.Fields) {
	log.WithFields(fields).Panic(message)
}

// PanicF logs a formatted message with the PANIC level and then panics.
func PanicF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Panicf(format, args...)
}
