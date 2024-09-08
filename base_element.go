package compton

import (
	"bytes"
	"golang.org/x/net/html/atom"
	"io"
)

type BaseElement struct {
	Attributes
	ClassList
	Children []Element
	TagName  atom.Atom
	Markup   []byte
}

func (be *BaseElement) Append(children ...Element) {
	be.Children = append(be.Children, children...)
}

func (be *BaseElement) WriteRequirements(w io.Writer) error {
	for _, child := range be.Children {
		if err := child.WriteRequirements(w); err != nil {
			return err
		}
	}
	return nil
}

func (be *BaseElement) WriteStyles(w io.Writer) error {
	for _, child := range be.Children {
		if err := child.WriteStyles(w); err != nil {
			return err
		}
	}
	return nil
}

func (be *BaseElement) WriteContent(w io.Writer) error {
	return WriteContents(bytes.NewReader(be.Markup), w, be.WriteFragment)
}

func (be *BaseElement) WriteDeferrals(w io.Writer) error {
	for _, child := range be.Children {
		if err := child.WriteDeferrals(w); err != nil {
			return err
		}
	}
	return nil
}

func (be *BaseElement) WriteFragment(t string, w io.Writer) error {
	switch t {
	case ContentToken:
		for _, child := range be.Children {
			if err := child.WriteContent(w); err != nil {
				return err
			}
		}
	case AttributesToken:
		// set class attribute
		if be.Attributes.attributes == nil {
			be.Attributes.attributes = make(map[string]string)
		}
		be.Attributes.attributes[ClassAttr] = be.ClassList.String()
		if err := be.Attributes.Write(w); err != nil {
			return err
		}
	default:
		return ErrUnknownToken(t)
	}
	return nil
}

func (be *BaseElement) SetId(id string) {
	be.SetAttribute(IdAttr, id)
}

func (be *BaseElement) GetTagName() atom.Atom {
	return be.TagName
}

func (be *BaseElement) GetElementById(id string) Element {
	for _, child := range be.Children {
		if cid := child.GetAttribute(IdAttr); cid == id {
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
