package main

func main() {

	test()
	println("-------------------------")
	test01()

}

func test() {
	const (
		a = iota //0
		b        //1
		c = 100  //100
		d        //100（与上一行常量右值表达式相同）
		e = iota //4（恢复iota自增，计数包括c、d）
		f        //5
	)
	println(a, b, c, d, e, f)
}

func test01() {
	const (
		a         = iota //int
		b float64 = iota //float32
		//c          //int（如不显式指定iota，则与b数据类型相同）
		c = iota //int（如不显式指定iota，则与b数据类型相同）
	)
	println(a, b, c)
}
