package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/form.html"
	markupForm []byte
)

func Form(action, method string) compton.Element {
	form := compton.NewElement(atom.Form, markupForm)
	form.SetAttribute(compton.ActionAttr, action)
	form.SetAttribute(compton.MethodAttr, method)
	return form
}
