// +build wireinject

package main

import "github.com/google/wire"

func InitLoggerViaWire() Logger {
	wire.Build(InitMessage, InitLogger)
	return Logger{}
}