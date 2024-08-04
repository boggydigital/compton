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
	SetClass(names ...string)
	HasClass(names ...string) bool
	SetAttr(name, val string)
	GetAttr(name string) string
	GetElementById(id string) Element
	GetElementsByTagName(tagName atom.Atom) []Element
	GetElementsByClassName(names ...string) []Element
}
