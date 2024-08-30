package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/script.html"
	markupScript []byte
)

func NewScript(code []byte) compton.Element {
	script := compton.NewElement(atom.Script, markupScript)
	script.Append(NewText(string(code)))
	return script
}
