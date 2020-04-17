package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
)

func main() {

	// 将返回值作为打印参数
	fmt.Println(resolveTime(1000))
	// 只获取消息和分钟
	_, hour, minute := resolveTime(18000)
	//fmt.Println(hour, minute)
	// 只获取天
	day, _, _ := resolveTime(90000)
	fmt.Println(day)
	//panic("Go语言将秒转换为具体的时间")

	//$ go run -gcflags "-m -l" main/main.go
	var a1 int
	void()
	fmt.Println(a1, dummy(0))
	fmt.Println(dummy2())
	//panic("stop")

	zhizhen()

	CreateImage()

	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
	c := math.Ceil(123.4567890)
	fmt.Printf("%f", c)
	//i, j:=1,"abc"
	//fmt.Printf("i:%d, j:%s",i,j)
	// Engin
	router := gin.New()
	// 指定地址和端口号
	router.Run("127.0.0.1:15001")
}

const (
	// 定义每分钟的秒数
	SecondsPerMinute = 60
	// 定义每小时的秒数
	SecondsPerHour = SecondsPerMinute * 60
	// 定义每天的秒数
	SecondsPerDay = SecondsPerHour * 24
)

// 将传入的“秒”解析为3种时间单位
func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}

func zhizhen() {
	str := new(string)
	*str = "创建指针的另一种方法——new() 函数"
	fmt.Println(*str)
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}

// 交换函数
func swap(a, b *int) {
	// 取a指针的值, 赋给临时变量t
	t := *a
	// 取b指针的值, 赋给a指针指向的变量
	*a = *b
	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func CreateImage() {
	// 图片大小
	const size = 300
	// 根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	// 遍历每个像素
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			// 填充为白色
			pic.SetGray(x, y, color.Gray{255})
		}
	}
	// 从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		// 让sin的值的范围在0~2Pi之间
		s := float64(x) * 2 * math.Pi / size
		// sin的幅度为一半的像素。向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		// 用黑色绘制sin轨迹
		pic.SetGray(x, int(y), color.Gray{0})
	}
	// 创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	// 使用png格式将数据写入文件
	png.Encode(file, pic) //将image信息写入文件中
	// 关闭文件
	file.Close()
}

func GetData() (int, int) {
	return 100, 200
}

func QueryParam(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	//name := c.Request.URL.Query().Get("name")

	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"user": name,
	})
}

/*func GetUser(c *gin.Context) {
	id := c.Param("id")
	user := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"id":id,
		"user":user,
	})
}*/

func loginEndpoint(c *gin.Context) {
	//println(">>>> hello gin start <<<<")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"user": name,
	})
	//c.String(http.StatusOK,"user:%s, login gin start", name)
}

func submitEndpoint(c *gin.Context) {

}

func readEndpoint(c *gin.Context) {

}

func hello(c *gin.Context) {
	println(">>>> hello gin start <<<<")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"success": true,
	})
}
