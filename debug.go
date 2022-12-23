package golog

import (
	"log"
	"os"
	"strings"
)

func Debug(message string) {
	logdebug(message)
}

func Debugf(message string, args ...any) {
	logdebugf(message, args...)
}

var logdebug = (func() func(string) {
	if debug {
		return func(message string) {
			log.Print("DEBUG: " + message)
		}
	}
	return func(string) {}
})()

var logdebugf = (func() func(string, ...any) {
	if debug {
		return func(message string, args ...any) {
			log.Printf("DEBUG: "+message, args...)
		}
	}
	return func(string, ...any) {}
})()

var debug = (func() bool {
	return strings.TrimSpace(os.Getenv("DEBUG")) != ""
})()
