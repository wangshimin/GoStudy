/*
 * @Descripttion : 列表（list） - 可以快速增删的非连续空间容器
 * @version      :
 * @Author       :
 * @Date         : 2022-11-09 15:17:08
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-09 16:19:01
 */

package main

import (
	"container/list"
	"fmt"
)

/*
	在GO中，列表使用container/list包来实现，内部的实现原理是：双链表
	列表可以高效地进行任意位置的元素插入和删除操作

	初始化
	方式1    变量名 := list.New()
	方式2    var 变量名 list.List

	列表的插入函数的返回值回提供一个*list.Element结构

这个结构记录着列表元素的值及和其他节点之间的关系等信息。
*/
func main() {

	l := list.New() // 创建列表实例

	//	尾部添加
	l.PushBack("canon") // canon

	// 尾部添加后保存元素句柄
	element := l.PushBack("two") // canon two。将two字符串插入到列表的尾部，并将这个元素的内部结构保存到element变量中

	//	头部添加
	l.PushFront(67) //	67 canon two

	l.PushBack("first") // 67 canon two first

	// 在two之后添加high
	l.InsertAfter("high", element) // 67 canon two high first。使用element变量，在element的位置后面插入high字符串

	// 在two之前添加noon
	l.InsertBefore("noon", element) // 67 canon noon two high first

	/*
	 遍历双链表需要配合Front()函数获取头元素
	 遍历时只要元素不为空就可以继续进行
	 每一次遍历调用元素的Next
	*/
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
