package golog

import (
	"fmt"
	"log"
	"os"
	"time"
)

const RFC3339MilliZulu = "2006-01-02T15:04:05.000Z"
const TimestampFormat = RFC3339MilliZulu

func init() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

func Print(message string) {
	log.Print(message)
}

func Printf(message string, args ...any) {
	log.Printf(message, args...)
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
