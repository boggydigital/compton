package compton

import (
	_ "embed"

	"github.com/boggydigital/compton/consts/compton_atoms"
)

var (
	//go:embed "script/copy_to_clipboard.js"
	scriptCopyToClipboard []byte
)

func CopyToClipboard(r Registrar, cta, success, err Element, value string) Element {

	r.RegisterDeferrals(compton_atoms.ScriptName(compton_atoms.CopyToClipboard),
		ScriptAsync(scriptCopyToClipboard))

	copyToClipboard := NewElement(tacMarkup(compton_atoms.CopyToClipboard))
	cta.AddClass("copy-to-clipboard-cta")

	copyToClipboard.Append(cta, success, err)
	copyToClipboard.SetAttribute("style", "display:flex; cursor:pointer")

	success.SetAttribute("style", "display:none")
	success.AddClass("copy-to-clipboard-success")

	err.SetAttribute("style", "display:none")
	err.AddClass("copy-to-clipboard-error")

	copyToClipboard.SetAttribute("data-value", value)

	return copyToClipboard
}
