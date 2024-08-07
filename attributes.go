package compton

import (
	"io"
	"strings"
)

const (
	IdAttr    = "id"
	ClassAttr = "class"
	HrefAttr  = "href"
)

const (
	TrueVal  = "true"
	FalseVal = "false"
)

type Attributes struct {
	attributes map[string]string
}

func (a *Attributes) SetAttr(name, val string) {
	if a.attributes == nil {
		a.attributes = make(map[string]string)
	}
	a.attributes[name] = val
}

func (a *Attributes) GetAttr(name string) string {
	return a.attributes[name]
}

func (a *Attributes) Write(w io.Writer) error {
	attrs := make([]string, 0, len(a.attributes))
	for name, val := range a.attributes {
		attrs = append(attrs, name+"='"+val+"'")
		if _, err := w.Write([]byte(strings.Join(attrs, " "))); err != nil {
			return err
		}
	}
	return nil
}
