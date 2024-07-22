package compton

import "io"

type Component interface {
	Append(children ...Component)
	Write(w io.Writer) error
}
