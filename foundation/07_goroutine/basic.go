package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
goroutine基础: 并发显示消息


针对一个需要较长运行时间的goroutine, 调度器会适时将该goroutine放回调度队列, 之后再拿出继续执行

以下代码使用time.Sleep模拟耗时任务
 */

func main() {
	// 使用一个逻辑处理器
	runtime.GOMAXPROCS(1)
	//runtime.GOMAXPROCS(runtime.NumCPU())

	// 使用wg来等待goroutine完成, WaitGroup是一个计数信号量
	var wg sync.WaitGroup
	// 表示等待2个goroutine
	wg.Add(2)

	fmt.Println("start ...")

	// 声明匿名函数, 开启第一个goroutine
	go func() {
		// 在函数退出时通知该goroutine已经结束, 类似-1
		defer wg.Done()

		for i := 0; i < 3; i++ {
			fmt.Printf("First goroutine: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 开启第二个goroutine
	go func() {
		defer wg.Done()
		num := 0
		for {
			num += 1
			if num > 3 {
				break
			}
			fmt.Printf("Second goroutine: %d\n", num)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("Waiting ...")
	// 等待2个goroutine完成, 如果值大于0会阻塞
	wg.Wait()
	fmt.Println("All goroutine done!")

	/*
	start ...
	Waiting ...
	Second goroutine: 1
	First goroutine: 0
	First goroutine: 1
	Second goroutine: 2
	Second goroutine: 3
	First goroutine: 2
	All goroutine done!
	 */
}
