package iframe_expand

import (
	_ "embed"
	"github.com/boggydigital/compton/elements/script"
	"github.com/boggydigital/compton/page"
)

var (
	//go:embed "script/post.js"
	scriptPost []byte
	//go:embed "style/iframe-content.css"
	styleIframeContent []byte
)

// IframeExpandContent creates an iframe content page and attaches
// script/post.js that send the message on iframe content size change
// to the host page that contains script/receive.js to size host
// iframe element and remove `loading` class
func IframeExpandContent(id, title string) *page.PageElement {
	p := page.Page(title).AppendStyle("style-iframe-content", styleIframeContent)
	p.SetBodyId(id)
	p.Append(script.ScriptAsync(scriptPost))
	return p
}
