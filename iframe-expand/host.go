package iframe_expand

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/els"
	"io"
)

const elementName = "iframe-expand"

var (
	//go:embed "script/receive.js"
	markupReceiveScript []byte
)

type IframeExpand struct {
	compton.BaseElement
	r             compton.Registrar
	iframe        compton.Element
	receiveScript compton.Element
}

func (ife *IframeExpand) WriteRequirements(w io.Writer) error {
	if ife.r.RequiresRegistration(elementName) {
		return ife.receiveScript.WriteContent(w)
	}
	return nil
}

func (ife *IframeExpand) WriteContent(w io.Writer) error {
	return ife.iframe.WriteContent(w)
}

// New creates iframe-expand that will expand height to content height.
// In order to achieve that, two scripts need to be present
// script/receive.js on the host page (the page that contains iframe element)
// script/post.js within the iframe page. See NewContent that creates the page
// with that script. Initially host iframe has opacity: 0 through `loading`
// class to avoid flash of white content as iframe loads
func New(r compton.Registrar, id, src string) compton.Element {
	iframe := els.NewIframeLazy(src)
	iframe.SetId(id)
	iframe.SetClass("loading")
	receiveScript := els.NewScript(markupReceiveScript)
	return &IframeExpand{
		r:             r,
		iframe:        iframe,
		receiveScript: receiveScript,
	}
}
