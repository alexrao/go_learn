package main

import (
	"fmt"
	"runtime"
	"sync"
)

var Counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	Counter++
	fmt.Println(Counter)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go Count(lock)
	}

	for {
		lock.Lock()
		c := Counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
