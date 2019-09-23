package main

import (
	"fmt"
	"runtime"
	"sync"
)

// go run -race main.go
func main() {
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())
	fmt.Println(counter)
}
