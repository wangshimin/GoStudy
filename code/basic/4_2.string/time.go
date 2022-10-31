/*
 * @Descripttion :
 * @version      :
 * @Author       :
 * @Date         : 2022-10-31 13:28:14
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-10-31 13:59:28
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	// basic()
	exampleSleep()
}

func basic() {
	// 1.获取当前时间
	now := time.Now()
	fmt.Printf("now:%v \ntype:%T", now, now)

	// 2.通过now可以获取到年月日、时分秒,通过Printf或SPrintf格式化
	fmt.Printf("现在是%v年%v月%v日%v时%v分%v秒 \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// 3.使用now.Format()格式化
	fmt.Printf(now.Format("示例2006年01/02 13:15:23"))

}

func exampleSleep() {
	// 每隔（1秒/0.1秒）打印一个数字，打印到100时就退出
	i := 0
	for {
		i++
		// 休眠
		// time.Sleep(time.Second)       // 1秒
		time.Sleep(time.Millisecond * 100) // 0.1秒
		fmt.Println(i)
		if i == 100 {
			break
		}
	}
}
