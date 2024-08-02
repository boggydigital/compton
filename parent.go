package compton

type Parent struct {
	Children []Element
}

func (pc *Parent) Append(children ...Element) {
	pc.Children = append(pc.Children, children...)
}
