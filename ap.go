package compton

import "strings"

const (
	IdAttr    = "id"
	ClassAttr = "class"
)

type AP struct {
	Attributes
	Parent
}

func (ap *AP) SetId(id string) {
	ap.SetAttr(IdAttr, id)
}

func (ap *AP) SetClass(classes ...string) {
	ap.SetAttr(ClassAttr, strings.Join(classes, " "))
}
