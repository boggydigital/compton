package compton

import "io"

type Registrar interface {
	Register(name, extends string, template []byte, mode EncapsulationMode, w io.Writer) error
}
