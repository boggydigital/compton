package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/input_types"
	"golang.org/x/exp/maps"
	"golang.org/x/net/html/atom"
	"slices"
	"strconv"
	"time"
)

const (
	rnInput       = "input"
	rnDatalistPfx = "datalist-"
)

var (
	//go:embed "style/input.css"
	styleInputs []byte
)

type InputElement struct {
	*BaseElement
	r        Registrar
	it       input_types.Type
	dataList Element
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

func (ie *InputElement) SetDisabled(condition bool) *InputElement {
	if condition {
		ie.SetAttribute("disabled", "")
	}
	return ie
}

func (ie *InputElement) SetDatalist(list map[string]string, listId string) *InputElement {

	if listId == "" {
		listId = ie.GetAttribute(attr.Id)
		if listId == "" {
			listId = strconv.FormatInt(time.Now().Unix(), 10)
		}
		listId += "-list"
	}

	if len(list) > 0 {
		dataList := Datalist(listId)

		values := maps.Keys(list)
		slices.Sort(values)

		for _, value := range values {
			dataList.Append(Option(value, list[value]))
		}
		ie.dataList = dataList
	}

	ie.SetAttribute(attr.List, listId)

	ie.r.RegisterDeferrals(rnDatalistPfx+listId, ie.dataList)

	return ie
}

func (ie *InputElement) FontWeight(w font_weight.Weight) *InputElement {
	ie.AddClass(class.FontWeight(w))
	return ie
}

func Input(r Registrar, it input_types.Type) *InputElement {
	input := &InputElement{
		BaseElement: NewElement(tacMarkup(atom.Input)),
		r:           r,
		it:          it,
	}
	input.SetAttribute(attr.Type, it.String())

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(atom.Input))

	return input
}

func InputValue(r Registrar, it input_types.Type, value string) *InputElement {
	input := Input(r, it)
	input.SetAttribute(attr.Value, value)
	return input
}

func Switch(r Registrar) *InputElement {
	input := Input(r, input_types.Checkbox)
	input.SetAttribute("switch", "")
	return input
}
