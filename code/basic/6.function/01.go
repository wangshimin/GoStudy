package main

import (
	"flag"
	"fmt"
)

func main() {
	// 函数变量 - 把函数作为值保存到变量中
	example1()

	// 匿名函数 - 没有函数名字的函数
	example2()

	// 匿名函数用作回调函数
	exampleCallback()

	// 使用匿名函数实现操作封装
	exampleFlag()
}

// 函数变量 - 把函数作为值保存到变量中
func example1() {
	var f func() // 将变量f声明为func()类型，此时f就被俗称为“回调函数”。此时f的值为nil
	fmt.Printf("类型：%T， 值：%v \n", f, f)
	f = fire
	f()
}

func fire() {
	fmt.Println("fire")
}

// 匿名函数 - 没有函数名字的函数
// 经常被用于实现回调函数、闭包等
func example2() {
	//	将匿名函数体保存到f()中
	f := func(data int) {
		fmt.Println("hello", data)
	}

	// 使用f()调用
	f(100)
}

// 匿名函数用作回调函数
func exampleCallback() {
	// 实现：对切片的遍历操作，遍历中访问每个元素的操作使用匿名函数来实现

	// 使用匿名函数打印切片内容
	visit([]int{1, 2, 3, 4}, func(v int) {
		fmt.Println(v)
	})
}

// 遍历切片的每个元素，通过给定函数进行元素访问
func visit(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func exampleFlag() {

	var skillParam = flag.String("skill", "", "skill to perform") // 定义命令行参数，从命令行输入skill可以将空格后的字符串传入skillparam指针变量
	flag.Parse()                                                  // 解析命令行参数，skillParam指针变量将指向命令行传入的值

	var skill = map[string]func(){ // 定义一个从字符串映射到func()的map，然后填充这个map
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok { // skillParam是一个*string类型的指针变量，使用*skillParam获取到值，并在map中查找对应字符串的函数
		f()
	} else {
		fmt.Println("skill not found")
	}
}

// 命令行运行
// go run 01.go -skill=fly
// go run 01.go -skill=run
