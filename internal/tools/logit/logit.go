package logit

import (
	"fmt"
	"github.com/TwiN/go-color"
	"log"
	"os"
	"runtime"
)

func appendFilePath() string {
	_, file, no, ok := runtime.Caller(3)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s:%d -", file, no)
}

func general(prefix string, c string, v ...any) {
	if os.Getenv("COLOR_DEBUG") != "" {
		coloredPrefix := color.InBold(color.Ize(c, prefix))
		log.Println(append([]interface{}{coloredPrefix, appendFilePath()}, v...)...)
	} else {
		log.Println(append([]interface{}{prefix, appendFilePath()}, v...)...)
	}

}

func Info(v ...any) {
	general("[INFO]", color.Green, v...)
}

func Warn(v ...any) {
	general("[WARN]", color.Yellow, v...)
}

func Error(v ...any) {
	general("[ERROR]", color.Red, v...)
}

func Debug(v ...any) {
	general("[DEBUG]", color.Blue, v...)
}
