package logger_test

import (
	"logger"
	"testing"
)

func TestInit(t *testing.T) {
	config := &logger.LoggerConfig{LogFile: "sample.log", LogLevel: logger.INFO}
	if !logger.Init(config) {
		t.Fail()
	}
}
