package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// go run -race main.go
func main() {
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	var counter int64
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	fmt.Println(counter)
}
