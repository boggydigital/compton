package compton

import (
	"net/http"
)

type PageElement interface {
	Element
	Registrar

	SetBodyId(id string) PageElement

	AppendStyle(id string, style []byte) PageElement
	AppendManifest() PageElement
	AppendIcon() PageElement

	WriteResponse(w http.ResponseWriter) error
}
