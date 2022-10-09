package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

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

	/************ 格式化字符串 *****************/
	p := point{1, 2}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true)
	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 14)
	fmt.Printf("%c\n", 96)
	fmt.Printf("%x\n", 456)
	fmt.Printf("%X\n", 456)
	fmt.Printf("%f\n", 456.8889900)
	fmt.Printf("%e\n", 456.8889900)
	fmt.Printf("%E\n", 456.8889900)
	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")
	fmt.Printf("%x\n", "hey this")
	fmt.Printf("%p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 345, 13)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
