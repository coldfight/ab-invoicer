package logit

import (
	"fmt"
	"log"
	"runtime"
)

func appendFilePath() string {
	_, file, no, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d -", file, no)
}

func general(prefix string, v ...any) {
	log.Println(append([]interface{}{prefix, appendFilePath()}, v...)...)
}

func Info(v ...any) {
	general("[INFO]", v...)
}

func Warn(v ...any) {
	general("[WARN]", v...)
}

func Error(v ...any) {
	general("[ERROR]", v...)
}

func Debug(v ...any) {
	general("[DEBUG]", v...)
}
