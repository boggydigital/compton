package page

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/net/html/atom"
	"io"
)

var (
	//go:embed "markup/page.html"
	markupPage []byte
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

type PageElement struct {
	compton.BaseElement
	customElementsRegistry map[string]any
	title                  string
	favIconEmoji           string
	customStyles           [][]byte
}

func (p *PageElement) WriteContent(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupPage), w, p.writeFragment)
}

func (p *PageElement) writeFragment(t string, w io.Writer) error {
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
	case ".StyleClasses":
		if _, err := w.Write(class.StyleClasses()); err != nil {
			return err
		}
	case ".StyleApp":
		for _, customStyle := range p.customStyles {
			style := els.Style(customStyle)
			if err := style.WriteContent(w); err != nil {
				return err
			}
		}
	case ".ElementsStyles":
		if err := p.WriteStyles(w); err != nil {
			return err
		}
	case compton.RequirementsToken:
		if err := p.WriteRequirements(w); err != nil {
			return err
		}
	case compton.DeferralsToken:
		if err := p.WriteDeferrals(w); err != nil {
			return err
		}
	case compton.AttributesToken:
		fallthrough
	case compton.ContentToken:
		if err := p.WriteFragment(t, w); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func (p *PageElement) RequiresRegistration(name string) bool {
	if _, ok := p.customElementsRegistry[name]; !ok {
		p.customElementsRegistry[name] = nil
		return true
	}
	return false
}

func (p *PageElement) SetCustomStyles(customStyles ...[]byte) *PageElement {
	p.customStyles = customStyles
	return p
}

func (p *PageElement) SetFavIconEmoji(favIconEmoji string) *PageElement {
	p.favIconEmoji = favIconEmoji
	return p
}

func Page(title string) *PageElement {
	return &PageElement{
		BaseElement: compton.BaseElement{
			Markup:  markupPage,
			TagName: atom.Body,
		},
		title:                  title,
		customElementsRegistry: make(map[string]any),
	}
}
