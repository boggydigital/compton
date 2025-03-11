package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

var (
	//go:embed "script/one_liner.js"
	scriptOneLiner []byte
)

type OneLinerElement struct {
	*BaseElement
}

func OneLiner(r Registrar) Element {

	ol := &OneLinerElement{NewElement(tacMarkup(compton_atoms.OneLiner))}

	r.RegisterStyles(DefaultStyle, compton_atoms.StyleName(compton_atoms.OneLiner))
	r.RegisterDeferrals(compton_atoms.ScriptName(compton_atoms.OneLiner),
		ScriptAsync(scriptOneLiner))

	return ol
}
