package main

import "fmt"

func main() {

	exampleBreak()

	exampleGoto()
}

// 传统方式：Break
func exampleBreak() {

	var breakAgain bool

	// 外循环
	for x := 0; x < 10; x++ {

		// 内循环
		for y := 0; y < 10; y++ {

			fmt.Printf("x:%d - y:%d \n", x, y)
			// 满足某个条件时，退出循环
			if y == 2 {

				// 设置退出标记
				breakAgain = true

				// 退出本次循环
				break
			}
		}

		// 根据标记，还需要退出一次循环
		if breakAgain {
			break
		}
	}

	fmt.Println("done")
}

// goto语句通过标签进行代码间的无条件跳转。
// goto语句可以在快速跳出循环、避免重复退出上有一定帮助

// 使用goto退出多层循环
func exampleGoto() {

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Printf("x:%d - y:%d \n", x, y)
			if y == 2 {
				// 跳转到标签
				goto breakHere

			}
		}
	}

	//	手动返回，避免执行进入标签
	return // 标签只能被goto使用，但不影响代码执行流程，此处如果不手动返回，在不满足条件时，也会执行第65行代码

breakHere:
	fmt.Println("done")
}
