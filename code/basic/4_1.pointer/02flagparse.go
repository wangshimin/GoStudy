// 示例：使用指针变量获取命令行的输入信息
package main

import (
	"flag"
	"fmt"
)

// 定义命令行参数(参数名称、默认值、说明)，以指针类型返回
var mode = flag.String("mode", "", "process mode")

func main() {
	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode)
}

// 命令后运行
// go run 02flagparse.go --mode=fast
