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
