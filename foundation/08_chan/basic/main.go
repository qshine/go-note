package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func test_1() {
	// 声明一个bool型的channel
	c := make(chan bool)
	go func() {
		fmt.Println("go go go")
		// 把true放入channel
		c <- true
	}()
	// 因为没有内容, 会在这里阻塞, 直到拿出true, 拿出后主程序结束, 并不需要close掉channel
	<-c

	// go go go
}

func test_2() {
	c := make(chan bool)
	go func() {
		fmt.Println("go go go")
		c <- true
		close(c) // 关闭channel后for循环才会结束
	}()

	// 使用range对chan循环
	for v := range c {
		fmt.Println(v)
	}

	/*
	go go go
	true
	 */
}

func test_sum(wg *sync.WaitGroup, channel chan int) {
	defer wg.Done()
	sum := 1
	for i := 0; i < 100000; i++ {
		sum += i
	}
	channel <- sum
}

// 使用缓冲型通道
func test_3() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	channel := make(chan int, 10)

	wg := sync.WaitGroup{}
	// 添加10个任务
	wg.Add(10)

	// 开启10个goroutine
	for i := 0; i < 10; i++ {
		go test_sum(&wg, channel)
	}

	wg.Wait()
	// 必须先关闭
	close(channel)

	// 关闭channel后循环才会结束
	for res := range channel {
		fmt.Println(res)
	}

	/*
	4999950001
	4999950001
	4999950001
	4999950001
	4999950001
	4999950001
	4999950001
	4999950001
	4999950001
	4999950001
	 */
}

// 使用select同时操作多个通道的读写
func test_4() {
	c1, c2 := make(chan int), make(chan string)
	signal := make(chan bool) // 信号通道

	go func() {
		for {
			// select只能接收一次, 使用无限循环来不断接收
			select {
			case v, ok := <-c1:
				if !ok {
					signal <- true
					break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					signal <- true
					break
				}
				fmt.Println("c2", v)
			}
		}

	}()

	fmt.Println("Start ...")

	c1 <- 1
	c2 <- "Hello"
	c1 <- 3
	c2 <- "world"

	close(c2)
	// 阻塞
	<-signal
	fmt.Println("End !")

	/*
	Start ...
	c1 1
	c2 Hello
	c1 3
	c2 world
	End !
	 */

}

// 使用select作为生产者
func test_5() {
	channel := make(chan int)
	go func() {
		for v := range channel {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case channel <- 0:
		case channel <- 1:
		}
	}

	/*
	0
	1
	1
	1
	1
	0
	1
	1
	0
	1
	 */
}

// select 读取设置超时
func test_6() {
	channel := make(chan int)
	signal := make(chan bool)
	go func() {
		for {
			select {
			case v := <-channel:
				fmt.Println(v)
			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				signal <- true
				break

			}
		}
	}()

	<-signal

	// timeout
}

// 斐波那契数列
func test_fib() {
	n := 10
	channel := make(chan int)
	go func() {
		// 一定要关闭通道, 否则最后会报错
		defer close(channel)

		x, y := 1, 1
		for i := 0; i < n; i++ {
			channel <- x
			x, y = y, x+y
		}
	}()

	for j := range channel {
		fmt.Printf("%d,", j)
	}

	// 1,1,2,3,5,8,13,21,34,55,
}

// 开启一个goroutine池
func test_7() {
	// 设置要开启的goroutine数
	const numberGoroutines = 4
	// 任务数
	const taskLoad = 10

	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	var wg sync.WaitGroup
	wg.Add(numberGoroutines)
	// 启动goroutine来开始消费
	for i := 0; i < numberGoroutines; i++ {
		go worker(tasks, i, &wg)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// 在close之后还能读取, 但是读取完成后再次读取会出错, 所以可以马上进行关闭
	close(tasks)
	// 等待所有工作完成
	wg.Wait()
}

// 从chan中不断取出任务消费
func worker(tasks chan string, worker int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			// 说明chan被关闭了
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)
		time.Sleep(2 * time.Second)
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}

func main() {
	test_1()
	test_2()
	test_3()
	test_4()
	test_5()
	test_6()
	test_fib()
	test_7()
}
