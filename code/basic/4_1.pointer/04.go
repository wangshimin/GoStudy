package main

import "fmt"

func main() {
	var i int = 10
	fmt.Println("i的地址=", &i)

	var ptr *int = &i
	fmt.Printf("ptr=%v \n", ptr)
	fmt.Printf("ptr 的地址=%v \n", &ptr)
	fmt.Printf("ptr 指向的值=%v \n", *ptr)

	*ptr = 20
	fmt.Printf("i=%v \n", i)
}
