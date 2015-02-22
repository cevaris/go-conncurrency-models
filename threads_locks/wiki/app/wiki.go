package main

// http://blog.davidsingleton.org/parsing-huge-xml-files-with-go

import (
	"fmt"
	"flag"
	"github.com/cevaris/go_concurrency_models/threads_locks/wiki"
)

var filePath = flag.String("infile", "/data/enwiki-20150205-pages-meta-current27.xml", "Input file path")

func main() {
	flag.Parse()
	
	file, _ := wiki.OpenFileHandler(*filePath)
	defer file.Close()

	wiki.Parse(file)

	fmt.Println("Done")
}


