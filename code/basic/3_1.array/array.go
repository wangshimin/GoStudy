package main

import "fmt"

func main() {

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

	//	数组
	// a := [...]int{5, 2, 6, 3, 9}
	// fmt.Println(a)

	// //	冒泡排序
	// num := len(a)
	// for i := 0; i < num; i++ {
	// 	for j := i + 1; j < num; j++ {
	// 		if a[i] < a[j] {
	// 			tmp := a[i]
	// 			a[i] = a[j]
	// 			a[j] = tmp
	// 		}

	// 	}
	// }
	// fmt.Println(a)

}
