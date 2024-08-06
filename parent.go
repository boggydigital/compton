package compton

import "io"

type Parent struct {
	Children []Element
}

func (pc *Parent) Append(children ...Element) {
	pc.Children = append(pc.Children, children...)
}

func (pc *Parent) Register(w io.Writer) error {
	for _, child := range pc.Children {
		if err := child.Register(w); err != nil {
			return err
		}
	}
	return nil
}
