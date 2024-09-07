package compton

import (
	"golang.org/x/net/html/atom"
	"io"
)

type Element interface {
	Append(children ...Element)
	WriteStyles(w io.Writer) error
	WriteRequirements(w io.Writer) error
	WriteContent(w io.Writer) error
	WriteDeferrals(w io.Writer) error

	GetTagName() atom.Atom

	SetId(id string)

	SetClass(names ...string)
	HasClass(names ...string) bool

	SetAttr(name, val string)
	GetAttr(name string) string

	GetElementById(id string) Element
	GetElementsByTagName(tagName atom.Atom) []Element
	GetElementsByClassName(names ...string) []Element
}
