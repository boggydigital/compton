package custom_elements

import (
	"errors"
	"github.com/boggydigital/compton"
)

var (
	ErrNilOptions            = errors.New("custom element options cannot be nil")
	ErrMissingElementName    = errors.New("missing element name")
	ErrMissingExtendsElement = errors.New("missing extends element")
)

type Options struct {
	elementName       string
	extendsElement    string
	encapsulationMode compton.EncapsulationMode
}

func Defaults(name string) *Options {
	return &Options{
		elementName:       name,
		extendsElement:    "HTMLElement",
		encapsulationMode: compton.EncapsulationClosed,
	}
}

func Validate(opt *Options) error {
	if opt == nil {
		return ErrNilOptions
	}
	if opt.elementName == "" {
		return ErrMissingElementName
	}
	if opt.extendsElement == "" {
		return ErrMissingExtendsElement
	}
	return nil
}
