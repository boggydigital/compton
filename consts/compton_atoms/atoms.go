package compton_atoms

import "golang.org/x/net/html/atom"

const (
	DetailsClosed atom.Atom = 0xfffff000 + iota
	DetailsOpen
	FlexItems
	GridItems
	TitleValues
	NavLinks
	Section
	IssaImage
	TextLine
)
