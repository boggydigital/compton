package text

import (
	"github.com/boggydigital/compton"
	"io"
)

type Text struct {
	content string
}

func (t *Text) Append(_ ...compton.Component) compton.Component {
	// do nothing
	return nil
}

func (t *Text) Write(w io.Writer) error {
	if _, err := io.WriteString(w, t.content); err != nil {
		return err
	}
	return nil
}

func New(content string) compton.Component {
	return &Text{
		content: content,
	}
}
