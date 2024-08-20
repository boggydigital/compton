package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/input_types"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/input.html"
	markupInput []byte
)

type Input struct {
	compton.BaseElement
	it input_types.Type
}

func (i *Input) SetPlaceholder(placeholder string) *Input {
	i.SetAttr("placeholder", placeholder)
	return i
}

func (i *Input) SetName(name string) *Input {
	i.SetAttr("name", name)
	return i
}

func NewInput(it input_types.Type) *Input {
	input := &Input{
		BaseElement: compton.BaseElement{
			TagName: atom.Input,
			Markup:  markupInput,
		},
		it: it,
	}
	input.SetAttr(compton.TypeAttr, it.String())
	return input
}
