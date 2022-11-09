/*
 * @Descripttion : 切片 slice
 * @version      :
 * @Author       :
 * @Date         : 2022-10-12 11:58:27
 * @LastEditors  : Please set LastEditors
 * @LastEditTime : 2022-11-09 14:54:29
 */

package main

import (
	"fmt"
	"reflect"
)

/*
 * 切片是数组一个连续片段的引用，是引用类型。
 * 是一个 长度可变的数组。
 * 0 <= len(s)切片长度 <= cap(s)切片容量
 *
 * 多个切片如果表示同一个数组的片段，它们可以共享数据；因此一个切片和相关数组的其他切片是共享存储的
 * 相反，不同的数组总是代表不同的存储。数组实际上是切片的构建块。
 *
 * 切片声明格式：var identufier []type
 *
 * 未初始化前默认为nil，长度为0
 * var slice1 []type = arr1[start:end]
 * 这表示 slice1 是由数组 arr1 从 start 索引到 end-1 索引之间的元素构成的子集（切分数组，start:end 被称为切片表达式）。
 */

func main() {
	// example1()

	// example2()

	// exampleToFunc()

	// exampleMake()

	// compareNewMake()

	// example3()

	// example4()

	exampleCopy()

	// exampleAppend()

	// 重置切片，清空拥有的元素
	// a := []int{1, 2, 3}
	// fmt.Println(a[0:0])
}

func example1() {
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片
	var numListEmpty = []int{} // 本来会在“{}”中填充切片的初始化元素，这里没填充，所以切片是空的。但此时numListEmpty已被分配了内存，只是没有元素

	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty) // [] [] []

	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty)) // 0 0 0

	// 切片判定空的结果
	fmt.Println(strList == nil)      // true
	fmt.Println(numList == nil)      // true
	fmt.Println(numListEmpty == nil) // false。numListEmpty已被分配内存，但没有元素，因此为false

	b := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("初始化了一个数组b，值为", b)
	s1 := b[:6] // 由b从 0 索引到 5 索引之间的元素构成的子集
	fmt.Println("b[:6]初始化一个切片，s1=", s1)
	fmt.Printf("s2 = b[:]，等于完整的b数组,此时s2的值为%v \n", b[:])
	fmt.Println("还有一种表述方式是 &b，譬如 s3=&b，s3的值为", &b)

	if reflect.DeepEqual(b[2:], b[2:len(b)]) {
		fmt.Println("b[2:] == b[2:len(b)]，都包含了数组从第三个到最后的所有元素 \n")
	}

	if reflect.DeepEqual(b[:3], b[0:3]) {
		fmt.Println("b[:3] == b[0:3]，包含了从第一个到第三个元素(不包括第四个) \n")
	}

	fmt.Printf("去掉最后一个元素，可以这么写：b[:len(b)-1]，值为%v\n", b[:len(b)-1])

}

func example2() {
	var arr1 [6]int              // 声明一个长度为6的整型数组
	var slice1 []int = arr1[2:5] // 初始化一个由arr1索引2到索引4之间的切片

	// 数组赋值，[0,1,2,3,4,5]
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	// 遍历打印slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice[%d] = %d \n", i, slice1[i])
	}
	fmt.Printf("arr1数组长度：%d，slice1切片长度：%d，slice1切片容量：%d \n", len(arr1), len(slice1), cap(slice1))

	slice1 = slice1[0:4] // 索引0到3
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice[%d] = %d \n", i, slice1[i])
	}
	fmt.Printf("slice1切片长度：%d，slice1切片容量：%d \n", len(slice1), cap(slice1))
}

/** 计算数组元素和
 * @Descripttion:将切片传递给函数
 * @return {*}
 */
func exampleToFunc() {
	var sum = func(a []int) int {
		s := 0
		for i := 0; i < len(a); i++ {
			s += a[i]
		}
		return s
	}
	var arr = [5]int{0, 1, 2, 3, 4}
	fmt.Println(sum(arr[:]))
}

/**
 * make()函数创建切片，同时创建好相关数组
 *
 * make(数据类型,数据个数，容量)
 */
func exampleMake() {
	var slice1 []int = make([]int, 10, 12)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
		fmt.Printf("slice[%d] = %d \n", i, slice1[i])
	}
	fmt.Printf("slice长度为%d，容量为%d\n", len(slice1), cap(slice1))
}

//!! 总结 [ 如何理解 new、make、slice、map、channel 的关系 ] !!
// 1.slice、map 以及 channel 都是 golang 内建的一种引用类型，三者在内存中存在多个组成部分， 需要对内存组成部分初始化后才能使用，而 make 就是对三者进行初始化的一种操作方式
// 2.new 获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行相应的初始化操作， 所以 slice、map、channel 需要 make 进行初始化并获取对应的内存地址，而非 new 简单的获取内存地址

/**  new() 和 make() 的区别
 * @Descripttion:
 * @return {*}
 */
func compareNewMake() {
	// new(T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 *T 的内存地址：
	// 这种方法 返回一个指向类型为 T，值为 0 的地址的指针
	// 它适用于值类型如数组和结构体；它相当于 &T{}
	var p *[]int = new([]int) // *p == nil
	p1 := new([]int)
	fmt.Println(*p, len(*p), cap(*p), *p1, len(*p1), cap(*p1))

	// make(T) 返回一个类型为 T 的初始值
	// 它只适用于 3 种内建的引用类型：切片、map 和 channel
	p2 := make([]int, 3)
	fmt.Println(p2, len(p2), cap(p2))

	v := make([]int, 10, 50)
	//	分配一个有 50 个 int 值的数组，
	// 并且创建了一个长度为 10，容量为 50 的切片 v，该切片指向数组的前 10 个元素
	fmt.Println(v, len(v), cap(v))

	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:] // e,m
	s2[1] = 't'
	fmt.Println(string(s1)) // poet

}

func example3() {
	// c1 := make([]int, 3, 10)
	// fmt.Println(len(c1), cap(c1))

	// d := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n'}
	// d1 := d[3:5]  //	de
	// d2 := d1[2:4] //	fg
	// fmt.Println(string(d1), string(d2), cap(d1), cap(d2))

	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5] // cde

	fmt.Println(string(a), len(a), cap(a), "\n", string(sa), len(sa), cap(sa))

	fmt.Println(string(sa[3:5])) //  fg
	fmt.Printf(" %p  \n %p \n", a, sa)

	//	append
	va := append(sa, 'l', 'm', 'n', 'o', 'p', 'q', 'r')

	fmt.Println(string(va), len(va), cap(va))

	fmt.Printf(" %p \n", va)

	t1 := []int{44, 55, 66}
	t2 := append(t1, 77, 88)

	fmt.Printf("%v %p \n %v %p", t1, t1, t2, t2)

}

func example4() {
	// 传入参数为一个 float32 切片，返回该切片的所有数字和
	var sum = func(a []float32) (sum float32) {
		for _, item := range a {
			sum += item
		}
		return
	}

	var a = []float32{1.0, 2.0, 3.0, 4.0}
	fmt.Printf("切片的和是%.2f\n", sum(a))
}

/** copy()函数的返回值表示实际发生复制的元素个数
 * @Descripttion:
 * @return {*}
 */
func exampleCopy() {
	slFrom := []int{1, 2, 3, 9}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("copied %d elements \n", n) // 4

	// 例子2
	const elementCount = 1000            // 设置元素数量为1000
	srcData := make([]int, elementCount) // 预分配足够多的元素切片
	for i := 0; i < elementCount; i++ {  // 将切片赋值
		srcData[i] = i
	}
	refData := srcData                                 // 引用切片数据
	copyData := make([]int, elementCount)              // 预分配足够多的元素切片
	copy(copyData, srcData)                            // 将数据复制到新的切片空间中
	srcData[0] = 999                                   // 修改原始数据的第一个元素
	fmt.Println(refData[0])                            // 打印引用切片的第一个元素
	fmt.Println(copyData[0], copyData[elementCount-1]) // 打印复印切片的第一个和最后一个元素
	copy(copyData, srcData[4:6])                       // 复制原始数据从4到6（不包含）
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}

/** append()函数的常见操作
 * @Descripttion:
 * @return {*}
 */
func exampleAppend() {

	var numbers []int //	声明一个整型切片
	for i := 0; i < 10; i++ {
		numbers = append(numbers, 1)                                                       // 	循环向numbers切片添加10个数
		fmt.Printf("len: %d，cap: %d, pointer: %p \n", len(numbers), cap(numbers), numbers) // 打印长度、容量、指针地址
	}

	// 上述代码总结：append()可以为切片动态添加元素。每个切片会指向一片内存空间，这片内存空间能容纳一定数量的元素
	// 当空间不能容纳足够多的元素时，切片就会进行“扩容”。
	// 切片在扩容时，容量的扩展规律按容量的2倍数扩充。如：1，2，4，8，16...

	var car []string

	// 添加1个元素
	car = append(car, "OldDriver")
	// 添加多个元素
	car = append(car, "Ice", "Sniper", "Monk")
	// 	添加切片
	team := []string{"Pig", "Flyingcake", "Chicken"}
	car = append(car, team...) // 在team后面加上了“...”，表示将team整个添加到car的后面
	fmt.Println(car)

	a := []string{"aa", "bb", "cc", "dd", "ee", "ff"}
	b := []string{"gg", "hh"}

	// 将切片 b 的元素追加到切片 a 之后: append(a, b...)
	a = append(a, b...)
	fmt.Println(a) // [aa, bb, cc, dd, ee, ff, gg, hh]

	//复制切片 a 的元素到新的切片 c 上
	s1 := make([]string, len(a))
	copy(s1, a)
	fmt.Println(s1) // [aa, bb, cc, dd, ee, ff, gg, hh]

	// 删除位于索引 i 的元素：a = append(a[:i], a[i+1:]...)
	a = append(a[:2], a[3:]...) // 删除cc
	fmt.Println(a)              // [aa, bb, dd, ee, ff, gg, hh]

	// 切除切片 a 中从索引 i 至 j 位置的元素：a = append(a[:i], a[j+1:]...)
	a = append(a[:1], a[3:]...)
	fmt.Println(a) // [aa, ee, ff, gg, hh ]

	// 为切片 a 扩展 j 个元素长度：a = append(a, make([]T, j)...)
	a = append(a, make([]string, 2)...)
	fmt.Println(len(a)) // 7

	// 在索引 i 的位置插入元素 x：a = append(a[:i], append([]T{x}, a[i:]...)...)
	a = append(a[:2], append([]string{"gogogo", "hello"}, a[2:]...)...)
	fmt.Println(a) // [aa, ee, hello, world, ff, gg, hh  ]

	// 在索引 i 的位置插入长度为 j 的新切片：a = append(a[:i], append(make([]T, j), a[i:]...)...)
	a = append(a[:2], append(make([]string, 3), a[2:]...)...)
	fmt.Println(a)

	// 在索引 i 的位置插入切片 b 的所有元素：a = append(a[:i], append(b, a[i:]...)...)
	c := []string{"you", "are", "best"}
	a = append(a[:2], append(c, a[2:]...)...)
	fmt.Println(a)
}
