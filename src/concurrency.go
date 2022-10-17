package main

// Goroutine是由GO语言的运行时调度完成，而线程是由操作系统调度完成

/*
 * 并发 concurrency：把任务在不同的时间点交给处理器进行处理。在同一时间点，任务并不会同时运行。
 * 并行 parallelism：把每一个任务分配给每一个处理器独立完成。在同一时间点，任务一定是同时运行。
 *
 * GO在GOMAXPROCS数量与任务数量相等时，可以做到并行执行，但一般情况下都是并发执行。
 *
 * 在任何时候，同时只能有一个goroutine访问通道进行发送和获取数据。goroutine间通过通道就可以通信
 *
 *
 */
/*** Goroutine奉行通过通信来共享内存，而不是共享内存来通信 ****/
import (
	"fmt"
	"runtime"
	"time"
)

// 程序启动时，runtime会默认为main()函数创建一个goroutine
func main() {
	// 创建一个空接口channel
	c := make(chan interface{}) // 格式：  通道实例 := make(chan数据类型)

	//	开启一个并发匿名函数
	go func() {
		fmt.Println("start goroutine...")
		c <- 0       // 将0放入通道中
		c <- "hello" //	将hello字符串放入通道中
		c <- false   // 将false放入通道中
		// close关闭Channel
		close(c)
		fmt.Println("exit goroutine...")
	}()

	fmt.Println("wait goroutine ...")

	<-c
	// cd := <-c
	// fmt.Print("接收到：", cd)
	for v := range c {
		fmt.Println("*", v)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("GO1.5版本开始，默认执行`runtime.GOMAXPROCS( runtime.NumCPU())`，最大效率利用CPU，GO程序调度器可高效将CPU资源分配给每一个任务。", runtime.NumCPU())

	// // 并发执行程序
	// go running()

	// // 接受命令后输出，不做任何事情
	// var input string
	// fmt.Scanln(&input)

	/* 所有goroutine在main()函数结束时会一同结束 */
}

func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}
