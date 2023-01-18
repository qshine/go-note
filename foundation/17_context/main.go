package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
}

// 使用waitGroup控制goroutine
func impWait() {
	st := time.Now()
	wg := sync.WaitGroup{}
	for _, num := range []int{1, 2, 3, 4, 5} {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(2 * time.Second)
		}(num)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(st))

	/*
		5
		4
		1
		2
		3
		2.003851909s
	*/
}

// 通过channel通知来结束goroutine
func impStop() {
	// 定义一个channel
	stop := make(chan bool)
	go func() {
		for {
			select {
			// 如果接收到停止信号就结束
			case <-stop:
				fmt.Println("收到停止信号 !")
				return
			default:
				fmt.Println("持续运行中 ...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)

	// 发送停止信号
	stop <- true

	// 如果后续没有输出则说明停止了
	time.Sleep(3 * time.Second)

	/*
		持续运行中 ...
		持续运行中 ...
		持续运行中 ...
		持续运行中 ...
		持续运行中 ...
		收到停止信号 !
	*/
}

// 使用Context来控制goroutine
func impCancel0() {
	ctx, cancel := context.WithCancel(context.Background())

	// 模拟调用3个超时接口
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("timeout end !")
		}
	}(ctx)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("timeout end !")
		}
	}(ctx)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("timeout end !")
		}
	}(ctx)

	// 超时后直接取消所有goroutine
	time.Sleep(3 * time.Second)
	// 通知结束goroutine
	cancel()

	time.Sleep(1 * time.Second)
	/*
		timeout end !
		timeout end !
		timeout end !
	*/
}

func watch(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(num, " 收到结束信号!")
			return
		default:
			fmt.Println(num, " 运行 ...")
			time.Sleep(1 * time.Second)
		}
	}
}

// 控制多个goroutine
func impCancel1() {
	ctx, cancel := context.WithCancel(context.Background())

	// 开启5个goroutine
	go watch(ctx, 1)
	go watch(ctx, 2)
	go watch(ctx, 3)
	go watch(ctx, 4)
	go watch(ctx, 4)

	time.Sleep(5 * time.Second)

	// 发送停止信号
	cancel()

	time.Sleep(5 * time.Second)

	/*

		3  运行 ...
		2  运行 ...
		4  运行 ...
		1  运行 ...
		4  运行 ...
		3  运行 ...
		2  运行 ...
		1  运行 ...
		4  运行 ...
		4  运行 ...
		3  运行 ...
		1  运行 ...
		4  运行 ...
		4  运行 ...
		2  运行 ...
		2  运行 ...
		4  运行 ...
		1  运行 ...
		3  运行 ...
		4  运行 ...
		4  运行 ...
		3  运行 ...
		4  运行 ...
		1  运行 ...
		2  运行 ...
		3  收到结束信号!
		1  收到结束信号!
		4  收到结束信号!
		4  收到结束信号!
		2  收到结束信号!
	*/

}

// 使用WithValue传递数据
func impValue() {
	key := "name"
	// 新建一个子context并附加值
	ctx := context.WithValue(context.Background(), key, "xxxx")

	go func() {
		val := ctx.Value(key)
		fmt.Println(val, "g1 运行中...")
	}()
	go func() {
		val := ctx.Value(key)
		fmt.Println(val, "g2 运行中...")
	}()
	go func() {
		val := ctx.Value(key)
		fmt.Println(val, "g3 运行中...")
	}()

	time.Sleep(2 * time.Second)

	/*
		xxxx g1 运行中...
		xxxx g2 运行中...
		xxxx g3 运行中...
	*/

}

// WithTimeout  子goroutine不超时
func impWithTimeout0() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func(d time.Duration) {
		select {
		case <-ctx.Done():
			fmt.Println("handle: ", ctx.Err())
		case <-time.After(d):
			fmt.Println("process request with: ", d)
		}
	}(500 * time.Millisecond)

	time.Sleep(5 * time.Second)

	/*
		process request with:  500ms
	*/
}

// impWithTimeout1 设置超时
func impWithTimeout1() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 子goroutine因为执行时间过长, 会收到主goroutine的错误结束
	go func(d time.Duration) {
		select {
		case <-ctx.Done():
			fmt.Println("handle: ", ctx.Err())
		case <-time.After(d):
			fmt.Println("process request with: ", d)
		}
	}(2 * time.Second)

	time.Sleep(5 * time.Second)

	/*
		handle:  context deadline exceeded
	*/
}
