package compton

import "io"

type Registrar interface {
	RegisterCustomElement(name, extends string, mode EncapsulationMode, w io.Writer) error
	RegisterMarkup(r io.Reader, w io.Writer, tw TokenWriter) error
}
