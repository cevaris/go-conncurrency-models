package main

import (
	"fmt"
	"flag"
	"github.com/cevaris/go_concurrency_models"
	"github.com/cevaris/go_concurrency_models/threads_locks/word_count"
	"github.com/cevaris/go_concurrency_models/threads_locks/wiki"
)

var SAMPLE_SIZE int64 = 100 * 1000
var total int64 = 0
var counts map[string]int64 = make(map[string]int64)


var filePath = flag.String("infile",
	"/data/enwiki-20150205-pages-meta-current27.xml",
	"Input file path")

// Using a Map, increment count on every word occurrence
func countWord(word string) {
	if val, ok := counts[word]; ok {
		counts[word] = val + 1
	} else {
		counts[word] = 1
	}
}

// Read pages off WikiPage parser channel
func pageHandler(pages <-chan *wiki.WikiPage) {
	for page := range pages {
		words := word_count.NewWords(page.GetText())
		for word := range words.Iterator() {
			countWord(word)
		}
	}	
}

func main() {
	// Start, stop and print runtime duration
	wallClock := go_concurrency_models.NewWallClock()
	wallClock.StartClock()
	defer wallClock.StopClock()
	
	flag.Parse()
	
	file := go_concurrency_models.OpenFileOrPanic(*filePath)
	defer file.Close()

	parser := wiki.NewWikiParser(SAMPLE_SIZE, file)
	pages := parser.Parse()
	pageHandler(pages)

	// for k, v := range counts {
	// 	fmt.Println(k, v)
	// }
	
	fmt.Println("done.")
}



