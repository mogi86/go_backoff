// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

// Injectors from wire.go:

func InitLoggerViaWire() Logger {
	message := InitMessage()
	logger := InitLogger(message)
	return logger
}
