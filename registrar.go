package compton

import (
	"embed"
)

type Registrar interface {
	RegisterStyles(efs embed.FS, names ...string)
	RegisterRequirements(name string, elements ...Element)
	RegisterDeferrals(name string, elements ...Element)
}
