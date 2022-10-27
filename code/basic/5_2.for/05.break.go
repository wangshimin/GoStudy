// 跳出指定循环（break） -- 可以跳出多层循环
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	example1()
	example2()
	// exaple()
}

// break 默认跳出最近的for循环
func example1() {
	var i int = 5
	for {
		i = i - 1
		fmt.Printf("The variable i is now : %d \n", i)
		if i < 0 {
			break
		}

	}
}

// break 后面可以指定标签，跳出标签对应的for循环
func example2() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("**")
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 1:
				fmt.Println("哈哈哈")
			case 4:
				fmt.Println("纳尼")
			}
		}
	}
}

func exaple() {
	// time.Now().Unix():返回一个从1970:01:01 的0时0分0秒到现在的秒数
	rand.Seed(time.Now().Unix())
	// 如何随机的生成1-100整数
	n := rand.Intn(100) + 1 // (0 -100)
	fmt.Println(n)

	// 随机生成1-100的一个数，直到生成了99这个数，看看你一共用了几次
	// 分析思路：
	// 编写一个无限循环的控制，然后不停的随机生成数，当生成了99时，就退出这个无限循环
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		c := rand.Intn(100) + 1
		fmt.Println("c= ", c)
		count++
		if n == 81 {
			break //	表示跳出for循环
		}
	}
}
