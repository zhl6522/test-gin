package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关内容

type LogLevel uint16

const (
	// 定义日志级别
	Debug LogLevel = iota
	Trace
	Info
	Warning
	Error
	Fatal
)

// Logger 日志结构体
type Logger struct {
	Level LogLevel
}

// Newlog 构造函数
func Newlog(levelStr string) Logger {
	return Logger{}
}

func (l Logger) Debug(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Debug] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Trace(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Trace] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Info(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Info] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Warning(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Warning] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Error(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Error] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}

func (l Logger) Fatal(msg string) {
	now := time.Now()
	fmt.Printf("[%s] [Fatal] %s\n", now.Format("2006-01-02 15:04:05"), msg)
}