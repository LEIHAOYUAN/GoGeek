package main

import "fmt"

/*
【接口使用技巧】
1、使用init，让编译器检查，确保类型实现了指定接口
2、定义函数类型，让相同签名的函数自动实现某个接口
*/
type FuncString func() string

func (f FuncString) String() string {
	return f()
}

func main() {
	var t fmt.Stringer = FuncString(func() string { // 转换类型，使其实现Stringer接口
		return "hello,world!"
	})

	fmt.Println(t)
}
