package compton

import (
	"embed"
	"github.com/boggydigital/compton/consts/attr"
	"golang.org/x/net/html/atom"
	"io"
	"path"
)

const (
	htmlExt = ".html"
	cssExt  = ".css"
)

type BaseElement struct {
	Attributes
	ClassList
	Children []Element
	TagName  atom.Atom
	Markup   embed.FS
	Filename string
}

func (be *BaseElement) Append(children ...Element) {
	be.Children = append(be.Children, children...)
}

func (be *BaseElement) HasChildren() bool {
	return len(be.Children) > 0
}

func (be *BaseElement) Write(w io.Writer) error {
	file, err := be.Markup.Open(be.Filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return WriteContents(file, w, be.WriteFragment)
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
		// set class attribute
		if be.Attributes.attributes == nil {
			be.Attributes.attributes = make(map[string]string)
		}
		if len(be.ClassList.classList) > 0 {
			be.Attributes.attributes[attr.Class] = be.ClassList.String()
		}
		if err := be.Attributes.Write(w); err != nil {
			return err
		}
	default:
		return ErrUnknownToken(t)
	}
	return nil
}

func (be *BaseElement) SetId(id string) {
	be.SetAttribute(attr.Id, id)
}

func (be *BaseElement) GetTagName() atom.Atom {
	return be.TagName
}

func (be *BaseElement) GetElementById(id string) Element {
	for _, child := range be.Children {
		if cid := child.GetAttribute(attr.Id); cid == id {
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

func (be *BaseElement) GetFirstElementByTagName(tagName atom.Atom) Element {
	if matches := be.GetElementsByTagName(tagName); len(matches) > 0 {
		return matches[0]
	}
	return nil
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

func NewElement(a atom.Atom, markup embed.FS, fn string) Element {
	return &BaseElement{
		Markup:   markup,
		TagName:  a,
		Filename: fn,
	}
}

func atomMarkupFilename(a atom.Atom) string {
	if an := a.String(); an != "" {
		return path.Join("markup", a.String()+htmlExt)
	}
	panic("unknown atom markup")
}

func atomStyleFilename(a atom.Atom) string {
	if an := a.String(); an != "" {
		return path.Join("style", a.String()+cssExt)
	}
	panic("unknown atom style")
}

func newElementAtom(a atom.Atom, markup embed.FS) Element {
	return NewElement(a, markup, atomMarkupFilename(a))
}
