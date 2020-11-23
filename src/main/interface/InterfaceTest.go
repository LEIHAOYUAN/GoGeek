package main

import "fmt"

/*
【接口】
1、接口代表一种调用契约，是多个方法声明的集合。
2、对于相同包，或者不会频繁变化的内部模块之间，并不需要抽象出接口来强行分离。接口最常见的使用场景，是对包外提供访问，或预留扩展空间。
3、Go接口实现机制很简洁，只要目标类型方法集内包含接口声明的全部方法，就被视为实现了该接口，无须做显示声明。当然，目标类型可实现多个接口。
4、可以像匿名字段那样，嵌入其他接口。目标类型方法集中必须拥有包含嵌入接口方法在内的全部方法才算实现了该接口
不能有字段。
不能定义自己的方法。
只能声明方法，不能实现。
可嵌入其他接口类型

*/

func main() {
	var o valueable
	o = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	var t f
	o = t
	showValue(o)
}

type valueable interface {
	getValue() float32
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

/**
定义结构体f，实现getValue方法
*/
type f float32

func (s f) getValue() float32 {
	return 0.1
}

func showValue(val valueable) {
	fmt.Printf("value of the asset is %f\n", val.getValue())
}
