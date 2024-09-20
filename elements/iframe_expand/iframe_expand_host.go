package iframe_expand

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "iframe-expand"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "script/receive.js"
	scriptReceive []byte
	//go:embed "style/iframe-host.css"
	styleIframeHost []byte
)

type IframeExpandElement struct {
	compton.BaseElement
	r      compton.Registrar
	iframe compton.Element
}

func (ife *IframeExpandElement) WriteStyles(w io.Writer) error {
	if ife.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleIframeHost, styleRegistrationName).WriteContent(w); err != nil {
			return err
		}
	}
	return nil
}

func (ife *IframeExpandElement) WriteRequirements(w io.Writer) error {
	if ife.r.RequiresRegistration(registrationName) {
		receiveScript := els.ScriptAsync(scriptReceive)
		return receiveScript.WriteContent(w)
	}
	return nil
}

func (ife *IframeExpandElement) WriteContent(w io.Writer) error {
	return ife.iframe.WriteContent(w)
}

// IframeExpandHost creates iframe-expand that will expand height to content height.
// In order to achieve that, two scripts need to be present
// script/receive.js on the host page (the page that contains iframe element)
// script/post.js within the iframe page. See NewContent that creates the page
// with that script. Initially host iframe has opacity: 0 through `loading`
// class to avoid flash of white content as iframe loads
func IframeExpandHost(r compton.Registrar, id, src string) compton.Element {
	iframe := els.IframeLazy(src)
	iframe.SetId(id)
	iframe.AddClass("loading")
	return &IframeExpandElement{
		r:      r,
		iframe: iframe,
	}
}
