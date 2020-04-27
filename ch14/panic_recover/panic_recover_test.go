package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {	//不推荐使用：1、形成僵尸服务进程 2、错误信息被忽略
			//错误恢复
			fmt.Println("recovered panic ", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))	//从输出你可以看出defer处理了panic中抛出的error
}
