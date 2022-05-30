/*
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)
var mutex sync.Mutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			mutex.Lock() //这里如果不加上锁的话，就会报：并发map写入的错误
			set(key, n)
			mutex.Unlock()
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

*/
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m2 = sync.Map{}
var wg sync.WaitGroup

//开箱即用
func main() {

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)
			value, _ := m2.Load(key)
			fmt.Println("value:", value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
