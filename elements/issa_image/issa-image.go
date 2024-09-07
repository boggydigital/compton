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
	registrationName       = "issa-image"
	styleRegistrationName  = "style-" + registrationName
	scriptRegistrationName = "script-" + registrationName
)

var (
	//go:embed "script/image_fadein.js"
	scriptImageFadeIn []byte
	//go:embed "script/hydrate_images.js"
	scriptHydrateImage []byte
	//go:embed "markup/issa-image.html"
	markupIssaImage []byte
	//go:embed "style/issa-image.css"
	styleIssaImage []byte
)

type IssaImageElement struct {
	compton.BaseElement
	r          compton.Registrar
	dehydrated bool
}

func (ii *IssaImageElement) WriteStyles(w io.Writer) error {
	if ii.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleIssaImage).WriteContent(w); err != nil {
			return err
		}
	}
	return nil
}

func (ii *IssaImageElement) WriteDeferrals(w io.Writer) error {
	if ii.r.RequiresRegistration(scriptRegistrationName) {
		hcScript := els.Script(issa.HydrateColorScript)
		if err := hcScript.WriteContent(w); err != nil {
			return err
		}
		hiScript := els.Script(scriptHydrateImage)
		if err := hiScript.WriteContent(w); err != nil {
			return err
		}
		ifiScript := els.Script(scriptImageFadeIn)
		return ifiScript.WriteContent(w)
	}
	return nil
}

func IssaImageHydrated(r compton.Registrar, placeholder, poster string) compton.Element {

	ii := &IssaImageElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.IssaImage,
			Markup:  markupIssaImage,
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
			Markup:  markupIssaImage,
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
