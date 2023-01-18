package common

import (
	"errors"
	"net/http"
	"runtime"
	"time"

	"github.com/fatih/color"
)

func RequestLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		color.Cyan("%s - [ Router ] 🥹 %s %s %s %s\n", time.Now().Format(time.RFC822), r.Host, r.Method, r.URL, r.UserAgent())
		handler.ServeHTTP(w, r)
	})
}

func DebugLogger(m string) {
	color.Green("%s - [ Logger ] 🚀 %s\n", time.Now().Format(time.RFC822), m)
}

func ErrorLogger(m string) {
	pc, _, line, _ := runtime.Caller(1)
	color.Red("%s - [ Exception ] 💩 '%s' at %s line :: %d\n", time.Now().Format(time.RFC822), errors.New(m), runtime.FuncForPC(pc).Name(), line)
}
