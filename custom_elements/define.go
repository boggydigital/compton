package custom_elements

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"io"
)

var (
	//go:embed "markup/define_script.html"
	markupDefineScript []byte
)

func Define(w io.Writer, opt *Options) error {

	if err := Validate(opt); err != nil {
		return err
	}

	if err := compton.WriteContents(
		bytes.NewReader(markupDefineScript), w,
		func(token string, w io.Writer) error {
			switch token {
			case ".ElementName":
				if _, err := io.WriteString(w, opt.elementName); err != nil {
					return err
				}
			case ".ExtendsElement":
				if _, err := io.WriteString(w, opt.extendsElement); err != nil {
					return err
				}
			case ".Mode":
				if _, err := io.WriteString(w, string(opt.encapsulationMode)); err != nil {
					return err
				}
			default:
				return compton.ErrUnknownToken(token)
			}
			return nil
		}); err != nil {
		return err
	}

	return nil
}
