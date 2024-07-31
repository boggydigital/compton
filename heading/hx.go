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
	compton.Parent
	Id        string
	ClassList []string
	Level     string
}

func (h *Heading) Append(children ...compton.Component) compton.Component {
	h.Children = append(h.Children, children...)
	return h
}

func (h *Heading) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupHx), w, h.writeHeadingFragment)
}

func (h *Heading) writeHeadingFragment(t string, w io.Writer) error {
	switch t {
	case ".Level":
		if _, err := io.WriteString(w, h.Level); err != nil {
			return err
		}
	case ".Id":
		if h.Id != "" {
			if err := compton.WriteId(w, h.Id); err != nil {
				return err
			}
		}
	case ".ClassList":
		if len(h.ClassList) > 0 {
			if err := compton.WriteClassList(w, h.ClassList...); err != nil {
				return err
			}
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

func New(level int, id string, classList ...string) compton.Component {
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	return &Heading{
		Level:     strconv.Itoa(level),
		Id:        id,
		ClassList: classList,
	}
}

func NewText(data string, level int, id string, classList ...string) compton.Component {
	return New(level, id, classList...).Append(text.New(data))
}
