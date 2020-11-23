package main

/*
1、在函数调用前，会为形参和返回值分配空间，并将实参拷贝到形参内存
2、一般来说指针类型的函数参数性能更好一些，但实际需要具体分析，被复制的指针会延长目标对象生命周期
还可能导致它被分配在堆上，那么其消耗就得加上堆内存分配和垃圾回收的成本
3、在栈上复制小对象只需要很少的指令即可完成，另外并发编程也提倡尽可能使用不可变对象（只读|复制），可以消除数据同步的麻烦
4、如果复制成本很高，或需要修改原对象的状态，自然使用指针更好
*/
func main() {

	//匿名函数，可以直接调用
	func(s string) {
		println(s)
	}("hello world")
}

// 定义函数类型
type TestFunc func(string, ...interface{}) (string, error)

func format(f TestFunc, s string, a ...interface{}) (string, error) {
	return f(s, a...)
}
