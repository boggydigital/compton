package issa_image

import (
	_ "embed"
	"github.com/boggydigital/compton"
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
)

type IssaImage struct {
	compton.BaseElement
	r          compton.Registrar
	dehydrated bool
	container  compton.Element
}

func (ii *IssaImage) WriteRequirements(w io.Writer) error {
	if ii.r.RequiresRegistration(elementName) {
		if ii.dehydrated {
			hcScript := els.NewScript(issa.HydrateColorScript)
			if err := hcScript.WriteContent(w); err != nil {
				return err
			}
			hiScript := els.NewScript(hydrateImageScript)
			if err := hiScript.WriteContent(w); err != nil {
				return err
			}
		}
		ifiScript := els.NewScript(imageFadeInScript)
		return ifiScript.WriteContent(w)
	}
	return nil
}

func (ii *IssaImage) WriteContent(w io.Writer) error {
	return ii.container.WriteContent(w)
}

func NewHydrated(r compton.Registrar, placeholder, poster string) compton.Element {
	container := els.NewDiv()
	container.SetClass(elementName)
	placeholderImg := els.NewImage(placeholder)
	placeholderImg.SetClass("placeholder")
	posterImg := els.NewImage(poster)
	posterImg.SetClass("poster", "loading")
	container.Append(placeholderImg, posterImg)

	return &IssaImage{
		r:          r,
		dehydrated: false,
		container:  container,
	}
}

func NewDehydrated(r compton.Registrar, placeholder, poster string) compton.Element {

	container := els.NewDiv()
	container.SetClass(elementName)
	placeholderImg := els.NewImage("")
	placeholderImg.SetClass("placeholder")
	placeholderImg.SetAttr("data-dehydrated", placeholder)
	posterImg := els.NewImage(poster)
	posterImg.SetClass("poster", "loading")
	container.Append(placeholderImg, posterImg)

	return &IssaImage{
		r:          r,
		dehydrated: true,
		container:  container,
	}
}
