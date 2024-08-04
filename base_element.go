package compton

import (
	"bytes"
	"golang.org/x/net/html/atom"
	"io"
	"strings"
)

type BaseElement struct {
	Attributes
	Parent
	TagName atom.Atom
	Markup  []byte
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

func (be *BaseElement) SetId(id string) {
	be.SetAttr(IdAttr, id)
}

func (be *BaseElement) SetClass(classes ...string) {
	be.SetAttr(ClassAttr, strings.Join(classes, " "))
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

func NewElement(a atom.Atom, markup []byte) Element {
	return &BaseElement{Markup: markup, TagName: a}
}
