1. **c#、Lua、Python语言都支持`coroutine`特性。`coroutine`和`goroutine`都可以将函数或语句在独立的环境中运行，它们的区别是什么？**
- goroutine可能发生并行执行；但coroutine是中顺序执行。
>goroutine可能发生在多线程环境下，goroutine无法控制自己获取高优先度支持；coroutine始终发生在单线程，coroutine程序需要主动交出控制权，宿主才能获得控制权并将控制权交给其他coroutine
- goroutine间使用channel通信；coroutine使用yield和resume操作。
   
2.**并行、 并发概念**
- 并发是指同一时刻只有一个线程被执行，在同一时间段内有多个线程被执行。线程间是互相抢占资源的。
- 并行是指同一时间点多个线程同时执行。线程间不互相抢占资源。

3.**你知道设置GOPAH有什么意义吗？**
你可以把GOPATH简单理解成GO语言的工作目录，它的值是一个目录的路径，也可以是多个目录路径。每个目录代表GO语言的一个工作区（workspace）。
我们需要利用这些工作区，去放置GO语言的源码文件（source file），以及安装（install）后的归档文件（archive file，也就是以“.a”为扩展名的文件）和可执行文件（excuable file）

4.**命令源码文件的用途是什么？怎样编写它？**
命令源码文件是程序的运行入口，是每个可独立运行的程序必须拥有的。我们可以通过构建或安装，生成其与其对应的可执行文件，后者一般会与该命令源码文件的直接父目录同名。
如果一个源码文件声明属于main包，并且包含一个无参数声明且无结果声明的main函数，那么它就是命令源码文件。

5. Go 语言的类型推断可以带来哪些好处？
Go 语言的类型推断可以明显提升程序的灵活性，使得代码重构变得更加容易，同时又不会给代码的维护带来额外负担（实际上，它恰恰可以避免散弹式的代码修改），更不会损失程序的运行效率。

6. 




- 一个代码包中可以包含任意个以.go 为扩展名的源码文件，这些源码文件都需要被声明属于同一个代码包。
- 某个工作区的src子目录下的源码文件在安装后一般会放置到当前工作区的pkg子目录下对应的目录中，或者是被直接放置到该工作去的bin子目录中。

