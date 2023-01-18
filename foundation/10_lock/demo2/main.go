package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

// 测试互斥锁
func test_mutex() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 模拟写操作
		go func() {
			lock.Lock()
			x = x + 1
			time.Sleep(10 * time.Millisecond)
			lock.Unlock()
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			time.Sleep(time.Millisecond)
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("互斥锁: %s", end.Sub(start))
	//fmt.Println(end.Sub(start))
}

// 测试读写锁
func test_rwmutex() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// 模拟写操作
		go func() {
			rwlock.Lock()
			x = x + 1
			time.Sleep(10 * time.Millisecond)
			rwlock.Unlock()
			wg.Done()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			rwlock.RLock()
			time.Sleep(time.Millisecond)
			rwlock.RUnlock()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Printf("读写锁: %s", end.Sub(start))
}

func main() {
	test_mutex()
	fmt.Println()
	test_rwmutex()
}

/*
互斥锁: 1.415241446s
读写锁: 115.773867ms
*/
