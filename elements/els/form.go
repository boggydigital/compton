package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/form.html"
	markupForm []byte
)

func Form(action, method string) compton.Element {
	form := compton.NewElement(atom.Form, markupForm)
	form.SetAttribute(attr.Action, action)
	form.SetAttribute(attr.Method, method)
	return form
}
