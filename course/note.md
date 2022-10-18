1. c#、Lua、Python语言都支持`coroutine`特性。`coroutine`和`goroutine`都可以将函数或语句在独立的环境中运行，它们的区别是什么？
- goroutine可能发生并行执行；但coroutine是中顺序执行。
>goroutine可能发生在多线程环境下，goroutine无法控制自己获取高优先度支持；coroutine始终发生在单线程，coroutine程序需要主动交出控制权，宿主才能获得控制权并将控制权交给其他coroutine
- goroutine间使用channel通信；coroutine使用yield和resume操作。
   
2. 并行、 并发概念
- 并发是指同一时刻只有一个线程被执行，在同一时间段内有多个线程被执行。线程间是互相抢占资源的。
- 并行是指同一时间点多个线程同时执行。线程间不互相抢占资源。


