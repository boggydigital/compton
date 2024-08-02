package compton

import "io"

type Element interface {
	Append(children ...Element)
	Write(w io.Writer) error
	SetId(id string)
	SetClass(classes ...string)
	SetAttr(name, val string)
}
