package main

import "fmt"

type Message string

func InitMessage() Message {
	return Message("hello!")
}

type Logger struct {
	Message Message
}

func InitLogger(m Message) Logger {
	return Logger{Message: m}
}

func main() {
	backID := 1
	fmt.Println(backID)

	hashID := 2
	fmt.Println(hashID)

	/*m := InitMessage()
	l := InitLogger(m)*/
	l := InitLoggerViaWire()
	fmt.Println(l.Message)
}
