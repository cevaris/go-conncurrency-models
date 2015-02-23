package main

import (
	"fmt"
	"flag"
	// "sync"
	// "runtime"
	"github.com/cevaris/go_concurrency_models"
	"github.com/cevaris/go_concurrency_models/threads_locks/word_count"
	"github.com/cevaris/go_concurrency_models/threads_locks/wiki"
)

// var WORKER_SIZE int64 = 100
var SAMPLE_SIZE int64 = 10 * 1000
// var mutex *sync.Mutex = &sync.Mutex{}
var total int64 = 0
var counts map[string]int64 = make(map[string]int64)


var filePath = flag.String( "infile",
	"/data/enwiki-20150205-pages-meta-current27.xml",
	"Input file path")

func countWord(word string) {
	if val, ok := counts[word]; ok {
		counts[word] = val + 1
	} else {
		counts[word] = 1
	}
}

func pageHandler(pages <-chan *wiki.WikiPage) {
	for page := range pages {
		words := word_count.NewWords(page.GetText())
		_ = words
		for word := range words.Iterator() {
			countWord(word)
		}
	}	
}

func main() {
	// maxCPUs := runtime.NumCPU() - 1
	// runtime.GOMAXPROCS(maxCPUs)
	
	flag.Parse()
	
	file := go_concurrency_models.OpenFileOrPanic(*filePath)
	defer file.Close()

	parser := wiki.NewWikiParser(SAMPLE_SIZE, file)
	pages := parser.Parse()
	pageHandler(pages)

	// for k, v := range counts {
	// 	fmt.Println(k, v)
	// }

	// wg := &sync.WaitGroup{}
	// wg.Add(WORKER_SIZE)
	// for i := 0; i < int(wp.NumOfReaders); i++ {
	// 	go wp.ParseWorker(i)
	// }
	
	// parser.ReadWaitGroup.Wait()

	// var wg sync.WaitGroup
	// for i := 0; i < int(WORKER_SIZE); i++ {
	// 	wg.Add(1)
	// 	go parser.ParseWorker(i)
	// }
	// wg.Wait()
	// fmt.Println("Read", total, "pages")
	fmt.Println("done.")
}



