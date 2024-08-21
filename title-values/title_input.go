package title_values

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/els"
	"github.com/boggydigital/compton/input_types"
)

type TitleInput struct {
	*TitleValues
	input *els.Input
}

func (ti *TitleInput) SetDataList(list map[string]string) *TitleInput {
	ti.input.SetDataList(list)
	return ti
}

func NewSearchInput(wcr compton.Registrar, title, inputId string) *TitleInput {
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
