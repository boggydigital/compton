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
	i.SetAttr("placeholder", placeholder)
	return i
}

func (i *InputElement) SetName(name string) *InputElement {
	i.SetAttr("name", name)
	return i
}

func (i *InputElement) SetDataList(list map[string]string) *InputElement {

	listId := i.GetAttr(compton.IdAttr)
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
	i.SetAttr(compton.ListAttr, listId)

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
	input.SetAttr(compton.TypeAttr, it.String())
	return input
}

func InputValue(it input_types.Type, value string) *InputElement {
	input := Input(it)
	input.SetAttr(compton.ValueAttr, value)
	return input
}
