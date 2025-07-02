// Race Condition

package main

import (
	"fmt"
	"sync"
	"time"
)

func EditAdd(a *int) { //mu *sync.Mutex
	// mu.Lock()
	*a += 1
	// mu.Unlock()
}

func EditSub(a *int) { //mu *sync.Mutex
	// mu.Lock()
	*a -= 1
	// mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	// var mu sync.Mutex // Mutex to protect shared variable 'a'
	var a int = 1000

	for i := 0; i < 1000; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			EditAdd(&a)
			// EditAdd(&a, &mu)
			time.Sleep(10 * time.Millisecond) // Simulate some work
		}()
		go func() {
			defer wg.Done()
			EditSub(&a)
			// EditSub(&a, &mu)
			time.Sleep(10 * time.Millisecond) // Simulate some work
		}()

	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("final a =", a)
}
