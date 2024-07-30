package compton

import (
	"bytes"
	"io"
)

type Container struct {
	Parent
	markup       []byte
	contentToken string
}

func (c *Container) Append(children ...Component) Component {
	c.Children = append(c.Children, children...)
	return c
}

func (c *Container) Write(w io.Writer) error {
	return WriteContents(bytes.NewReader(c.markup), w, c.writeFragment)
}

func (c *Container) writeFragment(t string, w io.Writer) error {
	switch t {
	case c.contentToken:
		for _, child := range c.Children {
			if err := child.Write(w); err != nil {
				return err
			}
		}
	default:
		return ErrUnknownToken(t)
	}
	return nil
}

func NewContainer(markup []byte, contentToken string) Component {
	return &Container{
		markup:       markup,
		contentToken: contentToken,
	}
}
