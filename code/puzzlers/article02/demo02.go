/* 实现功能：根据运行程序时给定的参数问候某人 */
package main

// GO语言标准库中的flag代码包：用于接收和解析命令参数
import (
	"flag"
	"fmt"
)

var name string

func init() {
	// 参数说明：参数地址、参数名、默认值、注释
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse() // 解析命令参数，并把它们的值赋给相应的变量
	fmt.Printf("Hello,%s!\n", name)
}

// 命令行传入参数
// $ go run demo02.go -name="tara"

// 查看参数使用说明
// $ go run demo02.go --help
