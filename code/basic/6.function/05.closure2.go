/*
 * @Descripttion : 闭包的应用
 * @version      :
 * @Author       :
 * @Date         : 2022-10-27 16:41:05
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-02 16:19:16
 */
package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

/**
 * 使用闭包调试
 * @Descripttion: 定义一个闭包函数，用来打印函数执行的位置
 * @return {*}
 */
var where func() = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("当前文件%s，第%d行", file, line)
}

func main() {
	// 应用闭包：将函数作为返回值
	example1()
	where()
	// 应用闭包：工厂函数
	example2()
}

/**
 * @Descripttion:  应用闭包：将函数作为返回值
 * @return {*}
 */
func example1() {
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives:%v \n", p2(3))
	TwoAddr := Adder(2)
	fmt.Printf("The result is: %v \n", TwoAddr(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}

/*
*

  - @Descripttion:  动态返回追加后缀

  - @param {*} string

  - @return {*}

    一个返回值为另一个函数的函数可以被称之为工厂函数，这在需要创建一系列相似的函数的时候非常有用：书写一个工厂函数而不是针对每种情况都书写一个函数。
*/
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func example2() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	where()
	fmt.Println(addBmp("file"))  // file.bmp
	fmt.Println(addJpeg("file")) // file.jpeg
}
