package main

import "fmt"

func main() {
	//primary()
	//null_pointer()
	//array_pointer()
	pointer_to_pointer()
}

func basic() {
	a := 10
	var ip *int // 指针变量, 存储变量的地址
	ip = &a
	fmt.Printf("a 变量的地址: %x\n", &a)
	fmt.Printf("ip 变量存储的指针地址: %x \n", ip)
	fmt.Printf("*ip 变量值: %d\n", *ip)

	/*
	a 变量的地址: c00008a008
	ip 变量存储的指针地址: c00008a008
	*ip 变量值: 10
	 */
}

func null_pointer() {
	var ip *int
	fmt.Println(ip)

	// <nil>
}

func array_pointer() {
	const max_value int = 3
	a := []int{10, 20, 30}

	var i int
	var ptr [max_value]*int

	for i = 0; i < max_value; i++ {
		ptr[i] = &a[i]
	}

	for i = 0; i < max_value; i++ {
		fmt.Printf("ptr[%d] = %d\n", i, *ptr[i])
	}

	/*
	ptr[0] = 10
	ptr[1] = 20
	ptr[2] = 30
	 */

}

func pointer_to_pointer() {
	a := 10
	var ptr *int
	var pptr **int

	ptr = &a
	pptr = &ptr

	fmt.Printf("变量a: %d\n", a)
	fmt.Printf("ptr=%x\n", ptr)
	fmt.Printf("pptr=%x\n", pptr)
	fmt.Printf("指针变量: *ptr=%d\n", *ptr)
	fmt.Printf("指向指针变量的指针变量 **pptr=%d", **pptr)

	/*
	变量a: 10
	ptr=c00001e090
	pptr=c00000e028
	指针变量: *ptr=10
	指向指针变量的指针变量 **pptr=10
	 */
}
