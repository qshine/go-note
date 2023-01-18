package main

import "fmt"

// if: 后面只能是bool类型, 不能是int型
func testIf() {
	a := 86
	if a > 90 {
		fmt.Println("A")
	} else if (a > 70 && a <= 90) {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// B
}

// switch
func testSwitch1() {
	level, scores := "B", 90

	switch scores {
	case 90:
		level = "A"
	case 80:
		level = "B"
	case 70, 60:
		level = "C"
	default:
		level = "D"
	}

	switch {
	case level == "A":
		fmt.Println("very good")
	case level == "B":
		fmt.Println("good")
	case level == "C":
		fmt.Println("normal")
	case level == "D":
		fmt.Println("bad")
	}
	fmt.Println(level)

	/*
	very good
	A
	 */
}

// 循环
func testForLoop() {
	// 普通循环
	a := []int{0, 1, 2, 3, 4, 5}
	for index, value := range a {
		fmt.Println(index, value)
	}

	/*
	0 0
	1 1
	2 2
	3 3
	4 4
	5 5
	 */

	// 死循环
	var b int = 0
	for {
		b += 1
		if b == 3 {
			break
		}
	}
	fmt.Println(b) // 3

	// 直接跟表达式
	var c int = 0
	for c <= 3 {
		c += 1
	}
	fmt.Println(c) // 4

}

// goto 语句, 可以实现条件转移
func testGoto() {
	a := 10

LOOP:
	for a < 15 {
		if a == 13 {
			a += 1
			// 如果==13 +1 后直接跳到LOOP处进行下一次循环
			goto LOOP
		}
		fmt.Println(a)
		a ++
	}

	/*
	10
	11
	12
	14
	 */
}

func main() {
	//testIf()
	//testForLoop()
	//testSwitch1()
	testGoto()
}
