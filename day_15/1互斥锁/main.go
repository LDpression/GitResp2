package main

import (
	"fmt"
	"sync"
)

var x int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 500000; i++ {
		mutex.Lock()   //上厕所的时候先关门
		x += 1         //正在上厕所
		mutex.Unlock() //上完了厕所，再去开门
	}
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
