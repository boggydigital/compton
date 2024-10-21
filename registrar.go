package compton

type Registrar interface {
	//IsRegistered(name string) bool
	RegisterStyle(name string, style []byte)
	RegisterRequirement(name string, element Element)
	RegisterDeferral(name string, elements Element)
}
