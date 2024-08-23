package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/input_types"
	"golang.org/x/exp/maps"
	"golang.org/x/net/html/atom"
	"io"
	"slices"
	"strconv"
	"time"
)

var (
	//go:embed "markup/input.html"
	markupInput []byte
)

type Input struct {
	compton.BaseElement
	it       input_types.Type
	dataList compton.Element
}

func (i *Input) SetPlaceholder(placeholder string) *Input {
	i.SetAttr("placeholder", placeholder)
	return i
}

func (i *Input) SetName(name string) *Input {
	i.SetAttr("name", name)
	return i
}

func (i *Input) SetDataList(list map[string]string) *Input {

	listId := i.GetAttr(compton.IdAttr)
	if listId == "" {
		listId = strconv.FormatInt(time.Now().Unix(), 10)
	}
	listId += "-list"
	dataList := NewDataList(listId)

	values := maps.Keys(list)
	slices.Sort(values)

	for _, value := range values {
		dataList.Append(NewOption(value, list[value]))
	}
	i.dataList = dataList
	i.SetAttr(compton.ListAttr, listId)

	return i
}

func (i *Input) WriteDeferrals(w io.Writer) error {
	if i.dataList != nil {
		return i.dataList.WriteContent(w)
	}
	return nil
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

func NewInputValue(it input_types.Type, value string) *Input {
	input := NewInput(it)
	input.SetAttr(compton.ValueAttr, value)
	return input
}
