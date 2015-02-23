package main

import (
	"fmt"
	"flag"
	"sync" 
	// "runtime"
	"time"

	"github.com/cevaris/go_concurrency_models"
	"github.com/cevaris/go_concurrency_models/threads_locks/word_count"
	"github.com/cevaris/go_concurrency_models/threads_locks/wiki"
)

var WORKER_SIZE int64 = 100
var SAMPLE_SIZE int64 = 100 * 1000
var mutex *sync.Mutex = &sync.Mutex{}
var counts map[string]int64 = make(map[string]int64)


var filePath = flag.String( "infile",
	"/data/enwiki-20150205-pages-meta-current27.xml",
	"Input file path")

func countWord(word string) {
	mutex.Lock()
	defer mutex.Unlock()
	
	if val, ok := counts[word]; ok {
		counts[word] = val + 1
	} else {
		counts[word] = 1
	}
}

func pageHandler(pages <-chan *wiki.WikiPage, wg *sync.WaitGroup) {
	for page := range pages {
		words := word_count.NewWords(page.GetText())
		for word := range words.Iterator() {
			countWord(word)
		}
	}
	wg.Done()
}

func main() {
	// maxCPUs := runtime.NumCPU() - 1
	// runtime.GOMAXPROCS(maxCPUs)
	start := time.Now().Unix()
	
	flag.Parse()
	
	file := go_concurrency_models.OpenFileOrPanic(*filePath)
	defer file.Close()

	parser := wiki.NewWikiParser(SAMPLE_SIZE, file)
	pages := parser.Parse()

	wg := &sync.WaitGroup{}
	wg.Add(int(WORKER_SIZE))
	for i := 0; i < int(WORKER_SIZE); i++ {
		go pageHandler(pages, wg)
	}
	wg.Wait()

	end := time.Now().Unix()
	duration := end - start
	fmt.Println(duration)
	// for k, v := range counts {
	// 	fmt.Println(k, v)
	// }

	fmt.Println("done.")
}



