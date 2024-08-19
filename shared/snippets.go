package shared

import (
	_ "embed"
)

var (
	//go:embed "styles/host-background-color.css"
	StyleHostBackgroundColor []byte
	//go:embed "styles/host-foreground-color.css"
	StyleHostForegroundColor []byte
	//go:embed "styles/host-row-gap.css"
	StyleHostRowGap []byte
	//go:embed "styles/host-column-gap.css"
	StyleHostColumnGap []byte
	//go:embed "styles/host-summary-margin.css"
	StyleHostSummaryMargin []byte
	//go:embed "styles/host-details-margin.css"
	StyleHostDetailsMargin []byte
	//go:embed "styles/host-align-content.css"
	StyleHostAlignContent []byte
	//go:embed "styles/host-justify-content.css"
	StyleHostJustifyContent []byte
	//go:embed "styles/host-flex-direction.css"
	StyleHostFlexDirection []byte
)
