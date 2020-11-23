package main

import "sync"

/*
【匿名字段】
1、可以像访问匿名字段成员那样调用其方法，由编译器负责查找
2、方法也会有同名遮蔽问题。但利用这种特性，可实现类似覆盖（override）操作。
3、尽管能直接访问匿名字段的成员和方法，但它们依然不属于继承关系
*/
func main() {
	testAnoCall()
	coverTest()
}

type data struct {
	sync.Mutex
	buf [1024]byte
}

func testAnoCall() {
	d := data{}
	d.Lock() // 编译会处理为sync.(*Mutex).Lock() 调用
	defer d.Unlock()
}

type user struct{}

type manager struct {
	user
}

func (user) toString() string {
	return "user"
}

func (m manager) toString() string {
	return m.user.toString() + ";manager"
}

/*
方法覆盖效果调用
*/
func coverTest() {
	var m manager

	println(m.toString())
	println(m.user.toString())
}
