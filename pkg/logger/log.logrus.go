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

	// log.SetReportCaller(true)

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

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	log.Info(args...)
}

// InfoF logs a message with the INFO level.
func InfoF(message string, fields logrus.Fields) {
	log.WithFields(fields).Info(message)
}

// InfoFF logs a formatted message with the INFO level.
func InfoFF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Infof(format, args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// DebugF logs a message with the DEBUG level.
func DebugF(message string, fields logrus.Fields) {
	log.WithFields(fields).Debug(message)
}

// DebugFF logs a formatted message with the DEBUG level.
func DebugFF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Debugf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	log.Error(args...)
}

// ErrorF logs a message with the ERROR level.
func ErrorF(message string, fields logrus.Fields) {
	log.WithFields(fields).Error(message)
}

// ErrorFF logs a formatted message with the ERROR level.
func ErrorFF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// FatalF logs a message with the FATAL level and exits the application.
func FatalF(message string, fields logrus.Fields) {
	log.WithFields(fields).Fatal(message)
}

// FatalFF logs a formatted message with the FATAL level and exits the application.
func FatalFF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Fatalf(format, args...)
}

// PanicF logs a message with the PANIC level and then panics.
func PanicF(message string, fields logrus.Fields) {
	log.WithFields(fields).Panic(message)
}

// PanicFF logs a formatted message with the PANIC level and then panics.
func PanicFF(format string, fields logrus.Fields, args ...interface{}) {
	log.WithFields(fields).Panicf(format, args...)
}
