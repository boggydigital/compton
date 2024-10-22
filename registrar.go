package compton

import (
	"embed"
)

type Registrar interface {
	//IsRegistered(name string) bool
	RegisterStyle(name string, efs embed.FS)
	RegisterRequirement(name string, element Element)
	RegisterDeferral(name string, elements Element)
}
