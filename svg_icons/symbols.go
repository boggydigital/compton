package svg_icons

import _ "embed"

type Symbol int

const (
	None Symbol = iota
	Windows
	MacOS
	Linux
	Plus
	Star
	Sparkle
	Stack
	Search
	BackToTop
)

var (
	MarkupSymbols = map[Symbol][]byte{
		Windows:   markupIconWindows,
		MacOS:     markupIconMacOS,
		Linux:     markupIconLinux,
		Plus:      markupIconPlus,
		Star:      markupIconStar,
		Sparkle:   markupIconSparkle,
		Stack:     markupIconStack,
		Search:    markupIconSearch,
		BackToTop: markupIconBackToTop,
	}

	//go:embed "markup/plus_icon.svg"
	markupIconPlus []byte
	//go:embed "markup/sparkle_icon.svg"
	markupIconSparkle []byte
	//go:embed "markup/search_icon.svg"
	markupIconSearch []byte
	//go:embed "markup/windows_icon.svg"
	markupIconWindows []byte
	//go:embed "markup/macos_icon.svg"
	markupIconMacOS []byte
	//go:embed "markup/linux.svg"
	markupIconLinux []byte
	//go:embed "markup/star.svg"
	markupIconStar []byte
	//go:embed "markup/stack.svg"
	markupIconStack []byte
	//go:embed "markup/back-to-top.svg"
	markupIconBackToTop []byte
)
