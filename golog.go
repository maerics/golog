package golog

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const TimestampFormat = "2006-01-02T15:04:05.000Z"

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

var logdebugf = (func() func(string, ...any) {
	if debug {
		return func(message string, args ...any) {
			log.Printf("DEBUG: "+message, args...)
		}
	}
	return func(string, ...any) {}
})()

func init() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

func Debug(message string) {
	logdebug(message)
}

func Debugf(message string, args ...any) {
	logdebugf(message, args...)
}

func Print(message string) {
	log.Print(message)
}

func Printf(message string, args ...any) {
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

func Errorf(message string, args ...any) {
	log.Printf("ERROR: "+message, args...)
}

func Fatal(message string) {
	log.Fatal("FATAL: " + message)
}

func Fatalf(message string, args ...any) {
	log.Fatalf("FATAL: "+message, args...)
}

func Must(err error) {
	if err != nil {
		Fatal(err.Error())
	}
}

func Must1[T any](t T, err error) T {
	Must(err)
	return t
}

func Must2[T any, U any](t T, u U, err error) (T, U) {
	Must(err)
	return t, u
}

func Must3[T any, U any, V any](t T, u U, v V, err error) (T, U, V) {
	Must(err)
	return t, u, v
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
