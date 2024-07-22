package compton

import "io"

type Component interface {
	Add(children ...Component)
	//AddCustomStyles(styles []byte)
	Write(w io.Writer) error
}
