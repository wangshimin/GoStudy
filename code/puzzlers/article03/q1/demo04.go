/* 功能实现：怎样把命令源码文件中的代码拆分到其他库源码文件？
 *
 * 回答：填入代码包声明语句package main。（见demo04_lib.go）
 * 切记：在同一个目录下的源码文件都需要被声明为属于同一个代码包
 */

/*
 * 【 代 码 包 声 明 的 基 本 规 则 】
 *
 * 1.同目录下的源码文件的代码包声明语句要一致。也就是说，它们要同属于一个代码包。
 * 2.源码文件声明的代码包的名称可以与其所在的目录的名称不同。在针对代码包进行构建时，生成的结果文件的主名称与其父目录的名称一致。
 */
package main

import (
	"flag"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	hello(name)
}

// 运行方式1
// $ go run demo04.go demo04_lib.go

// 运行方式2：先构建当前的代码包再运行
// $ cd puzzlers/article03/q1
// $ go build
// $ ./q1 -name=wsm
