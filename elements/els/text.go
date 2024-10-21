package els

import (
	"github.com/boggydigital/compton"
	"io"
)

type TextElement struct {
	compton.BaseElement
	content string
}

func (t *TextElement) Append(_ ...compton.Element) {
}

func (t *TextElement) Write(w io.Writer) error {
	if _, err := io.WriteString(w, t.content); err != nil {
		return err
	}
	return nil
}

func Text(content string) compton.Element {
	return &TextElement{
		content: content,
	}
}

func TextBytes(content []byte) compton.Element {
	return &TextElement{
		content: string(content),
	}
}
