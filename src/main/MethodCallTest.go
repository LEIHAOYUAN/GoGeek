package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("发送邮件给 %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	//user类型的值可以调用
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	//指向user类型值的指针也可以调用
	lisa := &user{"Lisa", "lisa@mailcom"}
	lisa.changeEmail("testXXX")
	bill.notify()

	//user 类型的值也可以调用
	bill.changeEmail("bill@newdomian.com")
	bill.notify()
}
