package main

import (
	"fmt"
	"runtime"
	"sync"
)

//! Goroutines incrementer
func main() {
	fmt.Println("Start goroutine:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	incrementer := 0
	const gs = 100

	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
		fmt.Println("Middle goroutine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Counted:", incrementer)
}
