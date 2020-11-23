package main

/*
如要等待多个任务结束，推荐使用sync.WaitGroup。通过设定计数器，让每个goroutine在退出前递减，直至归零时解除阻塞。
尽管WaitGroup.Add实现了原子操作，但建议在goroutine外累加计数器，以免Add尚未执行，Wait已经退出。
*/
import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 累加计数

		go func(id int) {
			defer wg.Done() // 递减计数

			time.Sleep(time.Second)
			println("goroutine", id, "done.")
		}(i)
	}

	println("main...")
	wg.Wait() // 阻塞，直到计数归零
	println("main exit.")
}
