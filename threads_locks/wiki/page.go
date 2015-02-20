package wiki

type Page interface {
	GetTitle() string
	GetText() string
	IsPoisonPill() bool
}
