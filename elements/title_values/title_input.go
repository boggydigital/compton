package title_values

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/inputs"
)

type TitleInputElement struct {
	*TitleValuesElement
	input *inputs.InputElement
}

func (ti *TitleInputElement) SetDatalist(list map[string]string, listId string) *TitleInputElement {
	ti.input.SetDatalist(list, listId)
	return ti
}

func Search(r compton.Registrar, title, inputId string) *TitleInputElement {
	titleInput := &TitleInputElement{
		TitleValuesElement: TitleValues(r, title),
	}

	label := els.Label(inputId)
	heading := els.HeadingText(title, 3)
	heading.SetId(title)
	label.Append(heading)
	titleInput.title = label

	input := inputs.Input(r, input_types.Search)
	input.SetPlaceholder(title).
		SetName(inputId).
		SetId(inputId)

	titleInput.Append(input)
	titleInput.input = input

	return titleInput
}

func SearchValue(wcr compton.Registrar, title, inputId, value string) *TitleInputElement {
	titleInput := Search(wcr, title, inputId)
	titleInput.input.SetAttribute(attr.Value, value)
	return titleInput
}
