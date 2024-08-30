package iframe_expand

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/els"
	"github.com/boggydigital/compton/page"
)

var (
	//go:embed "script/post.js"
	postScript []byte
)

// NewContent creates an iframe content page and attaches
// script/post.js that send the message on iframe content size change
// to the host page that contains script/receive.js to size host
// iframe element and remove `loading` class
func NewContent(id, title string) compton.Element {
	p := page.New(title)
	p.SetId(id)
	p.Append(els.NewScript(postScript))
	return p
}
