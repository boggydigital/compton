package compton_atoms

import "golang.org/x/net/html/atom"

const (
	Page atom.Atom = 0xfffff000 + iota
	Document
	Doctype
	Requirements
	Content
	Deferrals
	DetailsSummary
	FlexItems
	GridItems
	TitleValues
	NavLinks
	IssaImage
	Fspan
	Labels
)
