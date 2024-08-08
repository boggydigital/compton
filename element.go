package compton

import (
	"golang.org/x/net/html/atom"
	"io"
)

type Element interface {
	Append(children ...Element) Element
	Write(w io.Writer) error
	Register(w io.Writer) error

	GetTagName() atom.Atom

	SetId(id string) Element

	SetClass(names ...string) Element
	HasClass(names ...string) bool

	SetAttr(name, val string)
	GetAttr(name string) string

	GetElementById(id string) Element
	GetElementsByTagName(tagName atom.Atom) []Element
	GetElementsByClassName(names ...string) []Element
}
