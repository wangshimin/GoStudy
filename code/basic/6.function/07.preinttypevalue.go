// 示例：遍历可变参数列表 - 获取每一个参数的值
package main

import (
	"bytes"
	"fmt"
)

// 定义一个函数，参数数量为0～n，类型约束为字符串
func joinStrings(slist ...string) string {
	// 定义一个字节缓冲，快速地连接字符串
	var b bytes.Buffer
	// 遍历可变参数列表slist，类型为[]string
	for _, s := range slist {
		// 将遍历出的字符串连续写入字节数组
		b.WriteString(s)
	}
	// 将连接好的字节数组转换为字符串并输出
	return b.String()
}

// 获得每一个参数的类型和值
func printTypeValue(slist ...interface{}) string {
	// 字节缓冲作为快速字符串连接
	var b bytes.Buffer
	// 遍历参数
	for _, s := range slist {
		//	将interface{}类型格式化为字符串
		str := fmt.Sprintf("%v", s)

		// 类型的字符串描述
		var typeString string

		// 对s进行类型断言
		switch s.(type) {
		case bool:
			// 当s为布尔类型时
			typeString = "bool"
		case string:
			// 当s为字符串类型时
			typeString = "string"
		case int:
			// 当s为整型类型时
			typeString = "int"
		}
		// 写值字符串前缀
		b.WriteString("value：")

		// 写入值
		b.WriteString(str)

		// 写类型前缀
		b.WriteString(" type: ")

		// 写类型字符串
		b.WriteString(typeString)

		// 写入换行符
		b.WriteString("\n")
	}
	return b.String()
}

func main() {
	// 输入3个字符串，将它们连成一个字符串
	fmt.Println(joinStrings("pig ", "and", " rat"))
	fmt.Println(joinStrings("hammer", " mom", " and ", "hawk"))

	// 将不同类型的变量通过printTypeValue()打印出来
	fmt.Println(printTypeValue(100, "str", true))
}
