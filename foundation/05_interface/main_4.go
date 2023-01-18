package main

import (
	"fmt"
)

/*
利用interface实现多态
*/

type notifier2 interface {
	notify()
}

type user1 struct {
	name string
	age  int
}

func (u *user1) notify() {
	fmt.Println(u.name, u.age)
}

type admin struct {
	name string
	age  int
}

func (a *admin) notify() {
	fmt.Println(a.name, a.age)
}

func main() {
	tom := user1{"Tom", 10}
	sendNotification(&tom)

	bob := admin{"Bob", 40}
	sendNotification(&bob)
}

// 声明多态函数, 接受一个实现notifier2接口的值并调用notify方法
func sendNotification(n notifier2) {
	n.notify()

	/*
	Tom 10
	Bob 40
	 */
}
