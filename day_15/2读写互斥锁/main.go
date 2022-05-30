package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0
var wg sync.WaitGroup
var lock sync.RWMutex

func write() {
	defer wg.Done()
	lock.Lock()
	x += 1
	lock.Unlock()
}
func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println(x)
	lock.RUnlock()
}
func main() {

	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
