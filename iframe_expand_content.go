package compton

import (
	_ "embed"
)

var (
	//go:embed "script/iframe_expand_post.js"
	scriptIframeExpandPost []byte
	//go:embed "style/iframe-expand-content.css"
	styleIframeContent []byte
)

// IframeExpandContent creates an iframe content page and attaches
// script/post.js that send the message on iframe content size change
// to the host page that contains script/receive.js to size host
// iframe element and remove `loading` class
func IframeExpandContent(id, title string) PageElement {
	p := Page(title).
		AppendStyle("style-iframe-content", styleIframeContent).
		SetBodyId(id)
	p.Append(ScriptAsync(scriptIframeExpandPost))
	return p
}
