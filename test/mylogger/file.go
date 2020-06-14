package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志相关代码

// FileLogger 文件日志结构体
type FileLogger struct {
	Level		LogLevel
	filePath	string		// 日志文件保存的路径
	fileName	string		// 日志文件保存的文件名
	fileObj		*os.File	// 结构体的指针 （os包的file结构体）
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

// 根据指定的日志文件路径和文件名打开日志文件
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
func (f *FileLogger) enable(LogLevel LogLevel) bool {
	return LogLevel >= f.Level
}

// 判断文件是否需要切割
//func (f *FileLogger) checkSize(file *os.File) bool {
func (f *FileLogger) checkSize(file *os.File) (*os.File, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed, err:%v\n", err)
		return nil, err
	}
	// 如果当前文件大小大于等于文件的最大值 就应该返回true
	//return fileInfo.Size() >= f.maxFileSize
	if fileInfo.Size() < f.maxFileSize {
		return  file,nil
	}
	//return time.Now().Format("05") == "00" 	// 按照时间切割文件，判断是否为每分钟的00秒
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")

	logName := path.Join(f.filePath, fileInfo.Name())		// 拿到当前日志的完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)	// 拼接一个日志文件备份的名字
	// 1、关闭当前的日志文件
	file.Close()
	// 2、备份一下 rename	xx.log -> xx.log.bak202005231006
	os.Rename(logName, newLogName)
	// 3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v", err)
		return nil, err
	}
	// 4、将打开的信念日志哎文件对象赋值给f.fileObject
	return fileObj, nil
}

// 切割文件
func (f *FileLogger) splitFile(file *os.File)(*os.File, error) {
	// 需要切割日志文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed2, err:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())				// 拿到当前日志的完整路径
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)	// 拼接一个日志文件备份的名字
	// 1、关闭当前的日志文件
	file.Close()
	// 2、备份一下 rename	xx.log -> xx.log.bak202005231006
	os.Rename(logName, newLogName)
	// 3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed, err:%v", err)
		return nil, err
	}
	// 4、将打开的信念日志哎文件对象赋值给f.fileObject
	return fileObj, nil
}

// 记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, a...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		newFile, err := f.checkSize(f.fileObj)
		if err != nil {
			return
		}
		f.fileObj = newFile		//没有生效 ？	所有的上级f都应该传的是指针，不能使用值接收者
		/*if f.checkSize(f.fileObj)  {
			newFile, err := f.splitFile(f.fileObj)		// 日志文件
			if err != nil {
				return
			}
			f.fileObj = newFile
		}*/
		_, err = fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		if err != nil {
			fmt.Println("err:", err)
		}
		/*if lv >= ERROR {
			if f.checkSize(f.errFileObj)  {
				errNewFile, err := f.splitFile(f.errFileObj)		// 日志文件
				if err != nil {
					return
				}
				f.errFileObj = errNewFile
			}
			// 如果要记录的日志大于等于ERROR级别，我还要在err日志文件中在记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), funcName, fileName, lineNo, msg)
		}*/
	}
}

func (f *FileLogger) Debug(msg string) {
		f.log(DEBUG, msg)
}

func (f *FileLogger) Trace(msg string) {
		f.log(TRACE, msg)
}

func (f *FileLogger) Info(msg string) {
		f.log(INFO, msg)
}

func (f *FileLogger) Warning(msg string) {
		f.log(WARNING, msg)
}

func (f *FileLogger) Error(format string, a...interface{}) {
	// a = slice
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a...interface{}) {
	f.log(FATAL, format, a...)
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}