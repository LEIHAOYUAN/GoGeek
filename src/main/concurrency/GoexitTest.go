package main

import "runtime"

/*
【Goexit】
1、Goexit立即终止当前任务，运行时确保所有已注册延迟调用被执行。该函数不会影响其他并发任务，不会引发panic，自然也就无法捕获
2、无论身处哪一层，Goexit都能立即终止整个调用堆栈，这与return仅退出当前函数不同
标准库函数os.Exit可终止进程，但不会执行延迟调用
*/
func main() {
	exit := make(chan struct{})
	go func() {
		defer close(exit)  // 执行
		defer println("a") // 执行
		func() {
			defer func() {
				println("b", recover() == nil) // 执行，recover返回nil
			}()

			func() { // 在多层调用中执行Goexit
				println("c")
				runtime.Goexit()   // 立即终止整个调用堆栈
				println("c done.") // 不会执行
			}()

			println("b done.") // 不会执行
		}()
		println("a done.") // 不会执行
	}()

	<-exit
	println("main exit.")
}
