package title_values

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/elements/els"
)

type TitleInputElement struct {
	*TitleValuesElement
	input *els.InputElement
}

func (ti *TitleInputElement) SetDataList(list map[string]string) *TitleInputElement {
	ti.input.SetDataList(list)
	return ti
}

func Search(wcr compton.Registrar, title, inputId string) *TitleInputElement {
	titleInput := &TitleInputElement{
		TitleValuesElement: &TitleValuesElement{
			BaseElement: compton.BaseElement{
				Markup:  markupTitleValues,
				TagName: compton_atoms.TitleValues,
			},
			wcr: wcr,
		},
	}

	label := els.Label(inputId)
	label.Append(els.HeadingText(title, 3))
	titleInput.title = label

	input := els.Input(input_types.Search)
	input.
		SetPlaceholder(title).
		SetName(inputId).
		SetId(inputId)

	titleInput.Append(input)
	titleInput.input = input

	return titleInput
}

func SearchValue(wcr compton.Registrar, title, inputId, value string) *TitleInputElement {
	titleInput := Search(wcr, title, inputId)
	titleInput.input.SetAttr(compton.ValueAttr, value)
	return titleInput
}
