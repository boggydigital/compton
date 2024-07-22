package web_components

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"io"
)

var (
	//go:embed "markup/script.html"
	markupScript []byte
)

type Registrar struct {
	registry map[string]any
}

func (cr *Registrar) Register(elementName string, markupTemplate []byte, mode EncapsulationMode, w io.Writer) error {

	if _, ok := cr.registry[elementName]; ok {
		return nil
	}

	if err := compton.WriteContents(bytes.NewReader(markupScript), w, func(token string, w io.Writer) error {
		switch token {
		case ".ElementName":
			if _, err := io.WriteString(w, elementName); err != nil {
				return err
			}
		case ".Mode":
			if _, err := io.WriteString(w, string(mode)); err != nil {
				return err
			}
		default:
			return compton.ErrUnknownToken(token)
		}
		return nil
	}); err != nil {
		return err
	}

	if _, err := w.Write(markupTemplate); err != nil {
		return err
	}

	cr.registry[elementName] = nil

	return nil
}

func NewComponentRegistrar() *Registrar {
	return &Registrar{
		registry: make(map[string]any),
	}
}
