package compton

import (
	"bytes"
	"golang.org/x/net/html/atom"
	"io"
	"strings"
)

type BaseElement struct {
	Attributes
	Children []Element
	TagName  atom.Atom
	Markup   []byte
}

func (be *BaseElement) Append(children ...Element) Element {
	be.Children = append(be.Children, children...)
	return be
}

func (be *BaseElement) Register(w io.Writer) error {
	for _, child := range be.Children {
		if err := child.Register(w); err != nil {
			return err
		}
	}
	return nil
}

func (be *BaseElement) Write(w io.Writer) error {
	return WriteContents(bytes.NewReader(be.Markup), w, be.WriteFragment)
}

func (be *BaseElement) WriteFragment(t string, w io.Writer) error {
	switch t {
	case ContentToken:
		for _, child := range be.Children {
			if err := child.Write(w); err != nil {
				return err
			}
		}
	case AttributesToken:
		if err := be.Attributes.Write(w); err != nil {
			return err
		}
	default:
		return ErrUnknownToken(t)
	}
	return nil
}

func (be *BaseElement) SetId(id string) Element {
	be.SetAttr(IdAttr, id)
	return be
}

func (be *BaseElement) SetClass(names ...string) Element {
	be.SetAttr(ClassAttr, strings.Join(names, " "))
	return be
}

func (be *BaseElement) HasClass(names ...string) bool {
	class := be.GetAttr(ClassAttr)
	for _, name := range names {
		if !strings.Contains(class, name) {
			return false
		}
	}
	return true
}

func (be *BaseElement) GetTagName() atom.Atom {
	return be.TagName
}

func (be *BaseElement) GetElementById(id string) Element {
	for _, child := range be.Children {
		if cid := child.GetAttr(IdAttr); cid == id {
			return child
		}
		if el := child.GetElementById(id); el != nil {
			return el
		}
	}
	return nil
}

func (be *BaseElement) GetElementsByTagName(tagName atom.Atom) []Element {
	matches := make([]Element, 0)
	for _, child := range be.Children {
		if child.GetTagName() == tagName {
			matches = append(matches, child)
		}
		matches = append(matches, child.GetElementsByTagName(tagName)...)
	}
	return matches
}

func (be *BaseElement) GetElementsByClassName(names ...string) []Element {
	matches := make([]Element, 0)
	for _, child := range be.Children {
		if child.HasClass(names...) {
			matches = append(matches, child)
		}
		matches = append(matches, child.GetElementsByClassName(names...)...)
	}
	return matches
}

func NewElement(a atom.Atom, markup []byte) Element {
	return &BaseElement{Markup: markup, TagName: a}
}
