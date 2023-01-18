package main

// goroutine池

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	for j := range jobs {

		fmt.Printf("worker:%d start job:%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 100

		wg.Done()
	}
}

func printData(results <-chan int) {
	for x := range results {
		fmt.Println(x)
	}
}

func main() {
	wg := sync.WaitGroup{}

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	jobNum := 10
	wg.Add(jobNum)

	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, &wg, jobs, results)
	}
	// 5个任务
	for j := 1; j <= jobNum; j++ {
		jobs <- j
	}
	close(jobs)

	go printData(results)

	wg.Wait()
	close(results)
}
