package compton_atoms

import (
	"golang.org/x/net/html/atom"
	"path"
)

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
	IframeExpandHost
	IframeExpandContent
	Popup
)

var atomStrings = map[atom.Atom]string{
	Page:                "page",
	Document:            "document",
	Doctype:             "doctype",
	Requirements:        "requirements",
	Content:             "content",
	Deferrals:           "deferrals",
	DetailsSummary:      "details-summary",
	FlexItems:           "flex-items",
	GridItems:           "grid-items",
	TitleValues:         "title-values",
	NavLinks:            "nav-links",
	IssaImage:           "issa-image",
	Fspan:               "fspan",
	Labels:              "labels",
	IframeExpandHost:    "iframe-expand-host",
	IframeExpandContent: "iframe-expand-content",
	Popup:               "popup",
}

func String(ca atom.Atom) string {
	if str, ok := atomStrings[ca]; ok {
		return str
	}
	panic("no string for atom")
}

func MarkupName(ca atom.Atom) string {
	return path.Join("markup", String(ca)) + ".html"
}

func StyleName(ca atom.Atom) string {
	return path.Join("style", String(ca)) + ".css"
}
