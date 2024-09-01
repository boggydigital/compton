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

func NewForm(action, method string) compton.Element {
	form := compton.NewElement(atom.Form, markupForm)
	form.SetAttr(compton.ActionAttr, action)
	form.SetAttr(compton.MethodAttr, method)
	return form
}
