package wiki



import (
	"os"
	"encoding/xml"
)

var PAGE_ELEMENT string = "page"

type Parser interface {
	Parse(file *os.File)
}

type WikiParser struct {
	Pages chan Page
}

func NewWikiParser() *WikiParser {
	return &WikiParser{
		Pages: make(chan Page, 100),
	}
}

func (wp *WikiParser) Parse(file *os.File) <-chan Page {
	decoder := xml.NewDecoder(file)
	
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
	return wp.Pages
}
