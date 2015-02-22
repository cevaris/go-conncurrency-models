package main

import (
	"fmt"
	"flag"
	"os"
	"sync"
	"runtime"
	"github.com/cevaris/go_concurrency_models/threads_locks/wiki"
)

var WORKER_SIZE uint64 = 100
var SAMPLE_SIZE uint64 = 100 * 1000
var mutex *sync.Mutex = &sync.Mutex{}
var total uint64 = 0


var filePath = flag.String("infile", "/data/enwiki-20150205-pages-meta-current27.xml", "Input file path")


func HandlePage(id int, pageIter <-chan wiki.Page, wg *sync.WaitGroup) {
	for page := range pageIter {

		mutex.Lock()
		total++
		_ = page

		if total % 1000 == 0 {
			fmt.Printf("\r%d\t%d", id, total)
		}
		// Need to offset the number of goroutines which have already
		// grabbed a page off the channel
		if total > SAMPLE_SIZE - WORKER_SIZE {
			mutex.Unlock()
			break
		} 

		mutex.Unlock()
	}
	wg.Done()
}
func main() {
	maxCPUs := runtime.NumCPU() - 1
	runtime.GOMAXPROCS(maxCPUs)
	
	flag.Parse()
	
	file, fileErr := wiki.OpenFileHandler(*filePath)
	defer file.Close()
	if fileErr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", fileErr)
		os.Exit(1)
	}

	parser := wiki.NewWikiParser()
	pageIter := parser.Parse(file)

	var wg sync.WaitGroup
	for i := 0; i < int(WORKER_SIZE); i++ {
		wg.Add(1)
		go HandlePage(i, pageIter, &wg)
	}
	wg.Wait()
	fmt.Println("Read", total, "pages")
}



