package main

// 并发安全的单例模式

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var instance *singleton

var once sync.Once

func GetIns() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	instance := GetIns()
	fmt.Println(instance)
}
