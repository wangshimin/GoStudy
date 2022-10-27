// continue：结束当前循环，继续执行下一次循环
//
// 在contine语句后添加标签时，表示开始标签对应的循环
package main

import "fmt"

func main() {
	example1()
	example2()
}

// continue实现，打印1-100之内的奇数
func example1() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("奇数是：", i)
	}
}

func example2() {
OuterLoop:
	for i := 0; i < 2; i++ {
		fmt.Printf("【 外：%d 】\n", i)
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println("发现了，下一个！")
				continue OuterLoop
			}
			fmt.Printf("内：%d \n", j)
		}
	}
}
