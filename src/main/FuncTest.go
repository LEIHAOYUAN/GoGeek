package main

// 在函数调用前，会为形参和返回值分配空间，并将实参拷贝到形参内存
func main() {

}

// 定义函数类型
type TestFunc func(string, ...interface{}) (string, error)

func format(f TestFunc, s string, a ...interface{}) (string, error) {
	return f(s, a...)
}
