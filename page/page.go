package page

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"io"
)

var (
	//go:embed "markup/page.html"
	markupPage []byte
	//go:embed "markup/define_script.html"
	markupDefineScript []byte
)

var (
	//go:embed "style/colors.css"
	styleColors []byte
	//go:embed "style/units.css"
	styleUnits []byte
	//go:embed "style/page.css"
	stylePages []byte
)

type Page struct {
	compton.Parent
	registry     map[string]any
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
	return compton.WriteContents(bytes.NewReader(markupPage), w, p.writeFragment)
}

func (p *Page) writeFragment(t string, w io.Writer) error {
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
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func (p *Page) Register(name string, template []byte, mode compton.EncapsulationMode, w io.Writer) error {

	if _, ok := p.registry[name]; ok {
		return nil
	}

	if err := compton.WriteContents(
		bytes.NewReader(markupDefineScript), w,
		func(token string, w io.Writer) error {
			switch token {
			case ".ElementName":
				if _, err := io.WriteString(w, name); err != nil {
					return err
				}
			case ".Mode":
				if _, err := io.WriteString(w, string(mode)); err != nil {
					return err
				}
			default:
				return compton.ErrUnknownToken(token)
			}
			return nil
		}); err != nil {
		return err
	}

	if _, err := w.Write(template); err != nil {
		return err
	}

	p.registry[name] = nil

	return nil
}

func New(title string) *Page {
	return &Page{
		registry: make(map[string]any),
		Title:    title,
	}
}
