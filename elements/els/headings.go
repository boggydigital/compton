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

func Heading(level int) compton.Element {

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

func H1() compton.Element { return Heading(1) }
func H2() compton.Element { return Heading(2) }
func H3() compton.Element { return Heading(3) }
func H4() compton.Element { return Heading(4) }
func H5() compton.Element { return Heading(5) }
func H6() compton.Element { return Heading(6) }

func HeadingText(txt string, level int) compton.Element {
	heading := Heading(level)
	heading.Append(Text(txt))
	return heading
}

func H1Text(txt string) compton.Element { return HeadingText(txt, 1) }
func H2Text(txt string) compton.Element { return HeadingText(txt, 2) }
func H3Text(txt string) compton.Element { return HeadingText(txt, 3) }
func H4Text(txt string) compton.Element { return HeadingText(txt, 4) }
func H5Text(txt string) compton.Element { return HeadingText(txt, 5) }
func H6Text(txt string) compton.Element { return HeadingText(txt, 6) }
