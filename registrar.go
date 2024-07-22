package compton

import "io"

type Registrar interface {
	Register(n string, t []byte, m EncapsulationMode, w io.Writer) error
}
