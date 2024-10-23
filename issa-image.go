package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/issa"
)

const (
	rnIssaImage                        = "issa-image"
	scriptImageFadeInRegistrationName  = "script-image-fade-in-" + rnIssaImage
	scriptHydrateColorRegistrationName = "script-hydrate-color-" + rnIssaImage
	scriptHydrateImageRegistrationName = "script-hydrate-image" + rnIssaImage
)

var (
	//go:embed "script/image_fadein.js"
	scriptImageFadeIn []byte
	//go:embed "script/hydrate_images.js"
	scriptHydrateImage []byte
)

type IssaImageElement struct {
	*BaseElement
	dehydrated bool
}

func issaImage(r Registrar, placeholder, poster string, dehydrated bool) Element {
	ii := &IssaImageElement{
		BaseElement: NewElement(tacMarkup(compton_atoms.IssaImage)),
		dehydrated:  dehydrated,
	}

	placeholderImg := Image("")
	classes := []string{"placeholder"}

	if dehydrated {
		placeholderImg.SetAttribute("data-dehydrated", placeholder)
		classes = append(classes, "loading")
	} else {
		placeholderImg.SetAttribute("src", placeholder)
	}

	placeholderImg.AddClass(classes...)

	posterImg := ImageLazy(poster)
	posterImg.AddClass("poster", "loading")
	ii.Append(placeholderImg, posterImg)

	r.RegisterStyles(comptonAtomStyle,
		compton_atoms.StyleName(compton_atoms.IssaImage))
	r.RegisterDeferrals(compton_atoms.ScriptName(compton_atoms.IssaImage),
		ScriptAsync(scriptHydrateImage),
		ScriptAsync(scriptImageFadeIn),
		ScriptAsync(issa.HydrateColorScript))

	return ii
}

func IssaImageHydrated(r Registrar, placeholder, poster string) Element {
	return issaImage(r, placeholder, poster, false)
}

func IssaImageDehydrated(r Registrar, placeholder, poster string) Element {
	return issaImage(r, placeholder, poster, true)
}
