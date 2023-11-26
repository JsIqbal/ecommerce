package logger

import (
	"runtime"
	"strings"
)

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(3)
	funcName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(funcName, ".")

	return parts[len(parts)-1]
}
