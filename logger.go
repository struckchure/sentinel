package sentinel

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Logger struct{}

func (l *Logger) format(tag, message string) string {
	return fmt.Sprintf("[%s] %s\n", tag, message)
}

func (l *Logger) Warn(message string, v ...any) {
	color.Yellow(l.format("WARN", message), v...)
}

func (l *Logger) Info(message string, v ...any) {
	color.Green(l.format("INFO", message), v...)
}

func (l *Logger) Error(message string, v ...any) {
	color.Red(l.format("ERROR", message), v...)
	os.Exit(1)
}

var logger = new(Logger)
