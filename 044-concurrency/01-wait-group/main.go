package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS:", runtime.GOOS)
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	wg.Add(1)

	go foo()
	bar()
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GOROUTINES:", runtime.NumGoroutine())

	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar", i)
	}
}
