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

func Script(code []byte) compton.Element {
	script := compton.NewElement(atom.Script, markupScript)
	script.Append(Text(string(code)))
	return script
}

func ScriptAsync(code []byte) compton.Element {
	script := Script(code)
	script.SetAttribute("async", "")
	return script
}
