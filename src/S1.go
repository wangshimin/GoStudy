package main

import (
	"fmt"
)

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5] // cde

	fmt.Println(string(a), len(a), cap(a), "\n", string(sa), len(sa), cap(sa))

	fmt.Println(string(sa[3:5])) //  fg
	fmt.Printf(" %p  \n %p \n", a, sa)

	va := append(sa, 'l', 'm', 'n', 'o', 'p', 'q', 'r')

	fmt.Println(string(va), len(va), cap(va))

	fmt.Printf(" %p \n", va)

	t1 := []int{44, 55, 66}
	t2 := append(t1, 77, 88)

	fmt.Printf("%v %p \n %v %p", t1, t1, t2, t2)
}
