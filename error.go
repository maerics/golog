package golog

import "log"

func Err(err error) {
	message := "<nil>"
	if err != nil {
		message = err.Error()
	}
	log.Printf("ERROR: " + message)
}

func Error(message string) {
	log.Print("ERROR: " + message)
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

func Fatale(err error) {
	log.Fatal("FATAL: " + err.Error())
}
