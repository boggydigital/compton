package compton

import (
	"io"
	"strings"
)

const (
	idAttr    = "id"
	classAttr = "class"
)

func WriteAttr(w io.Writer, attr, val string) error {
	attrValStr := attr + "='" + val + "'"
	_, err := w.Write([]byte(attrValStr))
	return err
}

func WriteId(w io.Writer, id string) error {
	return WriteAttr(w, idAttr, id)
}

func WriteClassList(w io.Writer, classList ...string) error {
	return WriteAttr(w, classAttr, strings.Join(classList, " "))
}
