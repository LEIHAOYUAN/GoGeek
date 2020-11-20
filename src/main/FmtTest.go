package main

import "fmt"

// https://www.cnblogs.com/jiangchunsheng/p/10725053.html
func main() {
	fmt.Print("a", "\n")      //输出a
	fmt.Print("a", "b", "\n") //输出ab
	fmt.Print('a', "\n")      //输出97
	fmt.Print('a', 'b', "\n") //输出97 98   字符之间会输出一个空格
	fmt.Print(12, "\n")       //输出12
	fmt.Print(12, 13, "\n")   //输出12 13   数值之间输出一个空格

}
