package main

import "fmt"

type User struct {
	name string
	age  int
}

// 使用值 receiver 实现一个方法, 调用时会使用值的一个副本
func (u User) notify() {
	fmt.Println(u.name, u.age)
}

// 该方法并不能改变外层的值
func (u User) changeAge1(age int) {
	u.age = age
	fmt.Println(u.name, u.age)
}

// 使用指针 receiver 实现一个方法, 调用时只传指针
func (u *User) changeAge2(age int) {
	u.age = age
}

func main() {
	// User类型的值可以调用方法
	test_user := User{
		name: "zhang",
		age:  20,
	}
	test_user.notify()
	test_user.changeAge1(99)
	test_user.notify()

	// User类型的指针也能调用方法, go编译器背后其实变为了 (*test_user2).notify(), 指针被解引用为值
	test_user2 := &User{
		name: "li",
		age:  40,
	}
	test_user2.notify()
	test_user2.changeAge1(50)
	test_user2.notify()
	test_user2.changeAge2(100)
	test_user2.notify()

	/*
	zhang 20
	zhang 99
	zhang 20
	li 40
	li 50
	li 40
	li 100
	 */

}
