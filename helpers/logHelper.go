package helpers

import (
	"fmt"
	"log"
)

const(
	_=iota
	Panic
	Error
	Message
	Fatal
)

// Logger to define severiety of an error and handle appropriately throughout the application.
func Logger(msgType int,msg error){
	switch msgType {
	case Message:
		fmt.Print(msg)
	case Panic:
		log.Panic(msg)
	case Error:
		log.Print(msg)
	case Fatal:
		log.Fatal(msg)
	}
}

// LogMessage for printing msgs on console.
// For debugging purpose.
func LogMessage(msg ...interface{}){
	fmt.Println(msg)
}