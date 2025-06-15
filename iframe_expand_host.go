package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"io"
)

var (
	//go:embed "script/iframe_expand_receive.js"
	scriptIframeExpandReceive []byte
)

type IframeExpandElement struct {
	BaseElement
	r      Registrar
	iframe Element
}

func (ife *IframeExpandElement) Write(w io.Writer) error {
	return ife.iframe.Write(w)
}

// IframeExpandHost creates iframe-expand that will expand height to content height.
// In order to achieve that, two scripts need to be present
// script/receive.js on the host page (the page that contains iframe element)
// script/post.js within the iframe page. See NewContent that creates the page
// with that script. Initially host iframe has opacity: 0 through `loading`
// class to avoid flash of white content as iframe loads
func IframeExpandHost(r Registrar, id, src string) Element {
	iframe := IframeLazy(src)
	iframe.SetId(id)

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(compton_atoms.IframeExpandHost))
	r.RegisterRequirements(compton_atoms.ScriptName(compton_atoms.IframeExpandHost),
		Script(scriptIframeExpandReceive))

	iframe.SetAttribute("style", "view-transition-name:iframe-content-"+id)

	return &IframeExpandElement{
		r:      r,
		iframe: iframe,
	}
}
