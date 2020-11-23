package main

func main() {

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
