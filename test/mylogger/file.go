package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

type FileLogger struct {
	Level		LogLevel
	filePath	string		// 日志文件保存的路径
	fileName	string		// 日志文件保存的文件名
	fileObj		*os.File
	errFileObj	*os.File
	maxFileSize	int64
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile()	// 按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile()(error) {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed, err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName + ".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed, err:%v\n", err)
		return err
	}
	// 日志文件都已经打开
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

// 上线后的错误级别部分输出到日志
func (f FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

func (f FileLogger) log(lv LogLevel, format string, a...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			// 如果要记录的日志大于等于ERROR级别，我还要在err日志文件中在记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		}
	}
}

func (f FileLogger) Debug(msg string) {
		f.log(DEBUG, msg)
}

func (f FileLogger) Trace(msg string) {
		f.log(TRACE, msg)
}

func (f FileLogger) Info(msg string) {
		f.log(INFO, msg)
}

func (f FileLogger) Warning(msg string) {
		f.log(WARNING, msg)
}

func (f FileLogger) Error(format string, a...interface{}) {
		f.log(ERROR, format, a...)
}

func (f FileLogger) Fatal(format string, a...interface{}) {
		f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}