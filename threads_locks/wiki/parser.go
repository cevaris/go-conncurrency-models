package wiki



import (
	"encoding/xml"
	"fmt"
	"os"
	"sync"
)

var PAGE_ELEMENT string = "page"

type WikiParser struct {
	FileHandler *os.File
	Pages chan *WikiPage
	NumToParse int64
	NumOfReaders int64
	ReadBufferSize int
	ReadMutex *sync.Mutex
	ReadWaitGroup *sync.WaitGroup
	TotalParsed int64
}

func NewWikiParser(numOfPages int64, file *os.File) *WikiParser {
	return &WikiParser{
		FileHandler: file,		
		ReadMutex: &sync.Mutex{},
		ReadWaitGroup: &sync.WaitGroup{},
		NumToParse: numOfPages,
		ReadBufferSize: 500,
		TotalParsed: 0,
		NumOfReaders: 100,
	}
}

func (wp *WikiParser) Parse() <-chan *WikiPage {
	wp.Pages = make(chan *WikiPage, wp.ReadBufferSize)

	wp.ReadWaitGroup.Add(int(wp.NumOfReaders))
	for i := 0; i < int(wp.NumOfReaders); i++ {
		go wp.ParseWorker(i)
	}
	
	go parseRoutine(wp)
	return wp.Pages
}

func parseRoutine(wp *WikiParser){
	decoder := xml.NewDecoder(wp.FileHandler)
	for wp.hasNext() {

		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == PAGE_ELEMENT {
				var p WikiPage
				decoder.DecodeElement(&p, &se)
				wp.Pages <- &p
			}
		default:
			continue
		}
	}
	close(wp.Pages)
}

func (wp *WikiParser) hasNext() bool {
	wp.ReadMutex.Lock()
	defer wp.ReadMutex.Unlock()
	// Check if reached limit
	fmt.Printf("\r%d/%d", wp.TotalParsed, wp.NumToParse)
	if  wp.TotalParsed >= wp.NumToParse {
		return false
	}
	wp.TotalParsed++
	return true
}

func (wp* WikiParser) ParseWorker(id int) {
	defer wp.ReadWaitGroup.Done()
	for page := range wp.Pages {

		// wp.ReadMutex.Lock()
		// wp.TotalParsed++
		_ = page

		// if wp.TotalParsed % 1000 == 0 {
		// fmt.Printf("\r%d\t%d", id, wp.TotalParsed)
		// }
		// // Need to offset the number of goroutines which have already
		// // grabbed a page off the channel
		// if wp.TotalParsed > wp.NumToParse - wp.NumOfReaders {
		// 	wp.ReadMutex.Unlock()
		// 	break
		// } 

		// wp.ReadMutex.Unlock()
	}	
}
