package main

import (
	"fmt"
	"sync"
)

// 自带的map不是线程安全的, 以下实现线程安全的字典

type SafeDict struct {
	data  map[string]int
	*sync.Mutex // 使用锁的指针类型, 避免锁被复制不起作用
}

// 新建字典
func NewSafeDict(data map[string]int) *SafeDict {
	return &SafeDict{
		data,
		&sync.Mutex{},
	}
}

func (d *SafeDict) Len() int {
	d.Lock()
	defer d.Unlock()
	return len(d.data)
}

func (d *SafeDict) Get(key string) (int, bool) {
	d.Lock()
	defer d.Unlock()
	value, ok := d.data[key]
	return value, ok
}

func (d *SafeDict) Put(key string, value int) (int, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	d.data[key] = value
	return old_value, ok
}

func (d *SafeDict) Delete(key string) (int, bool) {
	d.Lock()
	defer d.Unlock()
	old_value, ok := d.data[key]
	if ok {
		delete(d.data, key)
	}
	return old_value, ok
}

func write(d *SafeDict) {
	d.Put("banana", 5)
}

func read(d *SafeDict) {
	fmt.Println(d.Get("banana"))
}

func main() {
	d := NewSafeDict(map[string]int{
		"apple": 2,
		"pear":  3,
	})
	go read(d)
	write(d)
	read(d)
}
