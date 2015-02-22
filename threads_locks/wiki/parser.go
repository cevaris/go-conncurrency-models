package wiki

// http://blog.davidsingleton.org/parsing-huge-xml-files-with-go

import (
	"os"
	"encoding/xml"
	// "fmt"
)

var PAGE_ELEMENT string = "page"

func Parse(file *os.File) {
	decoder := xml.NewDecoder(file)
	total := 0
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
				total++
				// fmt.Println(total, p.Title)
			}
		default:
			continue
		}
	}
}
