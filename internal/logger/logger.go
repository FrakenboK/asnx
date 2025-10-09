package logger

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

type Logger struct{}

func (l *Logger) Fail(msg string) {
	fmt.Println(failString, msg)
}

func (l *Logger) Info(msg string) {
	fmt.Println(infoString, msg)
}

func (l *Logger) Banner() {
	randomColor.Println(banner)
	fmt.Println(versionString)
}

func (l *Logger) Note(msg string) {
	totalWidth := 70
	msgWidth := len(msg)

	if msgWidth >= totalWidth-4 {
		fmt.Printf("%s  %s  %s\n", color.GreenString(">"), msg, color.GreenString("<"))
		return
	}

	availableWidth := totalWidth - msgWidth - 4
	leftWidth := availableWidth / 2
	rightWidth := availableWidth - leftWidth

	leftArrows := strings.Repeat("=", leftWidth) + ">"
	rightArrows := "<" + strings.Repeat("=", rightWidth)

	fmt.Printf("%s %s %s\n",
		color.GreenString(leftArrows),
		msg,
		color.GreenString(rightArrows))
}

func New() *Logger {
	return &Logger{}
}
