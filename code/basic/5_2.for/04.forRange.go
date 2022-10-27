package main

import "fmt"

func main() {
	// 遍历数组、切片 - 获得索引和元素
	exampleSlice()

	// 遍历字符串 - 获得字符
	exampleString()

	// 遍历map - 获得map的键和值
	exampleMap()

	// 遍历通道 - 接收通道数据
	exampleChannel()

}

// 遍历数组、切片 - 获得索引和元素
func exampleSlice() {
	// key代表切片的下标，value代表下标对应的值
	for key, value := range []int{11, 22, 33, 44, 55} {
		fmt.Printf("key:%d   value:%d\n", key, value)
	}
}

// 遍历字符串 - 获得字符
func exampleString() {
	var str = "hello您好"
	for key, value := range str {
		fmt.Printf("key:%d   value:%c - 0x%x\n", key, value, value)
	}
	// 代码中的value变量，实际类型是rune，实际上就是int32，以十六进制打印出来就是字符的编码
}

// 遍历map - 获得map的键和值
//
// 【！！注意！！】
// 对map遍历时，遍历输出的键值是无序的；
// 如果需要有序的键值对输出
// 需要对结果进行排序
func exampleMap() {
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	// key代表map的索引，value代表索引对应的值
	// 一般被称为map的键值对，因为它们总是一对一对的出现
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// 遍历通道 - 接收通道数据
// 遍历时，只输出一个值，即 管道内的类型对应的数据
func exampleChannel() {

	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
