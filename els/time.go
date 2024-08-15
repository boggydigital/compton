package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/time.html"
	markupTime []byte
)

func NewTime() compton.Element {
	return compton.NewElement(atom.Time, markupTime)
}

func NewTimeText(txt string) compton.Element {
	tm := NewTime()
	tm.Append(NewText(txt))
	return tm
}
