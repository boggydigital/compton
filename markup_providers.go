package compton

import (
	"embed"
	"errors"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"golang.org/x/net/html/atom"
	"strings"
)

const (
	tacMarkupTemplate        = "<{tag} {{.Attributes}}>{{.Content}}</{tag}>"
	voidTacMarkupTemplate    = "<{tag} {{.Attributes}}/>"
	transparentContentMarkup = "{{.Content}}"
)

type atomsEmbedMarkupProvider struct {
	efs embed.FS
	fn  string
}

func (caemp *atomsEmbedMarkupProvider) GetMarkup() ([]byte, error) {
	bts, err := caemp.efs.ReadFile(caemp.fn)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

func atomsEmbedMarkup(ca atom.Atom, efs embed.FS) (atom.Atom, MarkupProvider) {
	return ca, &atomsEmbedMarkupProvider{
		efs: efs,
		fn:  compton_atoms.MarkupName(ca),
	}
}

type tacMarkupProvider struct {
	a atom.Atom
}

func (tmp *tacMarkupProvider) GetMarkup() ([]byte, error) {
	if an := compton_atoms.Atos(tmp.a); an != "" {
		tacm := strings.Replace(tacMarkupTemplate, "{tag}", an, -1)
		return []byte(tacm), nil
	} else {
		return nil, errors.New("no tag-attribute-content markup for atom")
	}
}

func tacMarkup(a atom.Atom) (atom.Atom, MarkupProvider) {
	return a, &tacMarkupProvider{a: a}
}

type voidTacMarkupProvider struct {
	a atom.Atom
}

func (vtmp *voidTacMarkupProvider) GetMarkup() ([]byte, error) {
	if an := compton_atoms.Atos(vtmp.a); an != "" {
		vtacm := strings.Replace(voidTacMarkupTemplate, "{tag}", an, -1)
		return []byte(vtacm), nil
	} else {
		return nil, errors.New("no void-tag-attribute-content markup for atom")
	}
}

func voidTacMarkup(a atom.Atom) (atom.Atom, MarkupProvider) {
	return a, &voidTacMarkupProvider{a: a}
}

type contentMarkupProvider struct{}

func (cmp *contentMarkupProvider) GetMarkup() ([]byte, error) {
	return []byte(transparentContentMarkup), nil
}

func contentMarkup(a atom.Atom) (atom.Atom, MarkupProvider) {
	return a, &contentMarkupProvider{}
}

type bytesMarkupProvider struct {
	bts []byte
}

func (cmp *bytesMarkupProvider) GetMarkup() ([]byte, error) {
	return cmp.bts, nil
}

func BytesMarkup(a atom.Atom, bts []byte) (atom.Atom, MarkupProvider) {
	return a, &bytesMarkupProvider{bts: bts}
}
