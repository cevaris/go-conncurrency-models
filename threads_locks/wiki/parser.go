package wiki



import (
	"encoding/xml"
	"fmt"
	"os"
	"sync"
)

var PAGE_ELEMENT string = "page"

type Parser interface {
	Parse(file *os.File)
	ParseWorker(id int64)
}

type ParserConfig struct {
	NumOfReaders int64
	NumToRead int64
	TotalRead int64	
}

func NewParserConfig(numOfReaders int64, numToRead int64) (*ParserConfig) {
	return &ParserConfig{
		NumOfReaders: numOfReaders,
		NumToRead: numToRead,		
	}
}

type WikiParser struct {
	FileHandler *os.File
	Pages chan Page
	Config *ParserConfig
	ReadMutex *sync.Mutex
	ReadWaitGroup *sync.WaitGroup
}

func NewWikiParser(numOfPages int64, file *os.File) *WikiParser {
	config := NewParserConfig(1, numOfPages)
	return &WikiParser{
		FileHandler: file,
		Pages: make(chan Page),
		ReadMutex: &sync.Mutex{},
		ReadWaitGroup: &sync.WaitGroup{},
		Config: config,
	}
}

func (wp* WikiParser) SetConfig(config *ParserConfig) {
	wp.Config = config
	// Update to buffered channel to 5x number of readers
	wp.Pages = make(chan Page, config.NumOfReaders*5)
	wp.ReadWaitGroup.Add(int(config.NumOfReaders))
}

func (wp *WikiParser) Parse() {
	decoder := xml.NewDecoder(wp.FileHandler)
	
	go func() {
		for {
			t, _ := decoder.Token()
			if t == nil {
				break
			}
			// Inspect the type of the token just read.
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
	}()
}

func (wp* WikiParser) ParseWorker(id int) {
	
	config := wp.Config
	defer wp.ReadWaitGroup.Done()
	for page := range wp.Pages {

		wp.ReadMutex.Lock()
		config.TotalRead++
		_ = page

		if config.TotalRead % 1000 == 0 {
			fmt.Printf("\r%d\t%d", id, config.TotalRead)
		}
		// Need to offset the number of goroutines which have already
		// grabbed a page off the channel
		if config.TotalRead > config.NumToRead - config.NumOfReaders {
			wp.ReadMutex.Unlock()
			break
		} 

		wp.ReadMutex.Unlock()
	}	
}
