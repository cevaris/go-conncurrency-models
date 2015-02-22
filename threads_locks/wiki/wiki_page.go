package wiki


type WikiPage struct {
	Title string `xml:"title"`
	Text string `xml:"revision>text"`
}

func NewWikiPage(title string, text string) *WikiPage {
	return &WikiPage{Title: title, Text: text}
}
