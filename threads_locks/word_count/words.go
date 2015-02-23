package word_count

import (
	"fmt"
	"strings"
	"regexp"
)


type Words struct {
	Text string
}

func NewWords(text string) *Words {
	return &Words{Text: text}
}

func Sanitize(text string) string {
	// Capture AlphaNumeric, -, _
	reg, err := regexp.Compile("[^\\w-_]+")
	if err != nil {
		fmt.Println(err, text)
		return ""
	}
	sanitized := reg.ReplaceAllString(strings.ToLower(text), "")
	if len(sanitized) > 15 {
		// fmt.Println("Dropping", text)
		return ""
	}
	return sanitized
}

func (w *Words) Iterator() <-chan string {
	ch := make(chan string)
	go func() {
		words := strings.Fields(w.Text)
		for _, val := range words {
			ch <- Sanitize(strings.TrimSpace(val))
		}
		close(ch)
	}()
	return ch
}
