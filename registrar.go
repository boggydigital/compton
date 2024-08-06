package compton

type Registrar interface {
	RequiresRegistration(name string) bool
}
