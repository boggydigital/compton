package compton

import (
	"github.com/boggydigital/compton/consts/attr"
	"io"
	"strings"
)

type Attributes struct {
	attributes map[string]string
}

func (a *Attributes) SetAttribute(name, val string) {
	if name == attr.Class {
		panic("class attribute should be set with ClassList methods")
	}
	if a.attributes == nil {
		a.attributes = make(map[string]string)
	}
	a.attributes[name] = val
}

func (a *Attributes) GetAttribute(name string) string {
	return a.attributes[name]
}

func (a *Attributes) Write(w io.Writer) error {
	attrs := make([]string, 0, len(a.attributes))
	for name, val := range a.attributes {
		attrs = append(attrs, name+"='"+val+"'")
	}
	if _, err := w.Write([]byte(strings.Join(attrs, " "))); err != nil {
		return err
	}
	return nil
}
