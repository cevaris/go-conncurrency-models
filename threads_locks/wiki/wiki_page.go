package wiki


type WikiPage struct {
	Title string `xml:"title"`
	Text string `xml:"revision>text"`
}

func NewWikiPage(title string, text string) *WikiPage {
	return &WikiPage{Title: title, Text: text}
}

func (p *WikiPage) IsPoisonPill() bool {
	return false
}

func (p* WikiPage) GetTitle() string {
	return p.Title
}

func (p* WikiPage) GetText() string {
	return p.Text
}
