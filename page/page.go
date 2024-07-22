package page

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton"
	"io"
)

var (
	//go:embed "markup/page.html"
	markupPage []byte
)

var (
	//go:embed "styles/colors.css"
	styleColors []byte
	//go:embed "styles/units.css"
	styleUnits []byte
	//go:embed "styles/page.css"
	stylePages []byte
)

func ErrUnknownToken(t string) error {
	return fmt.Errorf("unknown token: %s", t)
}

type Page struct {
	compton.Parent
	Title        string
	CustomStyles []byte
}

func (p *Page) Add(children ...compton.Component) {
	p.Parent.Add(children...)
}

func (p *Page) AddCustomStyles(customStyles []byte) {
	p.CustomStyles = customStyles
}

func (p *Page) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupPage), w, p.writePageFragment)
}

func (p *Page) writePageFragment(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if _, err := io.WriteString(w, p.Title); err != nil {
			return err
		}
	case ".StyleColors":
		if _, err := w.Write(styleColors); err != nil {
			return err
		}
	case ".StyleUnits":
		if _, err := w.Write(styleUnits); err != nil {
			return err
		}
	case ".StylePage":
		if _, err := w.Write(stylePages); err != nil {
			return err
		}
	case ".StyleCustom":
		if len(p.CustomStyles) > 0 {
			if _, err := w.Write(p.CustomStyles); err != nil {
				return err
			}
		}
	case ".Body":
		for _, child := range p.Children {
			if err := child.Write(w); err != nil {
				return err
			}
		}
	default:
		return ErrUnknownToken(t)
	}
	return nil
}

func New(title string) compton.Component {
	return &Page{
		Title: title,
	}
}
