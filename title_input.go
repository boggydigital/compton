package compton

import (
	"github.com/boggydigital/compton/consts/attr"
	"github.com/boggydigital/compton/consts/input_types"
)

type TitleInputElement struct {
	*TitleValuesElement
	input *InputElement
}

func (ti *TitleInputElement) SetDatalist(list map[string]string, listId string) *TitleInputElement {
	ti.input.SetDatalist(list, listId)
	return ti
}

func newTitleInput(r Registrar, inputType input_types.Type, title, inputId string) *TitleInputElement {
	titleInput := &TitleInputElement{
		TitleValuesElement: TitleValues(r, title),
	}

	label := Label(inputId)
	heading := HeadingText(title, 3)
	heading.SetId(title)
	label.Append(heading)
	titleInput.title = label

	input := Input(r, inputType)
	input.SetPlaceholder(title).
		SetName(inputId).
		SetId(inputId)

	titleInput.Append(input)
	titleInput.input = input

	return titleInput
}

func TISearch(r Registrar, title, inputId string) *TitleInputElement {
	return newTitleInput(r, input_types.Search, title, inputId)
}

func TISearchValue(wcr Registrar, title, inputId, value string) *TitleInputElement {
	titleInput := TISearch(wcr, title, inputId)
	titleInput.input.SetAttribute(attr.Value, value)
	return titleInput
}

func TIText(r Registrar, title, inputId string) *TitleInputElement {
	return newTitleInput(r, input_types.Text, title, inputId)
}

func TIPassword(r Registrar, title, inputId string) *TitleInputElement {
	return newTitleInput(r, input_types.Password, title, inputId)
}
