package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	fmt.Println(counter)
}
