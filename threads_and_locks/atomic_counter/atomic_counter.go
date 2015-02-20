package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	Count uint64
	mutex *sync.Mutex
}

// Increment count, return updated value
func (c *AtomicCounter) IncrementAndGet() uint64 {
	return atomic.AddUint64(&c.Count, 1)
}

// Increment count, return previous value
func (c *AtomicCounter) GetAndIncrement() uint64 {
	old := c.Count
	atomic.AddUint64(&c.Count, 1)
	return old
}


func AtomicCountingThread(counter *AtomicCounter, wg *sync.WaitGroup) {
	// Signal done after function execution
	defer wg.Done()

	for i := 0; i < 10000; i++ {
		counter.GetAndIncrement()
	}
}

func main() {
	// Create Counter instance with Mutex
	counter := &AtomicCounter{mutex: &sync.Mutex{}}

	// Create thread pool of size 2
	var wg sync.WaitGroup
	wg.Add(2)
	
	// Launch 2 threads with reference to the same Counter
	go AtomicCountingThread(counter, &wg)
	go AtomicCountingThread(counter, &wg)

	// Wait for counters to finish
	wg.Wait()
	fmt.Println(counter.Count)
} 
