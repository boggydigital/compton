package shared

import (
	_ "embed"
)

var (
	//go:embed "styles/host-background-color.css"
	StyleHostBackgroundColor []byte
	//go:embed "styles/host-foreground-color.css"
	StyleHostForegroundColor []byte
	//go:embed "styles/host-summary-margin.css"
	StyleHostSummaryMargin []byte
	//go:embed "styles/host-details-margin.css"
	StyleHostDetailsMargin []byte
)
