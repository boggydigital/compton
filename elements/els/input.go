package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/input_types"
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

type InputElement struct {
	compton.BaseElement
	it       input_types.Type
	dataList compton.Element
}

func (i *InputElement) SetPlaceholder(placeholder string) *InputElement {
	i.SetAttribute("placeholder", placeholder)
	return i
}

func (i *InputElement) SetName(name string) *InputElement {
	i.SetAttribute("name", name)
	return i
}

func (i *InputElement) SetDataList(list map[string]string) *InputElement {

	listId := i.GetAttribute(compton.IdAttr)
	if listId == "" {
		listId = strconv.FormatInt(time.Now().Unix(), 10)
	}
	listId += "-list"
	dataList := DataList(listId)

	values := maps.Keys(list)
	slices.Sort(values)

	for _, value := range values {
		dataList.Append(Option(value, list[value]))
	}
	i.dataList = dataList
	i.SetAttribute(compton.ListAttr, listId)

	return i
}

func (i *InputElement) WriteDeferrals(w io.Writer) error {
	if i.dataList != nil {
		return i.dataList.WriteContent(w)
	}
	return nil
}

func Input(it input_types.Type) *InputElement {
	input := &InputElement{
		BaseElement: compton.BaseElement{
			TagName: atom.Input,
			Markup:  markupInput,
		},
		it: it,
	}
	input.SetAttribute(compton.TypeAttr, it.String())
	return input
}

func InputValue(it input_types.Type, value string) *InputElement {
	input := Input(it)
	input.SetAttribute(compton.ValueAttr, value)
	return input
}
