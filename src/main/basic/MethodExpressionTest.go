package main

import "fmt"

/*
【表达式】
1、方法和函数一样，除直接调用外，还可赋值给变量，或作为参数传递。依照具体引用方式的不同，可分为expression和value两种状态。


*/
func main() {
	var n MN = 25
	fmt.Printf("main.n: %p, %d\n", &n, n)

	f1 := MN.test //func(n N)
	f1(n)

	f2 := (*MN).test //func(n*N)
	f2(&n)           // 按方法集中的签名传递正确类型的参数

	MN.test(n)
	(*MN).test(&n) // 注意: *N须使用括号，以免语法解析错误
}

type MN int

func (n MN) test() {
	fmt.Printf("test.n: %p, %d\n", &n, n)
}
