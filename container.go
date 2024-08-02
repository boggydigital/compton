package compton

import (
	"bytes"
	"io"
)

type Container struct {
	AP
	markup       []byte
	contentToken string
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
	case ".Attributes":
		if err := c.Attributes.Write(w); err != nil {
			return err
		}
	default:
		return ErrUnknownToken(t)
	}
	return nil
}

func NewContainer(markup []byte, contentToken string) Element {
	return &Container{
		markup:       markup,
		contentToken: contentToken,
	}
}
