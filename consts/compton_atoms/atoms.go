package compton_atoms

import "golang.org/x/net/html/atom"

const (
	DetailsClosed atom.Atom = 0xffffff00 + iota
	DetailsOpen
	FlexItems
	GridItems
	TitleValues
	NavLinks
	SvgInlineIcon
	SectionHighlight
)
