package main

/*
在 64 位架构的机器上，一个切片需要 24 字节的内存: 指针字段需要8字节，长度8字节, 容量8字节

由于与切片关联的数据包含在底层数组里，不属于切片本身，所以将切片 复制到任意函数的时候，
对底层数组大小都不会有影响。复制时只会复制切片本身，不会涉及底 层数组

在函数间传递 24 字节的数据会非常快速, 这也是切片效率高的地方
 */

import (
	"fmt"
	"sort"
)

func slice1() {
	// 使用make构造slice, length=5, capacity=8
	var a = make([]int, 5, 8)
	for i := 0; i < len(a); i++ {
		a[i] = i * i
	}
	fmt.Println(a, len(a), cap(a))

	// 直接构建slice, 注意和array区别, 没有长度和...
	b := []string{"a", "b", "c"}
	fmt.Println(b, len(b), cap(b))

	/*
	[0 1 4 9 16] 5 8
	[a b c] 3 3
	 */
}

func slice2() {
	// slice共享底层元素
	a := []int64{0, 1, 2, 3}
	b := a

	// 更改一个元素
	a[0] = 999

	fmt.Println(a)
	fmt.Println(b)

	/*
	[999 1 2 3]
	[999 1 2 3]
	 */
}

func slice3() {
	// 容量变更对slice的影响
	s1 := []int64{1, 2, 3, 4, 5}
	fmt.Println(s1, len(s1), cap(s1))

	// 满容量切片追加会影响底层数组
	var s2 = append(s1, 6)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))

	// 非满容量切片追加仍然共享底层数组
	var s3 = append(s2, 7)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	// 此时s2和s3依然共享底层数组
	s2[0] = 255
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	/*
	[1 2 3 4 5] 5 5
	[1 2 3 4 5] 5 5
	[1 2 3 4 5 6] 6 10
	[1 2 3 4 5 6] 6 10
	[1 2 3 4 5 6 7] 7 10
	[255 2 3 4 5 6] 6 10
	[255 2 3 4 5 6 7] 7 10
	 */

}

func slice4() {
	var s1 = []int{1, 2, 3, 4, 5, 6, 7}
	var s2 = s1[:5]
	var s3 = s1[2:] // slice不能看到第一个元素之前的部分, 所以len=5, cap=5
	var s4 = s1[:]

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	fmt.Println(s4, len(s4), cap(s4))

	fmt.Println("change ...")
	s1[0] = 255

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	/*
	[1 2 3 4 5 6 7] 7 7
	[1 2 3 4 5] 5 7
	[3 4 5 6 7] 5 5
	[1 2 3 4 5 6 7] 7 7
	change ...
	[255 2 3 4 5 6 7]
	[255 2 3 4 5]
	[3 4 5 6 7]
	[255 2 3 4 5 6 7]
	 */
}

// append操作, 如果还在共享底层数组, 那么slice的追加也会影响array
// 在当前slice容量小于1000时, 每次扩容是原来2倍, 超过后扩大1.25倍
func slice5() {
	a := [5]int{0, 1, 2, 3, 4}

	b := a[1:3]

	fmt.Println(a)
	fmt.Println(b)

	b = append(b, 99)
	fmt.Println(a)
	fmt.Println(b)

	/*
	[0 1 2 3 4]
	[1 2]
	[0 1 2 99 4]
	[1 2 99]
	 */
}

// append追加多个元素
func slice6() {
	a := []int{0, 1, 2, 3}
	b := []int{4, 5, 6}
	c := append(a, b...)
	fmt.Println(c)

	/*
	[0 1 2 3 4 5 6]
	 */

}

// 强制限制切片容量, 设置len=cap的切片好处是发生append操作可以使slice和原有array分离, 保证不影响array中的数据
func slice7() {
	a := [4]int{0, 1, 2, 3}
	b := a[1:2:3] // 此时len = 2-1 = 1, cap = 3-1 = 2
	fmt.Println(b)
	c := a[2:3:3] // 设置len = cap = 1
	fmt.Println(c)

	/*
	[1]
	[2]
	 */

}

// range操作
func slice8() {
	a := []int{10, 20, 30}
	fmt.Println(a)
	for index, value := range a {
		fmt.Println(index, value)
	}

	// 注意此处value是原始值的副本, 并不是指向原始值的应用, 所以此处不能直接获取value的地址当做原始值地址
	for index, value := range a {
		// 迭代过程中返回的变量是依次赋值的新变量, value的地址都是一样的
		fmt.Println(index, &value, &a[index])
	}

	/*
	[10 20 30]
	0 10
	1 20
	2 30
	0 0xc00008a060 0xc00008c000
	1 0xc00008a060 0xc00008c008
	2 0xc00008a060 0xc00008c010
	 */
}

func slice_sort() {
	// 排序, 默认是升序, 原地排序
	a := []int{4, 2, 1, 3}
	sort.Ints(a)
	fmt.Println(a)

	b := []string{"b", "c", "a"}
	sort.Strings(b)
	fmt.Println(b)

	c := []int{4, 2, 1, 3}
	sort.Sort(sort.Reverse(sort.IntSlice(c)))
	fmt.Println(c)

	/*
	[1 2 3 4]
	[a b c]
	[4 3 2 1]
	 */

}

func main() {
	//slice1()
	//slice2()
	//slice3()
	//slice4()
	//slice5()
	//slice6()
	//slice7()
	slice8()
	//slice_sort()
}
