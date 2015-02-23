package wiki

import (
	"encoding/xml"
	// "fmt"
	"os"
)

var PAGE_ELEMENT string = "page"

type WikiParser struct {
	FileHandler *os.File
	Pages chan *WikiPage
	NumToParse int64
	ReadBufferSize int
	TotalParsed int64
}

func NewWikiParser(numOfPages int64, file *os.File) *WikiParser {
	return &WikiParser{
		FileHandler: file,		
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
	// While we have not reached our target Page count
	for wp.TotalParsed <= wp.NumToParse {
		// fmt.Printf("\r%d/%d", wp.TotalParsed, wp.NumToParse)
		wp.TotalParsed++

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
