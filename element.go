package compton

import (
	"golang.org/x/net/html/atom"
	"io"
)

type Element interface {
	GetTagName() atom.Atom
	Append(children ...Element)
	Write(w io.Writer) error
	SetId(id string)
	SetClass(classes ...string)
	SetAttr(name, val string)
	GetAttr(name string) string
	GetElementById(id string) Element
	GetElementsByTagName(tagName atom.Atom) []Element
}
