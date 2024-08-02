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
	//go:embed "style/elements.css"
	styleElements []byte
)

type Page struct {
	compton.BaseElement
	registry     map[string]any
	title        string
	favIconEmoji string
	customStyles []byte
}

func (p *Page) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupPage), w, p.writeFragment)
}

func (p *Page) writeFragment(t string, w io.Writer) error {
	switch t {
	case ".Title":
		if _, err := io.WriteString(w, p.title); err != nil {
			return err
		}
	case ".FavIconEmoji":
		if _, err := io.WriteString(w, p.favIconEmoji); err != nil {
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
	case ".StyleElements":
		if _, err := w.Write(styleElements); err != nil {
			return err
		}
	case ".StyleCustom":
		if len(p.customStyles) > 0 {
			if _, err := w.Write(p.customStyles); err != nil {
				return err
			}
		}
	case compton.AttributesToken:
		fallthrough
	case compton.ContentToken:
		if err := p.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func (p *Page) RegisterCustomElement(name, extends string, mode compton.EncapsulationMode, w io.Writer) error {

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
			case ".ExtendsElement":
				if _, err := io.WriteString(w, extends); err != nil {
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

	p.registry[name] = nil

	return nil
}

func (p *Page) RegisterMarkup(r io.Reader, w io.Writer, tw compton.TokenWriter) error {
	return compton.WriteContents(r, w, tw)
}

func (p *Page) SetCustomStyles(customStyles []byte) {
	p.customStyles = customStyles
}

func New(title, favIconEmoji string) *Page {
	return &Page{
		BaseElement:  compton.BaseElement{Markup: markupPage},
		registry:     make(map[string]any),
		title:        title,
		favIconEmoji: favIconEmoji,
	}
}
