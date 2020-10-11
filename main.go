package main

import "fmt"

type Message string

func InitMessage() Message {
	return Message(fmt.Sprintf("hello!"))
}

type Logger struct {
	Message Message
}

func InitLogger(m Message) Logger {
	return Logger{Message: m}
}

func main() {
	/*m := InitMessage()
	l := InitLogger(m)*/
	l := InitLoggerViaWire()
	fmt.Println(l)
}
