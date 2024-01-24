package logger

import "go.uber.org/zap"

var Log *zap.Logger

func inti() {
	var err error
	Log, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}
