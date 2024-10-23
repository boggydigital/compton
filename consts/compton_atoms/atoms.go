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
	SvgUse
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
	SvgUse:              "svg-use",
}

func Atos(a atom.Atom) string {
	if str, ok := atomStrings[a]; ok {
		return str
	} else if an := a.String(); an != "" {
		return an
	}
	panic("no string for atom")
}

func MarkupName(a atom.Atom) string {
	return path.Join("markup", Atos(a)+".html")
}

func StyleName(a atom.Atom) string {
	return path.Join("style", Atos(a)+".css")
}

func ScriptName(a atom.Atom) string {
	return path.Join("script", Atos(a)+".js")
}
