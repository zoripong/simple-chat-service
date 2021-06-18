package public

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func GetWarningLogger() *Logger {
	return &Logger{log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime|log.Lshortfile)}
}

func GetInfoLogger() *Logger {
	return &Logger{log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)}
}
