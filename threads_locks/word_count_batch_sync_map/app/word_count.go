
package main

import (
	"fmt"
	"flag"
	// "math"
	"sync" 
	"runtime"

	"github.com/cevaris/go_concurrency_models"
	"github.com/cevaris/go_concurrency_models/threads_locks/word_count"
	"github.com/cevaris/go_concurrency_models/threads_locks/word_count_batch_sync_map"
	"github.com/cevaris/go_concurrency_models/threads_locks/wiki"
)

var WORKER_SIZE int64 = 2500
var SAMPLE_SIZE int64 = 100 * 1000
// var SAMPLE_SIZE int64 = math.MaxInt64
var mutex *sync.Mutex = &sync.Mutex{}
var counts map[string]int64 = make(map[string]int64)

var startTime int64

var filePath = flag.String( "infile",
	"/data/enwiki-20150205-pages-meta-current27.xml",
	"Input file path")


func pageHandler(pages <-chan *wiki.WikiPage, wg *sync.WaitGroup) {
	// Each goroutine gets it own map counter
	c := word_count_batch_sync_map.NewCounter()

	for page := range pages {
		words := word_count.NewWords(page.GetText())
		for word := range words.Iterator() {
			c.CountWord(word)
		}
	}

	// Need to sync local counts to source, need to protect
	// main map counts variable.
	mutex.Lock()
	c.MergeMap(counts)
	

	wg.Done()
}

func main() {
	// Start, stop and print runtime duration
	wallClock := go_concurrency_models.NewWallClock()
	wallClock.StartClock()
	defer wallClock.StopClock()
	// Use all but 2 CPU to execute work in parallel
	maxCPUs := runtime.NumCPU() - 2
	runtime.GOMAXPROCS(maxCPUs)

	flag.Parse()

	file := go_concurrency_models.OpenFileOrPanic(*filePath)
	defer file.Close()

	parser := wiki.NewWikiParser(SAMPLE_SIZE, file)
	parser.ReadBufferSize = 10000
	pages := parser.Parse()

	// Create worker pool and wait till work is completed
	wg := &sync.WaitGroup{}
	wg.Add(int(WORKER_SIZE))
	for i := 0; i < int(WORKER_SIZE); i++ {
		go pageHandler(pages, wg)
	}
	wg.Wait()

	// Print all results
	for k, v := range counts {
		fmt.Println(k, v)
	}

	fmt.Println("done.")
}



