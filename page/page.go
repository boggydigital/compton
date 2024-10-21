package page

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/units"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/script"
	"golang.org/x/net/html/atom"
	"io"
)

var (
	//go:embed "markup/page.html"
	markupPage []byte
)

var (
	//go:embed "style/page.css"
	stylePage []byte
)

type PageElement struct {
	compton.BaseElement
	registry map[string]any
	document compton.Element
}

func (p *PageElement) Write(w io.Writer) error {

	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		if styleClasses := head.GetElementById("style-classes"); styleClasses == nil {
			p.AppendStyle("style-classes", class.StyleClasses())
		}
	}

	return p.document.Write(w)
}

func (p *PageElement) Append(children ...compton.Element) {
	if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
		if content := body.GetFirstElementByTagName(compton_atoms.Content); content != nil {
			content.Append(children...)
		}
	}
}

func (p *PageElement) RegisterStyle(name string, style []byte) {
	if _, ok := p.registry[name]; !ok {
		p.registry[name] = nil
		p.AppendStyle(name, style)
	}
}

func (p *PageElement) RegisterRequirement(name string, element compton.Element) {
	if _, ok := p.registry[name]; !ok {
		p.registry[name] = nil
		if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
			if req := body.GetFirstElementByTagName(compton_atoms.Requirements); req != nil {
				req.Append(element)
			}
		}
	}
}

func (p *PageElement) RegisterDeferral(name string, element compton.Element) {
	if _, ok := p.registry[name]; !ok {
		p.registry[name] = nil
		if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
			if def := body.GetFirstElementByTagName(compton_atoms.Deferrals); def != nil {
				def.Append(element)
			}
		}
	}
}

func (p *PageElement) IsRegistered(name string) bool {
	_, ok := p.registry[name]
	return ok
}

func (p *PageElement) AppendStyle(id string, style []byte) *PageElement {
	if len(style) > 0 {
		if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
			head.Append(els.Style(style, id))
		}
	}
	return p
}

func (p *PageElement) AppendManifest() *PageElement {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(els.Link(map[string]string{attr.Rel: attr.Manifest, attr.Href: attr.ManifestJson}))
		head.Append(els.Meta(map[string]string{attr.Name: attr.MobileWebAppCapable, attr.Content: attr.Yes}))
		head.Append(els.Meta(map[string]string{attr.Name: attr.AppleMobileWebAppStatusBarStyle, attr.Content: attr.BlackTranslucent}))
	}
	return p
}

func (p *PageElement) AppendIcon() *PageElement {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(els.Link(map[string]string{attr.Rel: attr.Icon, attr.Href: attr.IconPng, attr.Type: attr.ImagePng}))
	}
	return p
}

func (p *PageElement) SetBodyId(id string) *PageElement {
	if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
		body.SetId(id)
	}
	return p
}

func (p *PageElement) ScriptDigests() []string {

	if scripts := p.GetElementsByTagName(atom.Script); len(scripts) > 0 {
		digests := make([]string, 0, len(scripts))
		for _, s := range scripts {
			if se, ok := s.(*script.ScriptElement); ok {
				digests = append(digests, se.Sha256())
			}
		}
		return digests
	}
	return nil
}

func (p *PageElement) appendMetaCharset() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(els.Meta(map[string]string{attr.Charset: attr.Utf8}))
	}
}

func (p *PageElement) appendTitle(title string) {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(els.Title(title))
	}
}

func (p *PageElement) appendViewport() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(els.Meta(map[string]string{attr.Viewport: attr.ViewportDefaults}))
	}
}

func (p *PageElement) appendColorScheme() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(els.Meta(map[string]string{attr.ColorScheme: attr.DarkLight}))
	}
}

func Page(title string) *PageElement {
	page := &PageElement{
		BaseElement: compton.BaseElement{
			Markup:  markupPage,
			TagName: atom.Body,
		},
		registry: make(map[string]any),
	}

	page.document = els.Document()
	html := els.Html("en")
	page.document.Append(els.Doctype(), html)

	head := els.Head()
	body := els.Body()
	html.Append(head, body)

	body.Append(els.Requirements(), els.Content(), els.Deferrals())

	page.appendMetaCharset()
	page.appendTitle(title)
	page.appendViewport()
	page.appendColorScheme()

	page.AppendStyle("style-colors", color.StyleSheet)
	page.AppendStyle("style-units", units.StyleSheet)
	page.AppendStyle("style-page", stylePage)

	return page
}
