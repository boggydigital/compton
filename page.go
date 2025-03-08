package compton

import (
	"embed"
	_ "embed"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/font_weight"
	"golang.org/x/net/html/atom"
	"io"
	"net/http"
	"strings"
	"sync"
)

type pageElement struct {
	BaseElement
	registry map[string]any
	document Element
	mux      *sync.Mutex
}

func (p *pageElement) appendStyleClasses() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		if styleClasses := head.GetElementById("style-classes"); styleClasses == nil {
			classesStyle := Style(class.StyleClasses())
			classesStyle.SetId("style-classes")
			head.Append(classesStyle)
		}
	}
}

func (p *pageElement) Append(children ...Element) {
	p.document.Append(children...)
}

func (p *pageElement) WriteResponse(w http.ResponseWriter) error {
	return p.writeResponse(w, "text/html; charset=utf-8")
}

func (p *pageElement) writeResponse(w http.ResponseWriter, contentType string) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	if policy := p.contentSecurityPolicy(); policy != "" {
		w.Header().Set("Content-Security-Policy", policy)
	}
	w.Header().Set("Content-Type", contentType)
	return p.Write(w)
}

func (p *pageElement) Write(w io.Writer) error {
	p.appendStyleClasses()
	return p.document.Write(w)
}

func (p *pageElement) RegisterStyles(efs embed.FS, names ...string) {
	for _, name := range names {
		if _, ok := p.registry[name]; !ok {
			p.registry[name] = nil
			if content, err := efs.ReadFile(name); err == nil {
				if len(content) > 0 {
					if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
						head.Append(Style(content))
					}
				}
			} else {
				panic(err)
			}
		}
	}
}

func (p *pageElement) RegisterRequirements(name string, elements ...Element) {
	if _, ok := p.registry[name]; !ok {
		p.registry[name] = nil
		if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
			if req := body.GetFirstElementByTagName(compton_atoms.Requirements); req != nil {
				req.Append(elements...)
			}
		}
	}
}

func (p *pageElement) RegisterDeferrals(name string, elements ...Element) {
	if _, ok := p.registry[name]; !ok {
		p.registry[name] = nil
		if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
			if def := body.GetFirstElementByTagName(compton_atoms.Deferrals); def != nil {
				def.Append(elements...)
			}
		}
	}
}

func (p *pageElement) Error(err error) PageElement {
	msg := Fspan(p, err.Error()).BackgroundColor(color.Red).FontWeight(font_weight.Bolder)
	msg.AddClass("_compton_error_message")
	p.RegisterRequirements(err.Error(), FICenter(p, msg))
	return p
}

func (p *pageElement) IsRegistered(name string) bool {
	_, ok := p.registry[name]
	return ok
}

func (p *pageElement) AppendManifest() PageElement {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(Link(map[string]string{
			attr.Rel: attr.Manifest, attr.Href: attr.ManifestJson,
		}))
		head.Append(Meta(map[string]string{
			attr.Name: attr.MobileWebAppCapable, attr.Content: attr.Yes,
		}))
		head.Append(Meta(map[string]string{
			attr.Name: attr.AppleMobileWebAppStatusBarStyle, attr.Content: attr.BlackTranslucent,
		}))
	}
	return p
}

func (p *pageElement) AppendIcon() PageElement {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(Link(map[string]string{
			attr.Rel:  attr.Icon,
			attr.Href: attr.IconPng,
			attr.Type: attr.ImagePng,
		}))
	}
	return p
}

func (p *pageElement) AppendSpeculationRules(hrefMatches ...string) {

	if srBytes := SpeculationRulesBytes(hrefMatches...); len(srBytes) > 0 {
		srScript := ScriptAsync(srBytes)
		srScript.SetAttribute(attr.Type, speculationRulesName)

		p.RegisterDeferrals(speculationRulesName, srScript)
	}
}

func (p *pageElement) SetBodyAttribute(name, val string) {
	if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
		body.SetAttribute(name, val)
	}
}

func (p *pageElement) SetAttribute(name, val string) {
	if html := p.document.GetFirstElementByTagName(atom.Html); html != nil {
		html.SetAttribute(name, val)
	}
}

func (p *pageElement) SetBodyId(id string) PageElement {
	if body := p.document.GetFirstElementByTagName(atom.Body); body != nil {
		body.SetId(id)
	}
	return p
}

func (p *pageElement) contentSecurityPolicy() string {
	if scripts := p.document.GetElementsByTagName(atom.Script); len(scripts) > 0 {
		digests := make([]string, 0, len(scripts))
		for _, s := range scripts {
			if se, ok := s.(*ScriptElement); ok {
				digests = append(digests, "'"+se.Sha256()+"'")
			}
		}
		return "script-src " + strings.Join(digests, " ")
	}
	return ""
}

func (p *pageElement) appendMetaCharset() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(Meta(map[string]string{attr.Charset: attr.Utf8}))
	}
}

func (p *pageElement) appendMetaFormatDetectionTelephoneNo() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(Meta(map[string]string{
			attr.Name:    attr.FormatDetection,
			attr.Content: attr.TelephoneNo,
		}))
	}
}

func (p *pageElement) appendTitle(title string) {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(Title(title))
	}
}

func (p *pageElement) appendViewport() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(Meta(map[string]string{
			attr.Name:    attr.Viewport,
			attr.Content: attr.ViewportDefaults,
		}))
	}
}

func (p *pageElement) appendColorScheme() {
	if head := p.document.GetFirstElementByTagName(atom.Head); head != nil {
		head.Append(Meta(map[string]string{
			attr.ColorScheme: attr.DarkLight,
		}))
	}
}

func Page(title string) PageElement {
	page := &pageElement{
		BaseElement: BaseElement{
			TagName: compton_atoms.Page,
		},
		registry: make(map[string]any),
		mux:      &sync.Mutex{},
	}

	page.document = Document()
	html := Html("en")
	page.document.Append(Doctype(), html)

	body := Body()
	html.Append(Head(), body)

	body.Append(Requirements(), Content(), Deferrals())

	page.appendMetaCharset()
	page.appendTitle(title)
	page.appendViewport()
	page.appendColorScheme()
	page.appendMetaFormatDetectionTelephoneNo()

	page.RegisterStyles(DefaultStyle,
		"style/colors.css", "style/units.css", "style/page.css")

	return page
}
