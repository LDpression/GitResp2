package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup

func add() {
	atomic.AddInt64(&x, 1) //使用了原子类型的，就可以不用上锁了
	wg.Done()
}
func main() {

	for i := 0; i < 500000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
