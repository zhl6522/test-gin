package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address		string	`ini:"address"`
	Port		int		`ini:"port"`
	Username	string	`ini:"username"`
	Password	string	`ini:"password"`
}

type RedisConfig struct {
	Host		string	`ini:"host"`
	Port		int		`ini:"port"`
	Password	string	`ini:"password"`
	Datebase	int		`ini:"database"`
}

type Config struct {
	MysqlConfig	`ini:"mysql"`
	RedisConfig	`ini:"redis"`
}

func loadIni(fileName string,data interface{}) (err error) {
	// 0、参数的校验
	// 0.1、传进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	//fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer.")		// 新创建一个错误
		return
			//fmt.Errorf("data should be a pointer.")		// 格式化输出之后返回一个error类型
	}
	// 0.2、传进来的data参数必须是结构体类型指针（因为配置文件中各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer.")		// 新创建一个错误
		return
	}
	// 1、读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//string(b)		// 将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\r\n")
	//fmt.Printf("%#v\n", lineSlice)
	// 2、一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		// 去掉字符串首尾的空格
		line = strings.TrimSpace(line)
		// 如果是空行直接跳过
		if len(line) == 0 {
			continue
		}
		// 2.1、如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2、如果是[开头的就表示是节（section）
		if strings.HasPrefix(line, "[") {
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 去掉这一行首尾的[]后，把中间内容的首尾空格去掉
			sectionName := strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 2.3、如果不是[开头就是=分割的键值对
			// 2.3.1、以等号分割这一行，等号左边是key，右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") || strings.HasSuffix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			// 2.3.2、根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)		// 拿到嵌套结构体的值信息
			sType := sValue.Type()		// 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var filleName string
			// 2.3.3、遍历结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)		// tag信息是存储在类型信息中的
				// 2.3.4、如果key等于tag，给这个字段赋值
				if field.Tag.Get("ini") == key {
					// 找到了对应的字段
					fileName = field.Name
					break
				}
			}
		}
	}
	return


	/*v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {

	}*/
	// 反射中使用 Elem()方法获取指针对应的值
	/*if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}*/
}

func main() {
	var cfg Config
	//var x = new(int)
	//err := loadIni("./config.ini", x)
	err := loadIni("./config.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(cfg)
}
