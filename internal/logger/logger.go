package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// L is logger
var L *zap.SugaredLogger

func InitSugar() {
	l, _ := zap.NewDevelopment(zap.AddStacktrace(zap.ErrorLevel))
	L = l.Sugar()
}

func SyncForExit() {
	err := L.Sync()
	if err != nil {
		fmt.Println(err)
	}
}

// Init is not a popular mechanic but is justified for connecting the log to tests.
func init() {
	InitSugar()
}
