/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-07 15:15:02
 */
package main

import "fmt"

// 概念：数组是一段固定长度的连续内存区域。

func main() {
	//	数组声明、赋值
	// example1()

	// 冒泡排序的实现
	example2()

}

func example1() {
	//	声明数组的格式：var 数组变量名 [数据大小]数据类型
	var a1 [4]string
	a1[0] = "hammer"
	a1[1] = "solider"
	a1[2] = "mum"
	for k, v := range a1 {
		fmt.Println(k, v) // k为索引，v为元素值
	}

	// 声明并赋值
	// var a2 = [3]string{"hammer", "soldier", "mum"}
	//	...表示让编译器确定数组大小
	// var a3 = [...]string{"macos", "window", "linux"}
	var a4 = [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Printf("%T", a4)

	// 数组声明后，默认值为零值
	var arr01 [3]float32 // 1.数值（整数系列、浮点数系列） => 0
	var arr02 [3]string  // 2. 字符串 => ""
	var arr03 [3]bool    // 3. 布尔类型=> false
	fmt.Printf("arr01=%v arr02=%v arr03=%v \n", arr01, arr02, arr03)

	a := [...]int{3: 18, 9: 6, 0: 5}
	b := [6]int{3, 5: 8}
	//	指向数组的指针
	var ca *[len(a)]int = &a
	fmt.Println(a, b, ca)

	x, y := 3, 7
	//	指针数组
	d := [...]*int{&x, &y}
	//	new返回数组的指针
	p := new([10]int)

	//	多维数组
	e := [2][3]int{
		{1: 1, 2: 6},
		{2: 2}}
	fmt.Println(d, &p, e)
}

/**
 * @Descripttion: 冒泡排序(倒序)
 * @return {*}
 */
func example2() {
	// 数组
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println(a)

	//	冒泡排序
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				tmp := a[i]
				a[i] = a[j]
				a[j] = tmp
			}

		}
	}
	fmt.Println(a)

}
