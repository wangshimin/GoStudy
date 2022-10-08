package main

import (
	"fmt"
)

func main() {
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

	/********** slice ***********/
	b := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(b)
	s1 := b[:6]
	fmt.Println(s1)

	c1 := make([]int, 3, 10)
	fmt.Println(len(c1), cap(c1))

	d := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'}
	d1 := d[3:5]  //	de
	d2 := d1[2:4] //	fg
	fmt.Println(string(d1), string(d2), cap(d1), cap(d2))

}
