package main

import "fmt"

func array1() {
	// 声明一个数组, 设置长度, 并设置零值. 注意: 不同长度的array不是同一类型
	var a [10]int64
	a[2] = 222
	fmt.Println(a)

	// 声明并初始化
	b := [3]string{"a", "b", "c"}
	fmt.Println(b)

	// 自动计算数组长度
	d := [...]float64{1.1, 2.2, 3.3, 4.4}
	fmt.Println(d)

	// 指定特定位置的值
	e := [5]int{1: 10, 4: 50}
	fmt.Println(e)

	// 访问指针数组
	f := [5]*int{0: new(int), 1: new(int)}
	*f[0] = 10
	*f[1] = 20
	fmt.Println(f)
	fmt.Println(*f[0])
	fmt.Println(*f[1])

	/*
	[0 0 222 0 0 0 0 0 0 0]
	[a b c]
	[1.1 2.2 3.3 4.4]
	[0 10 0 0 50]
	[0xc00008a020 0xc00008a028 <nil> <nil> <nil>]
	10
	20
	 */

}

// 复制指针数组, 只会复制指针的值(即地址), 不会复制指针所指向的值
func array3() {
	var a [2]*int
	b := [2]*int{new(int), new(int)}
	*b[0] = 10
	*b[1] = 20

	// 将b的值复制给a
	a = b
	fmt.Println(a)
	fmt.Println(*a[0], *a[1])
	fmt.Println(b)
	fmt.Println(*b[0], *b[1])

	*a[0] = 99
	fmt.Println(a)
	fmt.Println(*a[0], *a[1])
	fmt.Println(b)
	fmt.Println(*b[0], *b[1])

	/*
	[0xc000090040 0xc000090048]
	10 20
	[0xc000090040 0xc000090048]
	10 20
	[0xc000090040 0xc000090048]
	99 20
	[0xc000090040 0xc000090048]
	99 20
	 */
}

// 复制值数组, 更改其中一个数组的值不会影响另一个. 类似传值和传引用
func array2() {
	var a [2]int
	b := [2]int{0, 1}

	a = b
	fmt.Println(a)
	fmt.Println(b)

	a[0] = 99
	fmt.Println(a)
	fmt.Println(b)

	/*
	[0 1]
	[0 1]
	[99 1]
	[0 1]
	 */
}

func main() {
	array1()
	array2()
	array3()
}
