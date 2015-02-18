// A _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Create thread pool of size 1
	var wg sync.WaitGroup
	wg.Add(1)
	
	go func() {
		fmt.Println("Hello from the new thread")
		wg.Done()
	}()

	fmt.Println("Hello from the main thread")
	// Wait till done
	wg.Wait()
}
