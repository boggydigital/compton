package iframe_expand

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/page"
)

var (
	//go:embed "script/post.js"
	postScript []byte
)

// IframeExpandContent creates an iframe content page and attaches
// script/post.js that send the message on iframe content size change
// to the host page that contains script/receive.js to size host
// iframe element and remove `loading` class
func IframeExpandContent(id, title string) compton.Element {
	p := page.Page(title)
	p.SetId(id)
	p.Append(els.Script(postScript))
	return p
}
