package logger

import (
	"fmt"
)

type Logger struct{}

func (l *Logger) Fail(msg string) {
	fmt.Println(failString, msg)
}

func (l *Logger) Info(msg string) {
	fmt.Println(infoString, msg)
}

func (l *Logger) Version() {
	fmt.Println(versionString)
}

func New() *Logger {
	return &Logger{}
}
