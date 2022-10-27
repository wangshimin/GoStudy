//在多个可变参数函数中传递参数

// 可变参数变量是一个包含所有参数的切片
// 如果要在多个可变参数中传递参数
// 可以在传递时在可变参数变量中默认添加“...”
// 将切片中的元素进行传递
// 而不是传递可变参数变量本身
package main

import "fmt"

// 实际打印的函数
func rawPrint(rawList ...interface{}) {
	// 遍历可变参数切片
	for _, a := range rawList {
		//	打印参数
		fmt.Println(a)
	}
}

// 打印函数封装
func print(slist ...interface{}) {
	// 将slist可变参数切片完整传递给下一个函数
	rawPrint(slist...)
}

func main() {
	print(1, 2, 3)
}
