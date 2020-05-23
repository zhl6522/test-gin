package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

// Logger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// Newlog 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

// 上线后的错误级别部分输出到日志
func (c ConsoleLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= c.Level
}

func (c ConsoleLogger) log(lv LogLevel, format string, a...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(msg string) {
	c.log(DEBUG, msg)
}

func (c ConsoleLogger) Trace(msg string) {
	c.log(TRACE, msg)
}

func (c ConsoleLogger) Info(msg string) {
	c.log(INFO, msg)
}

func (c ConsoleLogger) Warning(msg string) {
	c.log(WARNING, msg)
}

func (c ConsoleLogger) Error(format string, a...interface{}) {
	c.log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a...interface{}) {
	c.log(FATAL, format, a...)
}