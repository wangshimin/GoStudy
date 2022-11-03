/*
 * @Descripttion : 析构函数 defer
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-10-31 12:26:45
 */

package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	// example1()
	// example2()
	example3()
}

// 知识点：匿名函数、闭包
func example1() {
	for i := 0; i < 3; i++ {
		c := i + 10 //	声明并赋值变量c
		e := i * 3
		f := i + 2

		fmt.Printf("i=%d，c=%d,e=%d\n", i, c, e)

		defer func(ar *int, ab int) {
			e = 8
			ab = ab * 3
			fmt.Printf("\n three \n")
			fmt.Printf("i=%d，i+%d=%d\n", i, *ar, i+*ar)

			d := 2 * *ar //	声明并赋值变量d
			ar = &d      //	指针ar 重新指向变量d的地址，以此达到重新赋值的效果
			fmt.Printf("匿名函数体内给c指针的值 乘以2，此时c=%d, e=%d , ab=%d\n", *ar, e, ab)
		}(&c, f)

		fmt.Printf("------------c的值为%d,e=%d,ab=%d \n", c, e, f)

	}
	defer fmt.Println("two")

	//	defer按调用顺序的相反顺序逐个执行
	defer func() {
		fmt.Println("one")
	}()
}

func example2() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)                      // 赋值
		defer func() { fmt.Println("defer_closure i = ", i) }() //	闭包 4
		fs[i] = func() { fmt.Println("closure i = ", i) }       //	闭包 4
	}

	for _, f := range fs {
		f()
	}
}

func example3() {
	fmt.Println("defer begin")

	// 将defer放入延迟调用栈
	defer fmt.Println(1)

	defer fmt.Println(2)

	// 最后一个放入，位于栈顶，最先调用
	defer fmt.Println(3)

	fmt.Println("defer end")
}
