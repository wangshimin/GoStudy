package main

import (
	"fmt"
)

func main() {
	/***  例子：无缓冲通道 **/
	// ch := make(chan int)
	// go func() {
	// 	for i := 3; i >= 0; i-- {
	// 		ch <- i
	// 		time.Sleep(time.Second)
	// 	}
	// }()

	// for data := range ch {
	// 	fmt.Println(data)
	// 	if data == 0 {
	// 		break
	// 	}
	// }

	/**** 例子：并发打印 ****/
	// 创建一个channel
	// c := make(chan int)
	// // 并发执行printer，传入channel
	// go printer(c)

	// for i := 1; i <= 10; i++ {
	// 	// 将数据通过channel投送给printer
	// 	c <- i
	// }
	// //	通知并发的printer结束循环（没数据啦！）
	// c <- 0
	// //	等待printer结束（搞定喊我！）
	// <-c

	// fmt.Println("退出咯！")

	// 创建一个3个元素缓冲大小的整型通道
	ch := make(chan int, 3)
	// 查看当前通道的大小
	fmt.Println(len(ch))

	// 发送3个整型元素到通道
	ch <- 10
	ch <- 20
	ch <- 30

	// 查看当前通道的大小
	fmt.Println(len(ch))

}

func printer(c chan int) {
	// 开始无限循环等待数据
	for {
		// 从channel中获取一个数据
		data := <-c

		// 将0视为数据结束
		if data == 0 {
			break
		}
		// 打印数据
		fmt.Println(data)
	}

	// 通知main已经结束循环（我搞定了！）
	c <- 0
}
