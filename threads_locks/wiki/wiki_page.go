package wiki

import (
	_ "github.com/cevaris/go_concurrency_models/threads_and_locks/parser"
)

type WikiPage struct {
	Title string
	Text string
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
