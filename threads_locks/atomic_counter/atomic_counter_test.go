package main

import (
	"sync"
	"testing"
)

func TestIncrementAndGet(t *testing.T) {
	counter := &AtomicCounter{mutex: &sync.Mutex{}}
	
	if counter.IncrementAndGet() != uint64(1) {
		t.Error("Invalid size found, should be empty.")
	}

	if counter.Count != uint64(1) {
		t.Error("Invalid size found, should be empty.")
	}
}

func TestGetAndIncrement(t *testing.T) {
	counter := &AtomicCounter{mutex: &sync.Mutex{}}
	
	if counter.GetAndIncrement() != uint64(0) {
		t.Error("Invalid size found, should be empty.")
	}

	if counter.Count != uint64(1) {
		t.Error("Invalid size found, should be empty.")
	}
}

func TestThreadedCounter(t *testing.T) {
	
	counter := &AtomicCounter{mutex: &sync.Mutex{}}

	var wg sync.WaitGroup
	wg.Add(3)
	
	go AtomicCountingThread(counter, &wg)
	go AtomicCountingThread(counter, &wg)
	go AtomicCountingThread(counter, &wg)

	wg.Wait()

	var expected uint64 = 30000
	if counter.Count != expected {
		t.Error("Invalid count found: ", counter.Count, "!=", expected)
	}
}
