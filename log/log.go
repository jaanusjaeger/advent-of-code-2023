package log

import (
	"fmt"
	"log"
)

var debug bool = true

func SetDebug(d bool) {
	debug = d
}

func Debugf(format string, args ...any) {
	if !debug {
		return
	}
	printf("DEBUG", format, args...)
}

func Infof(format string, args ...any) {
	printf("INFO", format, args...)
}

func Errorf(format string, args ...any) {
	printf("ERROR", format, args...)
}

func printf(kind, format string, args ...any) {
	log.Printf("%-7s - %s", fmt.Sprintf("[%s]", kind), fmt.Sprintf(format, args...))
}
