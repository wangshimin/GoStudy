package main

import "fmt"

type (
	byte int8
	rune int32
)

var (
	// 使用常规方式
	aaa = "hello"
	// 使用并行方式以及类型推断
	sss, bbb = 1, 2
	// ccc :=3 // 不可以省略var
)

func main() {
	a, _, c, d := 1, 2, 3, 4
	var e float32 = 100.1

	fmt.Println(a, c, d, e) // 1 3 4 100.1
	f := int(e)
	fmt.Println(f)

	var g int = 66
	h := string(g) // 100
	fmt.Println(h) // B
}
