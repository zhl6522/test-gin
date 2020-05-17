package main
/*
1、支持往不同地方输出日志
2、日志分级别
	1.Debug
	2.Trace
	3.Info
	4.Warning
	5.Error
	6.Fatal
3、日志支持开关控制，比如开发的时候什么级别都输出，上线只有Info级别以上才能输出
4、完整的日志记录要包含时间、行号、文件名、日志级别、日志信息
5、日志文件要切割
*/
import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	//log.SetOutput(os.Stdout)	// os.Stdout 往终端写
	log.SetOutput(fileObj)
	for true {
		log.Println("这是一条log记录")
		time.Sleep(3 * time.Second)
	}
}
