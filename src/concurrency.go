package main

/*** Goroutine奉行通过通信来共享内存，而不是共享内存来通信 ****/
import (
	"fmt"
)

func main() {
	// 创建一个channel
	c := make(chan bool)
	fmt.Println("888")
	go func() {
		fmt.Println("GO GO GO!!!")
		c <- true
		// close(c)
	}()
	// <-c
	for v := range c {
		fmt.Println("*", v)
	}

	// time.Sleep(5 * time.Second)
	// fmt.Println("hahah")
}
