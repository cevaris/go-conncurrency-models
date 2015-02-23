package word_count

import (
	"strings"
)


type Words struct {
	Text string
}

func NewWords(text string) *Words {
	return &Words{Text: text}
}

func (w *Words) Iterator() <-chan string {
	ch := make(chan string)
	go func() {
		words := strings.Fields(w.Text)
		for _, val := range words {
			ch <- strings.TrimSpace(val)
		}
		close(ch)
	}()
	return ch
}
