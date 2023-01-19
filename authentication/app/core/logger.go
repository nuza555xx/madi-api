package core

import (
	"errors"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func DebugLogger(m string) {
	color.Green("%s - [ Logger ] 🚀 %s\n", time.Now().Format(time.RFC822), m)
}

func ErrorLogger(m string) {
	pc, _, line, _ := runtime.Caller(1)
	color.Red("%s - [ Exception ] 💩 '%s' at %s line :: %d\n", time.Now().Format(time.RFC822), errors.New(m), runtime.FuncForPC(pc).Name(), line)
}
