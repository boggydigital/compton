package issa_image

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/script"
	"github.com/boggydigital/issa"
)

const (
	registrationName                   = "issa-image"
	styleRegistrationName              = "style-" + registrationName
	scriptImageFadeInRegistrationName  = "script-image-fade-in-" + registrationName
	scriptHydrateColorRegistrationName = "script-hydrate-color-" + registrationName
	scriptHydrateImageRegistrationName = "script-hydrate-image" + registrationName
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
	//r          compton.Registrar
	dehydrated bool
}

//func (ii *IssaImageElement) WriteStyles(w io.Writer) error {
//	if ii.r.RequiresRegistration(styleRegistrationName) {
//		if err := els.Style(styleIssaImage, styleRegistrationName).Write(w); err != nil {
//			return err
//		}
//	}
//	return nil
//}

//func (ii *IssaImageElement) WriteDeferrals(w io.Writer) error {
//	if ii.r.RequiresRegistration(scriptRegistrationName) {
//		hcScript := script.ScriptAsync(issa.HydrateColorScript)
//		if err := hcScript.Write(w); err != nil {
//			return err
//		}
//		hiScript := script.ScriptAsync(scriptHydrateImage)
//		if err := hiScript.Write(w); err != nil {
//			return err
//		}
//		ifiScript := script.ScriptAsync(scriptImageFadeIn)
//		return ifiScript.Write(w)
//	}
//	return nil
//}

func issaImage(r compton.Registrar, placeholder, poster string, dehydrated bool) compton.Element {
	ii := &IssaImageElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.IssaImage,
			Markup:  markupIssaImage,
		},
		//r:          r,
		dehydrated: dehydrated,
	}

	placeholderImg := els.Image("")
	classes := []string{"placeholder"}

	if dehydrated {
		placeholderImg.SetAttribute("data-dehydrated", placeholder)
		classes = append(classes, "loading")
	} else {
		placeholderImg.SetAttribute("src", placeholder)
	}

	placeholderImg.AddClass(classes...)

	posterImg := els.ImageLazy(poster)
	posterImg.AddClass("poster", "loading")
	ii.Append(placeholderImg, posterImg)

	r.RegisterStyle(styleRegistrationName, styleIssaImage)
	r.RegisterDeferral(scriptHydrateImageRegistrationName, script.ScriptAsync(scriptHydrateImage))
	r.RegisterDeferral(scriptImageFadeInRegistrationName, script.ScriptAsync(scriptImageFadeIn))
	r.RegisterDeferral(scriptHydrateColorRegistrationName, script.ScriptAsync(issa.HydrateColorScript))

	return ii
}

func IssaImageHydrated(r compton.Registrar, placeholder, poster string) compton.Element {
	return issaImage(r, placeholder, poster, false)
}

func IssaImageDehydrated(r compton.Registrar, placeholder, poster string) compton.Element {
	return issaImage(r, placeholder, poster, true)
}
