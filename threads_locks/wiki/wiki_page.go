package wiki

import (
	"github.com/cevaris/go_concurrency_models/threads_and_locks/parser"
)

type WikiPage struct {
	Title string `xml:title`
	Text string `xml:text`
}

func NewWikiPage(title string, text string) *WikiPage {
	return &WikiPage{Title: title, Text: text}
}

func (p* WikiPage) GetTitle() string {
	return p.Title
}

func (p* WikiPage) GetText() string {
	return p.Text
}
