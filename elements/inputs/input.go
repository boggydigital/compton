package inputs

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/weight"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/exp/maps"
	"golang.org/x/net/html/atom"
	"io"
	"slices"
	"strconv"
	"time"
)

const (
	registrationName      = "input"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/input.html"
	markupInput []byte
	//go:embed "style/inputs.css"
	styleInputs []byte
)

type InputElement struct {
	compton.BaseElement
	r        compton.Registrar
	it       input_types.Type
	dataList compton.Element
}

func (ie *InputElement) WriteStyles(w io.Writer) error {
	if ie.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleInputs, styleRegistrationName).WriteContent(w); err != nil {
			return err
		}
	}
	return ie.BaseElement.WriteStyles(w)
}

func (ie *InputElement) SetPlaceholder(placeholder string) *InputElement {
	ie.SetAttribute("placeholder", placeholder)
	return ie
}

func (ie *InputElement) SetName(name string) *InputElement {
	ie.SetAttribute("name", name)
	return ie
}

func (ie *InputElement) SetValue(value string) *InputElement {
	ie.SetAttribute("value", value)
	return ie
}

func (ie *InputElement) SetChecked(condition bool) *InputElement {
	if condition {
		ie.SetAttribute("checked", "")
	}
	return ie
}

func (ie *InputElement) SetDataList(list map[string]string, listId string) *InputElement {

	if listId == "" {
		listId = ie.GetAttribute(compton.IdAttr)
		if listId == "" {
			listId = strconv.FormatInt(time.Now().Unix(), 10)
		}
		listId += "-list"
	}

	if len(list) > 0 {
		dataList := els.DataList(listId)

		values := maps.Keys(list)
		slices.Sort(values)

		for _, value := range values {
			dataList.Append(els.Option(value, list[value]))
		}
		ie.dataList = dataList
	}

	ie.SetAttribute(compton.ListAttr, listId)

	return ie
}

func (ie *InputElement) WriteDeferrals(w io.Writer) error {
	if ie.dataList != nil {
		return ie.dataList.WriteContent(w)
	}
	return nil
}

func (ie *InputElement) FontWeight(w weight.Weight) *InputElement {
	ie.AddClass(class.FontWeight(w))
	return ie
}

func Input(r compton.Registrar, it input_types.Type) *InputElement {
	input := &InputElement{
		BaseElement: compton.BaseElement{
			TagName: atom.Input,
			Markup:  markupInput,
		},
		r:  r,
		it: it,
	}
	input.SetAttribute(compton.TypeAttr, it.String())
	return input
}

func InputValue(r compton.Registrar, it input_types.Type, value string) *InputElement {
	input := Input(r, it)
	input.SetAttribute(compton.ValueAttr, value)
	return input
}

func Switch(r compton.Registrar) *InputElement {
	input := Input(r, input_types.Checkbox)
	input.SetAttribute("switch", "")
	return input
}
