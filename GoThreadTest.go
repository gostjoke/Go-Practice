// a go thread example
package main

import (
	"fmt"
	"sync"
	"time"
)

func CountGoat(v int, s string) {
	for i := 1; i <= v; i++ {
		fmt.Println(i, s)
		time.Sleep(100 * time.Millisecond) // Simulate some work
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		CountGoat(6, "dog")
		wg.Done()
	}()
	func() {
		CountGoat(4, "cat")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All goroutines finished executing.")
}
