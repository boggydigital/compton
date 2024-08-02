package heading

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/text"
	"io"
	"strconv"
)

var (
	//go:embed "markup/hx.html"
	markupHx []byte
)

type Heading struct {
	compton.AP
	level int
}

func (h *Heading) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupHx), w, h.writeHeadingFragment)
}

func (h *Heading) writeHeadingFragment(t string, w io.Writer) error {
	switch t {
	case ".Level":
		if _, err := io.WriteString(w, strconv.Itoa(h.level)); err != nil {
			return err
		}
	case ".Attributes":
		if err := h.Attributes.Write(w); err != nil {
			return err
		}
	case ".Content":
		for _, child := range h.Children {
			if err := child.Write(w); err != nil {
				return err
			}
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func New(level int) compton.Element {
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	return &Heading{
		level: level,
	}
}

func NewText(data string, level int) compton.Element {
	heading := New(level)
	heading.Append(text.New(data))
	return heading
}
