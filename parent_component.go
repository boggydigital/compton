package compton

type Parent struct {
	Children []Component
}

func (pc *Parent) Append(children ...Component) {
	pc.Children = append(pc.Children, children...)
}
