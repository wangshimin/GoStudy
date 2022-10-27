// 示例：使用指针修改值
package main

import (
	"fmt"
)

// 交换函数
func swap(a, b *int) {
	// 取a指针的值，赋给临时变量t
	t := *a

	// 取b指针的值，赋给a指针指向的变量
	*a = *b

	// 将a指针的值赋给b指针指向的变量
	*b = t
}

func main() {

	// a := 1
	// a++
	// var p *int = &a // p存的是a的地址
	// var d = a
	// fmt.Println(a, p, d, &a, &p, &d, *p)

	// 准备两个变量，赋值1和2
	x, y := 1, 2

	// 交换变量值
	swap(&x, &y)

	// 输出变量值
	fmt.Println(x, y)
}

//总结：“*”操作符的根本意义就是操作指针指向的变量。当操作右值时，就是取指向变量的值；当操作在左边时，就是将值设置给指向的变量。
