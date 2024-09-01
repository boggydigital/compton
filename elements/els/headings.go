package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/h1.html"
	markupH1 []byte
	//go:embed "markup/h2.html"
	markupH2 []byte
	//go:embed "markup/h3.html"
	markupH3 []byte
	//go:embed "markup/h4.html"
	markupH4 []byte
	//go:embed "markup/h5.html"
	markupH5 []byte
	//go:embed "markup/h6.html"
	markupH6 []byte
)

func NewHeading(level int) compton.Element {

	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	tn := atom.H1
	markup := markupH1
	switch level {
	case 1:
		tn = atom.H1
		markup = markupH1
	case 2:
		tn = atom.H2
		markup = markupH2
	case 3:
		tn = atom.H3
		markup = markupH3
	case 4:
		tn = atom.H4
		markup = markupH4
	case 5:
		tn = atom.H5
		markup = markupH5
	case 6:
		tn = atom.H6
		markup = markupH6
	}
	return compton.NewElement(tn, markup)
}

func NewHeadingText(txt string, level int) compton.Element {
	heading := NewHeading(level)
	heading.Append(NewText(txt))
	return heading
}
