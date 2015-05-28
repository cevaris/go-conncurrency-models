package sum

import (
	// "fmt"
	"sync"
	"sync/atomic"
	"runtime"
)





func IntBufferedIter(ints *[]int, buffSize int) (<-chan int){
	var output chan int = make(chan int, buffSize)
	go func() {
		for _, val := range *ints {
			output <- val
		}
		close(output)
	}()
	return output
}

func syncSumRoutine(total uint64, ints <-chan int, wg *sync.WaitGroup) {
	for v := range ints {
		atomic.AddUint64(&total, uint64(v))
	}
	wg.Done()
}

func SyncSum(numbers *[]int, numRoutines int, numCPUs int, buffSize int) int64 {
	runtime.GOMAXPROCS(numCPUs)
	wg := &sync.WaitGroup{}
	var total uint64 = 0
	var numbersChan <-chan int

	if buffSize > 0 {
		numbersChan = IntBufferedIter(numbers, buffSize)
	} else {
		numbersChan = IntBufferedIter(numbers, 1)
	}

	for i := 0; i < numRoutines; i++ {
		go syncSumRoutine(total, numbersChan, wg)
	}

	wg.Wait()
	return int64(total)
}


func SimpleSum(numbers *[]int) int64 {
	var total int64 = 0
	for _, v := range *numbers {
		total += int64(v)
	}
	return total
}
