package main

import (
	"errors"
	"fmt"
)

/*
Go在错误处理上采用了与C类似的检查返回值的方式，而不是其他多数主流语言采用的异常方式，这造成了代码编写
上的一个很大的缺点:错误处理代码的冗余，对于这种情况是我们通过复用检测函数来减少类似的代码。
*/
func main() {
	_, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("负数不能被开平方")
	}
	return f, nil
}
