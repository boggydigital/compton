package els

import (
	"github.com/boggydigital/compton"
	"io"
)

type Text struct {
	compton.BaseElement
	content string
}

func (t *Text) Append(_ ...compton.Element) {
}

func (t *Text) WriteContent(w io.Writer) error {
	if _, err := io.WriteString(w, t.content); err != nil {
		return err
	}
	return nil
}

func NewText(content string) compton.Element {
	return &Text{
		content: content,
	}
}
