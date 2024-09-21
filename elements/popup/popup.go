package popup

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "popup"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "style/popup.css"
	stylePopup []byte
	//go:embed "script/popup.js"
	scriptPopup []byte
)

type PopupElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (pe *PopupElement) WriteStyles(w io.Writer) error {
	if pe.r.RequiresRegistration(styleRegistrationName) {
		return els.Style(stylePopup, styleRegistrationName).WriteContent(w)
	}
	return nil
}

func (pe *PopupElement) WriteDeferrals(w io.Writer) error {
	if pe.r.RequiresRegistration(registrationName) {
		return els.ScriptAsync(scriptPopup).WriteContent(w)
	}
	return nil
}

func Attach(r compton.Registrar, actor, target compton.Element) *PopupElement {
	pe := &PopupElement{
		BaseElement: compton.BaseElement{},
		r:           r,
	}

	if targetId := target.GetAttribute("id"); targetId != "" {
		actor.SetAttribute("data-popup-target", targetId)
		target.SetAttribute("data-popup", "hide")
	}

	return pe
}
