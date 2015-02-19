package main

import (
	"sync"
	"testing"
)

func TestInsert(t *testing.T) {
	var test_value int64 = 100
	ls := NewConcurrentSortedList()
	
	if ls.Size() != 0 {
		t.Error("Invalid size found, should be empty.")
	}
	
	ls.Insert(test_value)
	if ls.Size() != 1 {
		t.Error("Invalid size found, should be size 1.")
	}
	if ls.Head.Next.Value != test_value {
		t.Error("Invalid value, should be", test_value)
	}
}

func TestThreadedInsert(t *testing.T) {

	sampleSize := 10000
	var wg sync.WaitGroup
	wg.Add(sampleSize)
	
	ls := NewConcurrentSortedList()
	for i := 0; i < sampleSize; i++ {
		go func() {
			ls.Insert(int64(i))
			wg.Done()
		}()
	}
	
	wg.Wait()
	
	if ls.Size() != int64(sampleSize) {
		t.Error("Invalid size found, should be size", sampleSize)
	}
}
