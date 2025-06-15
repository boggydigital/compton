package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

var (
	//go:embed "script/iframe_expand_post.js"
	scriptIframeExpandPost []byte
)

// IframeExpandContent creates an iframe content page and attaches
// script/post.js that send the message on iframe content size change
// to the host page that contains script/receive.js to size host
// iframe element and remove `loading` class
func IframeExpandContent(id, title string) PageElement {
	p := Page(title).
		SetBodyId(id)

	p.SetAttribute("style", "view-transition-name:iframe-content-"+id)

	p.RegisterStyles(DefaultStyle, compton_atoms.StyleName(compton_atoms.IframeExpandContent))
	p.RegisterDeferrals(compton_atoms.ScriptName(compton_atoms.IframeExpandContent),
		ScriptAsync(scriptIframeExpandPost))
	return p
}
