package main

// 使用函数参数的默认值

import "fmt"

type Option struct {
	timeout int8
	retries int8
}

// 先定义一个函数类型
type OptionFunc func(*Option)

// 用来修改值
func WithTimeout(t int8) OptionFunc {
	return func(o *Option) {
		o.timeout = t
	}
}

// 用来修改值
func WithRetries(r int8) OptionFunc {
	return func(o *Option) {
		o.retries = r
	}
}

// 定义一个默认的Option
var defaultOption = &Option{
	timeout: 3,
	retries: 2,
}

// 构造函数
func NewOption(opts ...OptionFunc) (opt *Option) {
	opt = defaultOption

	for _, o := range opts {
		o(opt)
	}
	return
}

func main() {
	x := NewOption()
	fmt.Printf("%+v\n", x)

	x = NewOption(
		WithRetries(5),
	)
	fmt.Printf("%+v\n", x)
}

/*
&{timeout:3 retries:2}
&{timeout:3 retries:5}
*/
