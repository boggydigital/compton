package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/anchors"
	"github.com/boggydigital/compton/colors"
	details_toggle "github.com/boggydigital/compton/details-toggle"
	"github.com/boggydigital/compton/directions"
	"github.com/boggydigital/compton/els"
	flex_items "github.com/boggydigital/compton/flex-items"
	grid_items "github.com/boggydigital/compton/grid-items"
	"github.com/boggydigital/compton/measures"
	nav_links "github.com/boggydigital/compton/nav-links"
	"github.com/boggydigital/compton/page"
	title_values "github.com/boggydigital/compton/title-values"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
	"time"
)

//go:embed "styles.css"
var appStyles []byte

func main() {

	p := page.New("test", "🤔")
	p.SetCustomStyles(appStyles)

	s := flex_items.New(p, directions.Column).
		SetRowGap(measures.Large)

	//h1 := els.NewHeadingText("Success", 1)
	//h1.SetClass("success")
	//s.Append(h1)
	//
	//t := table.New().
	//	AppendHead("Property", "Value", "Another one").
	//	AppendRow("Name", "John", "two").
	//	AppendRow("Last Name", "Smith", "three").
	//	AppendFoot("Summary", "123", "456")
	//t.SetClass("red")
	//s.Append(t)

	navLinks := map[string]string{
		"Description":   "#description",
		"Screenshots":   "#screenshots",
		"Videos":        "#videos",
		"Steam News":    "#steam_news",
		"Steam Reviews": "#steam_reviews",
		"Steam Deck":    "#steam_deck",
		"Downloads":     "#downloads",
	}

	nav := nav_links.NewLinks(p, navLinks)

	s.Append(nav)

	cdo := details_toggle.NewOpen(p, "Description").
		SetSummaryMargin(measures.XLarge).
		SetDetailsMargin(measures.Large).
		SetBackgroundColor(colors.LightBlue).
		SetForegroundColor(colors.Background)

	gridItems := grid_items.New(p).
		SetRowGap(measures.Large).
		SetColumnGap(measures.Large)
	//AlignContent(anchors.Center)
	//nso.Append(els.NewAText("One", "/one"), els.NewAText("Two", "/two"))

	tvLinks := map[string]string{
		"Achievements":       "/achievements",
		"Controller support": "/controller-support",
		"Overlay":            "/overlay",
		"Single-player":      "/single-player",
	}
	tv1 := title_values.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv2 := title_values.NewLinks(p, "Feature Links", tvLinks)
	tv3 := title_values.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv4 := title_values.NewLinks(p, "Feature Links", tvLinks)
	tv5 := title_values.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv6 := title_values.NewLinks(p, "Feature Links", tvLinks)

	gridItems.Append(tv1, tv2, tv3, tv4, tv5, tv6)
	cdo.Append(gridItems)
	s.Append(cdo)

	cdc := details_toggle.NewClosed(p, "Steam Deck").
		SetSummaryMargin(measures.Large)

	nsc := flex_items.New(p, directions.Column)
	//AlignContent(anchors.Center).
	nsc.Append(els.NewAText("One", "/one"), els.NewAText("Two", "/two"))
	cdc.Append(nsc)
	s.Append(cdc)

	footer := flex_items.New(p, directions.Row).
		JustifyContent(anchors.Center)
	div := els.NewDiv()
	div.SetClass("subtle")

	div.Append(els.NewText("Last updated: "),
		els.NewTimeText(time.Now().Format("2006-01-02 15:04:05")))

	footer.Append(div)

	s.Append(footer)

	p.Append(s)

	tempPath := filepath.Join(os.TempDir(), "test.html")
	tempFile, err := os.Create(tempPath)
	if err != nil {
		panic(err)
	}

	if err := p.Write(tempFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + tempPath)
}
