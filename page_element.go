package compton

import (
	"net/http"
)

type PageElement interface {
	Element
	Registrar

	SetBodyId(id string) PageElement

	AppendManifest() PageElement
	AppendIcon() PageElement

	WriteResponse(w http.ResponseWriter) error
}
