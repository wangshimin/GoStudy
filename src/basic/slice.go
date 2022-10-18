package main

import (
	"fmt"
)

func main() {

	/********** slice ***********/
	// b := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// fmt.Println(b)
	// s1 := b[:6]
	// fmt.Println(s1)

	// c1 := make([]int, 3, 10)
	// fmt.Println(len(c1), cap(c1))

	// d := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'}
	// d1 := d[3:5]  //	de
	// d2 := d1[2:4] //	fg
	// fmt.Println(string(d1), string(d2), cap(d1), cap(d2))

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5] // cde

	fmt.Println(string(a), len(a), cap(a), "\n", string(sa), len(sa), cap(sa))

	fmt.Println(string(sa[3:5])) //  fg
	fmt.Printf(" %p  \n %p \n", a, sa)

	//	append
	va := append(sa, 'l', 'm', 'n', 'o', 'p', 'q', 'r')

	fmt.Println(string(va), len(va), cap(va))

	fmt.Printf(" %p \n", va)

	t1 := []int{44, 55, 66}
	t2 := append(t1, 77, 88)

	fmt.Printf("%v %p \n %v %p", t1, t1, t2, t2)
}
