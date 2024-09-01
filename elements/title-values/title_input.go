package title_values

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/elements/els"
)

type TitleInput struct {
	*TitleValues
	input *els.Input
}

func (ti *TitleInput) SetDataList(list map[string]string) *TitleInput {
	ti.input.SetDataList(list)
	return ti
}

func NewSearch(wcr compton.Registrar, title, inputId string) *TitleInput {
	titleInput := &TitleInput{
		TitleValues: &TitleValues{
			BaseElement: compton.BaseElement{
				Markup:  markupTitleValues,
				TagName: compton_atoms.TitleValues,
			},
			wcr: wcr,
		},
	}

	label := els.NewLabel(inputId)
	label.Append(els.NewHeadingText(title, 3))
	titleInput.title = label

	input := els.NewInput(input_types.Search)
	input.
		SetPlaceholder(title).
		SetName(inputId).
		SetId(inputId)

	titleInput.Append(input)
	titleInput.input = input

	return titleInput
}

func NewSearchValue(wcr compton.Registrar, title, inputId, value string) *TitleInput {
	titleInput := NewSearch(wcr, title, inputId)
	titleInput.input.SetAttr(compton.ValueAttr, value)
	return titleInput
}
