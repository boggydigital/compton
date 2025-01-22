package compton

import (
	"net/http"
)

type PageElement interface {
	Element
	Registrar

	Error(err error) PageElement

	SetBodyId(id string) PageElement
	SetBodyAttribute(name, val string)

	AppendManifest() PageElement
	AppendIcon() PageElement
	AppendSpeculationRules(hrefMatches ...string)

	WriteResponse(w http.ResponseWriter) error
}
