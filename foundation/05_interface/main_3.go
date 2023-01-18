package main

import "fmt"

// 定义一个通知类行为的接口
type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

// 使用指针接收者实现该方法
func (u *user) notify(){
	fmt.Printf("Sending to %s, email is %s", u.name, u.email)
}

// 接受一个实现了notifier接口的值, 发送通知
func sendEmails(n notifier) {
	n.notify()
}


func main() {
	// 创建一个user类型的值并发送通知
	u := user{
		name: "bob",
		email: "bob.com",
	}
	/*
	这样会报错, 使用sendEmails必须传入一个指针类型, 因为user的notify方法是指针类型
	cannot use u (type user) as type notifier in argument to sendEmails:
	user does not implement notifier (notify method has pointer receiver)

	解释: user的notify方法是一个 pointer receiver, 也就是说user类型的值并没有实现notifier接口, 而是指针实现了该接口
	*/
	//sendEmails(u)
	sendEmails(&u)

	/*
	Sending to bob, email is bob.com
	 */
}


