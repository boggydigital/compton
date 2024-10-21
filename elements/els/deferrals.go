package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

func Deferrals() compton.Element {
	return compton.NewElement(compton_atoms.Deferrals, nil)
}
