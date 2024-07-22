package compton

type Parent struct {
	Children []Component
}

func (pc *Parent) Add(children ...Component) {
	pc.Children = append(pc.Children, children...)
}
