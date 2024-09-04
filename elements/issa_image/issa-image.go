package issa_image

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/issa"
	"io"
)

const (
	elementName = "issa-image"
)

var (
	//go:embed "script/image_fadein.js"
	imageFadeInScript []byte
	//go:embed "script/hydrate_images.js"
	hydrateImageScript []byte
	//go:embed "markup/template.html"
	templateMarkup []byte
)

type IssaImage struct {
	compton.BaseElement
	r          compton.Registrar
	dehydrated bool
}

func (ii *IssaImage) WriteRequirements(w io.Writer) error {
	if ii.r.RequiresRegistration(elementName) {
		hcScript := els.NewScript(issa.HydrateColorScript)
		if err := hcScript.WriteContent(w); err != nil {
			return err
		}
		hiScript := els.NewScript(hydrateImageScript)
		if err := hiScript.WriteContent(w); err != nil {
			return err
		}
		ifiScript := els.NewScript(imageFadeInScript)
		return ifiScript.WriteContent(w)
	}
	return nil
}

func NewHydrated(r compton.Registrar, placeholder, poster string) compton.Element {

	ii := &IssaImage{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.IssaImage,
			Markup:  templateMarkup,
		},
		r:          r,
		dehydrated: false,
	}

	placeholderImg := els.NewImage(placeholder)
	placeholderImg.SetClass("placeholder")
	posterImg := els.NewImageLazy(poster)
	posterImg.SetClass("poster", "loading")
	ii.Append(placeholderImg, posterImg)

	return ii
}

func NewDehydrated(r compton.Registrar, placeholder, poster string) compton.Element {

	ii := &IssaImage{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.IssaImage,
			Markup:  templateMarkup,
		},
		r:          r,
		dehydrated: true,
	}

	placeholderImg := els.NewImage("")
	placeholderImg.SetClass("placeholder")
	placeholderImg.SetAttr("data-dehydrated", placeholder)
	posterImg := els.NewImageLazy(poster)
	posterImg.SetClass("poster", "loading")
	ii.Append(placeholderImg, posterImg)

	return ii
}
