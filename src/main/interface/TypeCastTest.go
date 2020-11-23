package main

import "fmt"

/*
【类型转换】
类型推断可将接口变量还原为原始类型，或用来判断是否实现了某个更具体的接口类型。
*/

type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

func main() {
	var d data = 15
	var x interface{} = d

	if n, ok := x.(fmt.Stringer); ok { // 转换为更具体的接口类型
		fmt.Println(n)
	}

	if d2, ok := x.(data); ok { // 转换回原始类型
		fmt.Println(d2)
	}

	e := x.(error) // 错误:main.data is not error
	fmt.Println(e)
}
