package compton

import "io"

type Component interface {
	Append(children ...Component) Component
	Write(w io.Writer) error
}
