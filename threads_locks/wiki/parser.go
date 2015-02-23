package wiki

import (
	"encoding/xml"
	"fmt"
	"os"
	// "sync"
)

var PAGE_ELEMENT string = "page"

type WikiParser struct {
	FileHandler *os.File
	Pages chan *WikiPage
	NumToParse int64
	// NumOfReaders int64
	ReadBufferSize int
	// ReadMutex *sync.Mutex
	TotalParsed int64
}

func NewWikiParser(numOfPages int64, file *os.File) *WikiParser {
	return &WikiParser{
		FileHandler: file,		
		// ReadMutex: &sync.Mutex{},
		NumToParse: numOfPages,
		ReadBufferSize: 500,
		TotalParsed: 0,
	}
}

func (wp *WikiParser) Parse() <-chan *WikiPage {
	// Create new buffered channel
	wp.Pages = make(chan *WikiPage, wp.ReadBufferSize)
	// Launch goroutine to read, parse, and enqueue Page to a channel
	go parseRoutine(wp)
	// Return read-only channel for iterating
	return wp.Pages
}

func parseRoutine(wp *WikiParser){
	decoder := xml.NewDecoder(wp.FileHandler)
	for wp.hasNext() {

		t, _ := decoder.Token()
		if t == nil {
			break
		}

		// Get element type
		switch se := t.(type) {
		case xml.StartElement:
			// If start of page element
			if se.Name.Local == PAGE_ELEMENT {
				var p WikiPage
				// Parse page element
				decoder.DecodeElement(&p, &se)
				// enqueue Page to channel
				wp.Pages <- &p
			}
		default:
			continue
		}
	}
	// Close channel after done reading the file
	close(wp.Pages)
}

// Returns true if we have not reached our target page count
func (wp *WikiParser) hasNext() bool {
	// wp.ReadMutex.Lock()
	// defer wp.ReadMutex.Unlock()
	// Check if reached limit
	fmt.Printf("\r%d/%d", wp.TotalParsed, wp.NumToParse)
	if  wp.TotalParsed >= wp.NumToParse {
		return false
	}
	wp.TotalParsed++
	return true
}


