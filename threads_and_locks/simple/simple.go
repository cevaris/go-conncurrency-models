package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	
	go func() {
		defer wg.Done()
		fmt.Println("Hello from the new thread")
	}()

	fmt.Println("Hello from the main thread")
	wg.Wait()
	
}
