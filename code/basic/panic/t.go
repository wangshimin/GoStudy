/*
 * @Descripttion : panic 宕机，程序终止运行
 * @version      :
 * @Author       :
 * @Date         : 2022-11-02 16:29:36
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-02 16:33:41
 */
package main

// GO在宕机时，会将堆栈和goroutine信息输出到控制台
// 所以宕机也可以方便地知晓发生错误的位置

// 语法：func panic(v interface{})
// panic()的参数可以是任意类型
func main() {
	panic("cccc")
}
