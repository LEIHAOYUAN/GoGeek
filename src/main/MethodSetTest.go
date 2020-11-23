package main

import (
	"fmt"
	"reflect"
)

/**
【方法集】
类型有一个与之相关的方法集（method set），这决定了它是否实现某个接口
类型T方法集包含所有receiver T方法。
类型*T方法集包含所有receiver T+*T方法。
匿名嵌入S，T方法集包含所有receiver S方法。
匿名嵌入*S，T方法集包含所有receiver S+*S方法。
匿名嵌入S或*S，*T方法集包含所有receiver S+*S方法
*/
func main() {
	var t T
	methodSet(t) // 显示T方法集
	println("----------")
	methodSet(&t) // 显示 *T方法集
}

type S struct{}

type T struct {
	S
}

/*
利用反射来验证
*/
func methodSet(a interface{}) { // 显示方法集里所有方法名字
	t := reflect.TypeOf(a)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}
