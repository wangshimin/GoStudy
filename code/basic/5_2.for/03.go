package main

import "fmt"

func main() {
	exampleFor()
	exampleGoto()
}

// 使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值。
func exampleFor() {
	for i := 1; i <= 15; i++ {
		fmt.Printf("第%d次 \n", i)
	}
}

// 使用 goto 语句重写循环，要求不能使用 for 关键字
func exampleGoto() {
	i := 1
START:
	fmt.Printf("第%d次 \n", i)
	i++

	if i <= 15 {
		goto START
	}
}
