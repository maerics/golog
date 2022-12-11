package golog

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const TimestampFormat = "2006-01-02T15:04:05.999Z"

var debug = (func() bool {
	return strings.TrimSpace(os.Getenv("DEBUG")) != ""
})()

var logdebug = (func() func(string) {
	if debug {
		return func(message string) {
			log.Print("DEBUG: " + message)
		}
	}
	return func(string) {}
})()

var logdebugf = (func() func(string, ...interface{}) {
	if debug {
		return func(message string, args ...interface{}) {
			log.Printf("DEBUG: "+message, args...)
		}
	}
	return func(string, ...interface{}) {}
})()

func init() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

func Debug(message string) {
	logdebug(message)
}

func Debugf(message string, args ...interface{}) {
	logdebugf(message, args...)
}

func Print(message string) {
	log.Print(message)
}

func Printf(message string, args ...interface{}) {
	log.Printf(message, args...)
}

func Err(err error) {
	message := "<nil>"
	if err != nil {
		message = err.Error()
	}
	log.Printf("ERROR: " + message)
}

func Error(message string) {
	log.Printf("ERROR: " + message)
}

func Errorf(message string, args ...interface{}) {
	log.Printf("ERROR: "+message, args...)
}

func Fatal(message string) {
	log.Fatal("FATAL: " + message)
}

func Fatalf(message string, args ...interface{}) {
	log.Fatalf("FATAL: "+message, args...)
}

type logWriter struct{}

func (lw *logWriter) Write(bs []byte) (int, error) {
	return fmt.Fprintf(
		os.Stderr,
		"%v | %v",
		time.Now().UTC().Format(TimestampFormat),
		string(bs),
	)
}
