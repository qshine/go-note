package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 对比 原子操作 和 互斥锁 的性能开销

// 互斥锁操作
func test_mutex() {
	st := time.Now()
	var lock sync.Mutex

	var num int64 = 0
	for i := 0; i < 10000; i++ {
		lock.Lock()
		num++
		lock.Unlock()
	}
	fmt.Println("test_mutex: ", time.Now().Sub(st))
}

// 原子操作
func text_atomic() {
	st := time.Now()

	var num int64 = 0
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&num, 1)
	}
	fmt.Println("text_atomic: ", time.Now().Sub(st))
}

func main() {
	for i := 0; i < 10; i++ {
		test_mutex()
	}

	fmt.Println("----")

	for i := 0; i < 10; i++ {
		text_atomic()
	}
}

/*

test_mutex:  121.971µs
test_mutex:  123.655µs
test_mutex:  123.145µs
test_mutex:  123.395µs
test_mutex:  124.075µs
test_mutex:  136.818µs
test_mutex:  124.19µs
test_mutex:  127.375µs
test_mutex:  126.153µs
test_mutex:  125.244µs
----
text_atomic:  56.214µs
text_atomic:  56.441µs
text_atomic:  51.272µs
text_atomic:  52.399µs
text_atomic:  57.52µs
text_atomic:  56.049µs
text_atomic:  56.495µs
text_atomic:  57.203µs
text_atomic:  56.385µs
text_atomic:  56.185µs

*/
