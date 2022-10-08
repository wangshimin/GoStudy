package main

import "fmt"

/*********** 常量初始化 & 枚举 **************/
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

/*
6: 0110
11: 1011
---------
&   0010  = 2
|   1111  = 15
^   1101  = 13
&^  0100  = 4
*/
func main() {
	// fmt.Println(6 &^ 11)
	// fmt.Println(B, KB, MB, GB, TB)

	/*********** 指针 **************/
	// a := 1
	// a++
	// var p *int = &a // p存的是a的地址
	// var d = a
	// fmt.Println(a, p, d, &a, &p, &d, *p)

	/********** 判断语句if *************/
	// if a, b := 2, 5; a > 1 {
	// 	fmt.Println(a, b)
	// }

	/************ 循环语句for ************/
	// a := 1
	// for i := 0; i < 3; i++ {
	// 	a++
	// 	fmt.Println(a)
	// }

	// for a <= 3 {
	// 	a++
	// 	fmt.Println(a)
	// }

	// for {
	// 	fmt.Println(a)
	// 	if a >= 3 {
	// 		break
	// 	}
	// 	a++
	// }

	/************ 选择语句switch ************/
	// b := 1
	// switch b {
	// case 0:
	// 	fmt.Println("b=0")
	// case 1:
	// 	fmt.Println("b=1")
	// }

	// b := 1
	// switch {
	// case b >= 0:
	// 	fmt.Println("b大于等于0")
	// 	fallthrough
	// case b >= 1:
	// 	fmt.Println("b大于1")
	// }

	// switch b := 1; {
	// case b >= 0:
	// 	fmt.Println("b大于等于0")
	// 	fallthrough
	// case b >= 1:
	// 	fmt.Println("b大于1")
	// }

	// LABLE1:
	// 	for {
	// 		for i := 1; i < 10; i++ {
	// 			fmt.Println(i)
	// 			if i > 3 {
	// 				break LABLE1
	// 			}
	// 		}
	// 	}

	// fmt.Println("this is a test")
	// for c := 0; c < 10; c++ {
	// 	fmt.Printf("第%v行\n", c)
	// 	switch {
	// 	case c == 3:
	// 		fallthrough
	// 	case c == 6:
	// 		fallthrough
	// 	case c == 9:
	// 		fmt.Printf("哈哈哈，我现在是 %v !\n", c)
	// 	}
	// }

	/******* 数组 *******/
	// a := [...]int{3: 18, 9: 6, 0: 5}
	// b := [6]int{3, 5: 8}
	// //	指向数组的指针
	// var ca *[len(a)]int = &a
	// fmt.Println(a, b, ca)

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
