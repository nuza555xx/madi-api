package core

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

func DebugLogger(m string) {
	log.WithFields(log.Fields{
		"time": time.Now().Format(time.RFC822),
		"type": "Logger",
	}).Debug(m)
}

func ErrorLogger(m string) {
	pc, file, line, _ := runtime.Caller(1)
	log.WithFields(log.Fields{
		"time":     time.Now().Format(time.RFC822),
		"type":     "Exception",
		"error":    errors.New(m),
		"location": fmt.Sprintf("%s:%d", file, line),
		"func":     runtime.FuncForPC(pc).Name(),
	}).Error(m)
}
