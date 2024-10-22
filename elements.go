package compton

import (
	"bytes"
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/loading"
	"golang.org/x/net/html/atom"
	"io"
)

var (
	//go:embed "markup/*.html"
	markup embed.FS
	//go:embed "style/*.css"
	style embed.FS
)

/* https://developer.mozilla.org/en-US/docs/Web/API/Text */

type TextElement struct {
	BaseElement
	content string
}

func (t *TextElement) Append(_ ...Element) {
}

func (t *TextElement) Write(w io.Writer) error {
	if _, err := io.WriteString(w, t.content); err != nil {
		return err
	}
	return nil
}

func Text(content string) Element {
	return &TextElement{
		content: content,
	}
}

func TextBytes(content []byte) Element {
	return &TextElement{
		content: string(content),
	}
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a */

func A(href string) Element {
	anchor := newElementAtom(atom.A, markup)
	anchor.SetAttribute(attr.Href, href)
	return anchor
}

func AText(txt, href string) Element {
	anchor := A(href)
	anchor.Append(Text(txt))
	return anchor
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br */

var Br = Break

func Break() Element {
	return newElementAtom(atom.Br, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body */

func Body() Element {
	return newElementAtom(atom.Body, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist */

func Datalist(id string) Element {
	dataList := newElementAtom(atom.Datalist, markup)
	dataList.SetAttribute(attr.Id, id)
	return dataList
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details */

type DetailsElement struct {
	BaseElement
}

func (d *DetailsElement) AppendSummary(children ...Element) *DetailsElement {
	var summary Element
	if summaries := d.GetElementsByTagName(atom.Summary); len(summaries) > 0 {
		summary = summaries[0]
	}

	if summary == nil {
		summary = Summary()
		d.Append(summary)
	}

	for _, child := range children {
		summary.Append(child)
	}

	return d
}

func (d *DetailsElement) Open() *DetailsElement {
	d.SetAttribute("open", "")
	return d
}

func Details() *DetailsElement {
	return &DetailsElement{
		BaseElement{
			TagName:  atom.Details,
			Markup:   markup,
			Filename: atomMarkupFilename(atom.Details),
		},
	}
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div */

func Div() Element {
	return newElementAtom(atom.Div, markup)
}

func DivText(txt string) Element {
	div := Div()
	div.Append(Text(txt))
	return div
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form */

func Form(action, method string) Element {
	form := newElementAtom(atom.Form, markup)
	form.SetAttribute(attr.Action, action)
	form.SetAttribute(attr.Method, method)
	return form
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head */

func Head() Element {
	return newElementAtom(atom.Head, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements */

func Heading(level int) Element {

	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	tn := atom.H1
	switch level {
	case 1:
		tn = atom.H1
	case 2:
		tn = atom.H2
	case 3:
		tn = atom.H3
	case 4:
		tn = atom.H4
	case 5:
		tn = atom.H5
	case 6:
		tn = atom.H6
	}
	return newElementAtom(tn, markup)
}

func H1() Element { return Heading(1) }
func H2() Element { return Heading(2) }
func H3() Element { return Heading(3) }
func H4() Element { return Heading(4) }
func H5() Element { return Heading(5) }
func H6() Element { return Heading(6) }

func HeadingText(txt string, level int) Element {
	heading := Heading(level)
	heading.Append(Text(txt))
	return heading
}

func H1Text(txt string) Element { return HeadingText(txt, 1) }
func H2Text(txt string) Element { return HeadingText(txt, 2) }
func H3Text(txt string) Element { return HeadingText(txt, 3) }
func H4Text(txt string) Element { return HeadingText(txt, 4) }
func H5Text(txt string) Element { return HeadingText(txt, 5) }
func H6Text(txt string) Element { return HeadingText(txt, 6) }

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr */

var Hr = HorizontalRule

func HorizontalRule() Element {
	return newElementAtom(atom.Hr, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html */

func Html(lang string) Element {
	html := newElementAtom(atom.Html, markup)
	if lang != "" {
		html.SetAttribute("lang", lang)
	}
	return html
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe */

func Iframe(src string) Element {
	iframe := newElementAtom(atom.Iframe, markup)
	iframe.SetAttribute(attr.Src, src)
	return iframe
}

func IframeLazy(src string) Element {
	iframe := Iframe(src)
	iframe.SetAttribute(attr.Loading, loading.Lazy.String())
	return iframe
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img */

var (
	Img      = Image
	ImgLazy  = ImageLazy
	ImgEager = ImageEager
)

func Image(src string) Element {
	image := newElementAtom(atom.Img, markup)
	if src != "" {
		image.SetAttribute(attr.Src, src)
	}
	return image
}

func ImageLazy(src string) Element {
	image := Image(src)
	image.SetAttribute(attr.Loading, loading.Lazy.String())
	return image
}

func ImageEager(src string) Element {
	image := Image(src)
	image.SetAttribute(attr.Loading, loading.Eager.String())
	return image
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label */

func Label(forInput string) Element {
	label := newElementAtom(atom.Label, markup)
	if forInput != "" {
		label.SetAttribute(attr.For, forInput)
	}
	return label
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li */

var Li = ListItem

func ListItem() Element {
	return newElementAtom(atom.Li, markup)
}

func ListItemText(txt string) Element {
	listItem := ListItem()
	listItem.Append(Text(txt))
	return listItem
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link */

func Link(kv map[string]string) Element {
	link := newElementAtom(atom.Link, markup)
	for k, v := range kv {
		link.SetAttribute(k, v)
	}
	return link
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta */

func Meta(kv map[string]string) Element {
	meta := newElementAtom(atom.Meta, markup)
	for k, v := range kv {
		meta.SetAttribute(k, v)
	}
	return meta
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol */

var Ol = OrderedList

func OrderedList() Element {
	return newElementAtom(atom.Ol, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option */

func Option(value, label string) Element {
	option := newElementAtom(atom.Option, markup)
	option.SetAttribute(attr.Value, value)
	if label != "" {
		option.SetAttribute(attr.Label, label)
	}
	return option
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script */

var b64 = base64.StdEncoding

type ScriptElement struct {
	BaseElement
	hash []byte
}

func computeSha256(reader io.Reader) ([]byte, error) {
	h := sha256.New()
	var err error
	if _, err = io.Copy(h, reader); err == nil {
		return h.Sum(nil), nil
	}
	return nil, err
}

func (se *ScriptElement) Sha256() string {
	if len(se.hash) > 0 {
		return "sha256-" + b64.EncodeToString(se.hash)
	}
	return ""
}

func Script(code []byte) *ScriptElement {
	script := &ScriptElement{
		BaseElement: BaseElement{
			TagName:  atom.Script,
			Markup:   markup,
			Filename: atomMarkupFilename(atom.Script),
		},
	}

	if hash, err := computeSha256(bytes.NewReader(code)); err == nil {
		script.hash = hash
	}

	script.Append(Text(string(code)))
	return script
}

func ScriptAsync(code []byte) *ScriptElement {
	script := Script(code)
	script.SetAttribute("async", "")
	return script
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section */

func Section() Element {
	return newElementAtom(atom.Section, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span */

func Span() Element {
	return newElementAtom(atom.Span, markup)
}

func SpanText(txt string) Element {
	span := Span()
	span.Append(Text(txt))
	return span
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style */

func Style(styles []byte, id string) Element {
	style := newElementAtom(atom.Style, markup)
	style.Append(Text(string(styles)))
	if id != "" {
		style.SetId(id)
	}
	return style
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary */

func Summary() Element {
	return newElementAtom(atom.Summary, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody */

func Tbody() Element {
	return newElementAtom(atom.Tbody, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td */

func Td() Element {
	return newElementAtom(atom.Td, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot */

func Tfoot() Element {
	return newElementAtom(atom.Tfoot, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th */

func Th() Element {
	return newElementAtom(atom.Th, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead */

func Thead() Element {
	return newElementAtom(atom.Thead, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time */

func Time() Element {
	return newElementAtom(atom.Time, markup)
}

func TimeText(txt string) Element {
	tm := Time()
	tm.Append(Text(txt))
	return tm
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title */

func Title(txt string) Element {
	title := newElementAtom(atom.Title, markup)
	title.Append(Text(txt))
	return title
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr */

func Tr() Element {
	return newElementAtom(atom.Tr, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul */

var Ul = UnorderedList

func UnorderedList() Element {
	return newElementAtom(atom.Ul, markup)
}

/* https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video */

func Video(src string) Element {
	video := newElementAtom(atom.Video, markup)
	if src != "" {
		video.SetAttribute(attr.Src, src)
	}
	return video
}

/* required by compton.Page */

func Content() Element {
	return NewElement(compton_atoms.Content, markup, compton_atoms.MarkupName(compton_atoms.Content))
}

func Deferrals() Element {
	return NewElement(compton_atoms.Deferrals, markup, compton_atoms.MarkupName(compton_atoms.Deferrals))
}

func Doctype() Element {
	return NewElement(compton_atoms.Doctype, markup, compton_atoms.MarkupName(compton_atoms.Doctype))
}

func Document() Element {
	return NewElement(compton_atoms.Document, markup, compton_atoms.MarkupName(compton_atoms.Document))
}

func Requirements() Element {
	return NewElement(compton_atoms.Requirements, markup, compton_atoms.MarkupName(compton_atoms.Requirements))
}
