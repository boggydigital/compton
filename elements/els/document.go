package els

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

func Document() compton.Element {
	return compton.NewElement(compton_atoms.Document, nil)
}
