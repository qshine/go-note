package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
对一个共享资源的操作必须是原子化的
*/

var (
	// goroutine要增加的变量
	counter int64
	wg      sync.WaitGroup
)

func unLockIncCounter() {
	defer wg.Done()
	for num := 0; num < 2; num++ {
		value := counter
		// 使当前goroutine退出, 放回到调度队列
		runtime.Gosched()
		value ++
		counter = value
	}

}

// 不加锁, 有风险. 可以使用命令go build -race来检测竞争
func test_1() {
	wg.Add(2)
	go unLockIncCounter()
	go unLockIncCounter()

	wg.Wait()
	fmt.Println("Final counter: ", counter)

	/*
	Final counter:  2
	或
	Final counter:  4
	 */
}

func atomicIncCounter() {
	defer wg.Done()
	for num := 0; num < 2; num++ {
		// 强制同一时刻只能有一个goroutine能改变该值
		atomic.AddInt64(&counter, 1)
		// 使当前goroutine退出, 放回到调度队列
		runtime.Gosched()
	}
}

// 使用原子函数atomic. 原子函数能够以很底层的加锁机制来同步访问整型变量和指针
func test_2() {
	wg.Add(2)
	go atomicIncCounter()
	go atomicIncCounter()
	wg.Wait()
	fmt.Println("Final counter: ", counter)

	/*
	Final counter:  4
	 */
}

func lockIncCounter() {
	defer wg.Done()
	for num := 0; num < 2; num++ {
		// 加锁: 同一时刻只允许一个goroutine进入该临界区
		mutex.Lock()
		{
			value := counter
			// 使当前goroutine退出, 放回到调度队列
			runtime.Gosched()
			value ++
			counter = value
		}
		// 释放锁: 允许其它goroutine进入临界区
		mutex.Unlock()
	}

	/*
	Final counter:  4
	 */
}

var mutex sync.Mutex
// 使用互斥锁mutex
func test_3() {
	wg.Add(2)
	go lockIncCounter()
	go lockIncCounter()
	wg.Wait()
	fmt.Println("Final counter: ", counter)
}

func main() {
	//test_1()
	//test_2()
	test_3()
}
