package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func main() {
	/*************************** defer 析构函数 **********************/
	for i := 0; i < 3; i++ {
		c := i + 10 //	声明并赋值变量c
		e := i * 3
		f := i + 2

		fmt.Printf("i=%d，c=%d,e=%d\n", i, c, e)

		defer func(ar *int, ab int) {
			e = 8
			ab = ab * 3
			fmt.Printf("\n three \n")
			fmt.Printf("i=%d，i+%d=%d\n", i, *ar, i+*ar)

			d := 2 * *ar //	声明并赋值变量d
			ar = &d      //	指针ar 重新指向变量d的地址，以此达到重新赋值的效果
			fmt.Printf("匿名函数体内给c指针的值 乘以2，此时c=%d, e=%d , ab=%d\n", *ar, e, ab)
		}(&c, f)

		fmt.Printf("------------c的值为%d,e=%d,ab=%d \n", c, e, f)

	}
	defer fmt.Println("two")

	//	defer按调用顺序的相反顺序逐个执行
	defer func() {
		fmt.Println("one")
	}()

	// var fs = [4]func(){}

	// for i := 0; i < 4; i++ {
	// 	defer fmt.Println("defer i = ", i)
	// 	defer func() { fmt.Println("defer_closure i = ", i) }()
	// 	fs[i] = func() { fmt.Println("closure i = ", i) }
	// }

	// for _, f := range fs {
	// 	f()
	// }

}
