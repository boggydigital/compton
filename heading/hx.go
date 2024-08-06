package heading

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/text"
	"golang.org/x/net/html/atom"
	"io"
	"strconv"
)

var (
	//go:embed "markup/hx.html"
	markupHx []byte
)

type Heading struct {
	compton.BaseElement
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
	case compton.AttributesToken:
		fallthrough
	case compton.ContentToken:
		if err := h.BaseElement.WriteFragment(t, w); err != nil {
			return err
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
	tn := atom.H1
	switch level {
	case 1:
		tn = atom.H1
	case 2:
		tn = atom.H2
	case 3:
		tn = atom.H3
	case 4:
		tn = atom.H4
	case 5:
		tn = atom.H5
	case 6:
		tn = atom.H6
	}
	return &Heading{
		BaseElement: compton.BaseElement{
			TagName: tn,
		},
		level: level,
	}
}

func NewText(data string, level int) compton.Element {
	heading := New(level)
	heading.Append(text.New(data))
	return heading
}