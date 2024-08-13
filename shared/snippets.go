package shared

import (
	_ "embed"
)

var (
	//go:embed "styles/host-background-color.css"
	StyleHostBackgroundColors []byte
	//go:embed "styles/host-foreground-color.css"
	StyleHostForegroundColors []byte
)
