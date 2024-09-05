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

type IssaImageElement struct {
	compton.BaseElement
	r          compton.Registrar
	dehydrated bool
}

func (ii *IssaImageElement) WriteRequirements(w io.Writer) error {
	if ii.r.RequiresRegistration(elementName) {
		hcScript := els.Script(issa.HydrateColorScript)
		if err := hcScript.WriteContent(w); err != nil {
			return err
		}
		hiScript := els.Script(hydrateImageScript)
		if err := hiScript.WriteContent(w); err != nil {
			return err
		}
		ifiScript := els.Script(imageFadeInScript)
		return ifiScript.WriteContent(w)
	}
	return nil
}

func IssaImageHydrated(r compton.Registrar, placeholder, poster string) compton.Element {

	ii := &IssaImageElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.IssaImage,
			Markup:  templateMarkup,
		},
		r:          r,
		dehydrated: false,
	}

	placeholderImg := els.Image(placeholder)
	placeholderImg.SetClass("placeholder")
	posterImg := els.ImageLazy(poster)
	posterImg.SetClass("poster", "loading")
	ii.Append(placeholderImg, posterImg)

	return ii
}

func IssaImageDehydrated(r compton.Registrar, placeholder, poster string) compton.Element {

	ii := &IssaImageElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.IssaImage,
			Markup:  templateMarkup,
		},
		r:          r,
		dehydrated: true,
	}

	placeholderImg := els.Image("")
	placeholderImg.SetClass("placeholder", "loading")
	placeholderImg.SetAttr("data-dehydrated", placeholder)
	posterImg := els.ImageLazy(poster)
	posterImg.SetClass("poster", "loading")
	ii.Append(placeholderImg, posterImg)

	return ii
}
