package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
【Goexit】
如果在main.main里调用Goexit，它会等待其他任务结束，然后让进程直接崩溃
*/
func main() {
	for i := 0; i < 2; i++ {
		go func(x int) {
			for n := 0; n < 2; n++ {
				fmt.Printf("%c: %d\n", 'a'+x, n)
				time.Sleep(time.Millisecond)
			}
		}(i)
	}

	runtime.Goexit() // 等待所有任务结束
	println("main exit.")
}
