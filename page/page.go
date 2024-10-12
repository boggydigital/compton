package page

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/units"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/net/html/atom"
	"io"
	"strconv"
)

var (
	//go:embed "markup/page.html"
	markupPage []byte
)

var (
	//go:embed "style/page.css"
	stylePages []byte
)

type PageElement struct {
	compton.BaseElement
	customElementsRegistry map[string]any
	title                  string
	//favIconEmoji           string
	appStyles [][]byte
	favIcon   bool
	manifest  bool
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
	//case ".FavIconEmoji":
	//	if _, err := io.WriteString(w, p.favIconEmoji); err != nil {
	//		return err
	//	}
	case ".FavIcon":
		if p.favIcon {
			favIconLink := els.Link()
			favIconLink.SetAttribute("rel", "icon")
			favIconLink.SetAttribute("type", "image/png")
			favIconLink.SetAttribute("href", "icon.png")
			if err := favIconLink.WriteContent(w); err != nil {
				return err
			}
		}
	case ".Manifest":
		if p.manifest {
			manifestLink := els.Link()
			manifestLink.SetAttribute("rel", "manifest")
			manifestLink.SetAttribute("href", "manifest.json")
			if err := manifestLink.WriteContent(w); err != nil {
				return err
			}
		}
	case ".StyleColors":
		if _, err := w.Write(color.StyleSheet); err != nil {
			return err
		}
	case ".StyleUnits":
		if _, err := w.Write(units.StyleSheet); err != nil {
			return err
		}
	case ".StylePage":
		if _, err := w.Write(stylePages); err != nil {
			return err
		}
	case ".StyleClasses":
		if _, err := w.Write(class.StyleClasses()); err != nil {
			return err
		}
	case ".StyleApp":
		for ii, appStyle := range p.appStyles {
			id := "style-app-" + strconv.Itoa(ii)
			style := els.Style(appStyle, id)
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

func (p *PageElement) AppendStyle(styles ...[]byte) *PageElement {
	for _, style := range styles {
		if len(style) > 0 {
			p.appStyles = append(p.appStyles, style)
		}
	}
	return p
}

func (p *PageElement) AppendManifest() *PageElement {
	p.manifest = true
	return p
}

func (p *PageElement) AppendFavIcon() *PageElement {
	p.favIcon = true
	return p
}

//func (p *PageElement) SetFavIconEmoji(favIconEmoji string) *PageElement {
//	p.favIconEmoji = favIconEmoji
//	return p
//}

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
