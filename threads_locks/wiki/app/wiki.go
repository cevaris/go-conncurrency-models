package main

import (
	"fmt"
	"flag"
	"sync"
	"runtime"
	"github.com/cevaris/go_concurrency_models"
	"github.com/cevaris/go_concurrency_models/threads_locks/wiki"
)

var WORKER_SIZE int64 = 100
var SAMPLE_SIZE int64 = 10 * 1000
var mutex *sync.Mutex = &sync.Mutex{}
var total int64 = 0


var filePath = flag.String( "infile",
	"/data/enwiki-20150205-pages-meta-current27.xml",
	"Input file path")



func main() {
	maxCPUs := runtime.NumCPU() - 1
	runtime.GOMAXPROCS(maxCPUs)
	
	flag.Parse()
	
	file := go_concurrency_models.OpenFileOrPanic(*filePath)
	defer file.Close()

	parser := wiki.NewWikiParser(SAMPLE_SIZE, file)
	parser.Parse()

	var wg sync.WaitGroup
	for i := 0; i < int(WORKER_SIZE); i++ {
		wg.Add(1)
		go parser.ParseWorker(i)
	}
	wg.Wait()
	fmt.Println("Read", total, "pages")
}



