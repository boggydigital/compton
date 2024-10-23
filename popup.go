package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

const (
	rnPopup = "popup"
)

var (
	//go:embed "script/popup.js"
	scriptPopup []byte
)

type PopupElement struct {
	*BaseElement
}

func Attach(r Registrar, actor, target Element) *PopupElement {
	pe := &PopupElement{
		BaseElement: NewElement(contentMarkup(compton_atoms.Popup)),
	}

	if targetId := target.GetAttribute("id"); targetId != "" {
		actor.SetAttribute("data-popup-target", targetId)
		target.SetAttribute("data-popup", "hide")
	}

	r.RegisterStyles(comptonAtomStyle,
		compton_atoms.StyleName(compton_atoms.Popup))
	r.RegisterDeferrals(compton_atoms.ScriptName(compton_atoms.Popup), ScriptAsync(scriptPopup))

	return pe
}
