package main

import "fmt"

type student struct {
	name  string
	email string
}

func (u student) notify() {
	fmt.Printf("发送邮件给 %s<%s>\n", u.name, u.email)
}

func (u *student) changeEmail(email string) {
	u.email = email
}

func main() {
	//student类型的值可以调用
	bill := student{"Bill", "bill@email.com"}
	bill.notify()

	//指向student类型值的指针也可以调用
	lisa := &student{"Lisa", "lisa@mailcom"}
	lisa.changeEmail("testXXX")
	bill.notify()

	//student 类型的值也可以调用
	bill.changeEmail("bill@newdomian.com")
	bill.notify()
}
