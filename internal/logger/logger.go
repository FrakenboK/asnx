package logger

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/FrakenboK/asnx/internal/ui"
)

type Logger struct{}

func (l *Logger) Fail(msg string) {
	fmt.Println(ui.FailPrefix, msg)
}

func (l *Logger) Info(msg string) {
	fmt.Println(ui.InfoPrefix, msg)
}

func (l *Logger) RawLog(msg string) {
	fmt.Println(msg)
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
