package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	CRITICAL
	FATAL
)

var (
	l        *log.Logger // PRIVATE LOGGER INSTANCE
	logLevel LogLevel    // PRIVATE LOGGING LEVEL
)

type LoggerConfig struct {
	LogFile  string
	LogLevel LogLevel
}

func Init(config *LoggerConfig) bool {
	fileHandle := os.Stderr
	fileHandle, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if nil != err {
		panic(err)
		return false
	}
	l = log.New(fileHandle, "", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.SetOutput(fileHandle)
	log.Println("Logger initialised successfully.")
	setLogLevel(config.LogLevel)
	return true
}

func llog(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}

func Debug(format string, v ...interface{}) {
	//l.Output(3, fmt.Sprintf(format, v...))
	if DEBUG >= logLevel {
		llog("[DEBUG] "+format, v...)
	}
}

func setLogLevel(level LogLevel) {
	logLevel = level
}

func Info(format string, v ...interface{}) {
	if INFO >= logLevel {
		llog("[INFO] "+format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	if WARN >= logLevel {
		llog("[WARN] "+format, v...)
	}
}

func Error(format string, v ...interface{}) {
	if ERROR >= logLevel {
		llog("[ERROR] "+format, v...)
	}
}

func Fatal(format string, v ...interface{}) {
	if FATAL >= logLevel {
		llog("[FATAL] "+format, v...)
	}
}

func Critical(format string, v ...interface{}) {
	if CRITICAL >= logLevel {
		llog("[CRITICAL] "+format, v...)
		os.Exit(1)
	}
}

func main() {
	logConfig := &LoggerConfig{}
	Init(logConfig)
	Debug("This is sample test")
}
