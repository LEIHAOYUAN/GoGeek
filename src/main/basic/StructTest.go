package main

import "fmt"

/*
1、结构体（struct）将多个不同类型命名字段（field）序列打包成一个复合类型
2、字段名必须唯一，可用“_”补位，支持使用自身指针类型成员
*/
func main() {
	n1 := node{
		id: 1,
	}

	n2 := node{
		id:   2,
		next: &n1,
	}

	fmt.Println(n1, n2)
}

type node struct {
	_    int
	id   int
	next *node
}
