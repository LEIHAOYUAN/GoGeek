package main

import "fmt"

/*
1、方法
方法是与对象实例绑定的特殊函数
普通函数专注于算法流程，通过接受参数来完成特定的逻辑运算，并返回最终结果
方法是有关联状态的，而函数通常没有

2、语法区别
方法和函数语法区别在于方法有前置实例接受参数（receiver），编译器依次确定方法所属类型
（在其他语言中，尽管没有显示的定义，但在方法调用的时候饮食传递this实例参数）

可以为当前包，以及除了接口和指针以外的任何类型定义方法

3、指针调用、值调用
方法可看作特殊的函数，那么receiver的类型自然可以是基础类型或指针类型。这会关系到调用时对象实例是否被复制。
可使用实例值或指针调用方法，编译器会根据方法receiver类型自动在基础类型和指针类型间转换。

4、选择方法调用类型
要修改实例状态，用*T。
无须修改状态的小对象或固定值，建议用T。
大对象建议用*T，以减少复制成本。
引用类型、字符串、函数等指针包装对象，直接用T。
若包含Mutex等同步字段，用*T，避免因复制造成锁操作无效。
其他无法确定的情况，都用*T。
*/
func main() {
	var a N = 25
	println(a.toString())
}

type N int

func (n N) toString() string {
	return fmt.Sprintf("%#x", n)
}