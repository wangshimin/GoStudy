// 当前程序的包名
package main

// 导入其他的包
// import "fmt"
import (
	"fmt"
	// "os"
	// "io"
	// "time"
	// "strings"
)

// 常量的定义
// const PI = 3.14
const (
	PI    = 3.14
	one   = "1"
	two   = 2
	three = 3
)

// 全局变量的声明与赋值
// var name = "gopher"
var (
	name = "gopher"
	role = "admin"
	age  = 18
	sex  = "woman"
)

// 一般类型声明
// type newType int
type (
	newType int
	type1   float32
	type2   string
	type3   byte
)

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由 main 函数作为程序入口点启动
func main() {
	fmt.Println(PI + 1 + two)
	fmt.Println("hello" + " la" + sex)

}
