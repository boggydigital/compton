package human_resource

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"io"
)

const (
	elementName = "human-resource"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/human-resource.html"
	markupHumanResource []byte
)

type HumanResource struct {
	wcr        compton.Registrar
	FirstName  string
	LastName   string
	Department string
}

func (hr *HumanResource) Add(_ ...compton.Component) {
	// do nothing
}

func (hr *HumanResource) Write(w io.Writer) error {

	if err := hr.wcr.Register(elementName, markupTemplate, compton.Closed, w); err != nil {
		return err
	}

	if err := compton.WriteContents(bytes.NewReader(markupHumanResource), w, hr.writeHumanResourceFragment); err != nil {
		return err
	}

	return nil
}

func (hr *HumanResource) writeHumanResourceFragment(t string, w io.Writer) error {
	switch t {
	case ".FirstName":
		if _, err := io.WriteString(w, hr.FirstName); err != nil {
			return err
		}
	case ".LastName":
		if _, err := io.WriteString(w, hr.LastName); err != nil {
			return err
		}
	case ".Department":
		if _, err := io.WriteString(w, hr.Department); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

func New(wcr compton.Registrar, firstName, lastName, department string) compton.Component {
	return &HumanResource{
		wcr:        wcr,
		FirstName:  firstName,
		LastName:   lastName,
		Department: department,
	}
}
