package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Count int
	mutex *sync.Mutex
}

func (c *Counter) Increment()  {
	// Use Mutex to lock down access to the Count variable
	c.mutex.Lock()
	// Unlock after function execution
	defer c.mutex.Unlock()
	// Access mutable state here
	c.Count += 1	
}

func CountingThread(counter *Counter, wg *sync.WaitGroup) {
	// Signal done after function execution
	defer wg.Done()
	
	for i := 0; i < 10000; i++ {
		counter.Increment()
	}
}

func main() {
	// Create Counter instance with Mutex
	counter := &Counter{mutex: &sync.Mutex{}}

	// Create thread pool of size 2
	var wg sync.WaitGroup
	wg.Add(2)
	
	// Launch 2 threads with reference to the same Counter
	go CountingThread(counter, &wg)
	go CountingThread(counter, &wg)

	// Wait for counters to finish
	wg.Wait()
	fmt.Println(counter.Count)
}
