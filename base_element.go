package compton

import (
	"bytes"
	"io"
	"strings"
)

type BaseElement struct {
	Attributes
	Parent
	Markup []byte
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

func NewElement(markup []byte) Element {
	return &BaseElement{Markup: markup}
}
