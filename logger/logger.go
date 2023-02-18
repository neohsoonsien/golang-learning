package logger

import (
	"go.uber.org/zap"
	"fmt"
)

func Logger() {

	// initialize zap
	logger := zap.NewExample()
	defer logger.Sync()

	PrintLog(logger)
}

func PrintLog(log *zap.Logger) {

	// initialise logger
	logger := log.Sugar()

	fmt.Println("Logger Package")
	logger.Infof("This is the Zap Suger log")
}