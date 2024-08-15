package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/anchors"
	details_toggle "github.com/boggydigital/compton/details-toggle"
	"github.com/boggydigital/compton/directions"
	"github.com/boggydigital/compton/els"
	flex_items "github.com/boggydigital/compton/flex-items"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/table"
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

	s := flex_items.New(p, directions.Column).SetRowGap(measures.Large)

	h1 := els.NewHeadingText("Success", 1)
	h1.SetClass("success")
	s.Append(h1)

	t := table.New().
		AppendHead("Property", "Value", "Another one").
		AppendRow("Name", "John", "two").
		AppendRow("Last Name", "Smith", "three").
		AppendFoot("Summary", "123", "456")
	t.SetClass("red")
	s.Append(t)

	cdo := details_toggle.NewOpen(p, "Description").
		SetSummaryMargin(measures.Large)
	//SetBackgroundColor(colors.Red).
	//SetForegroundColor(colors.Background)

	nso := flex_items.New(p, directions.Row)
	//JustifyContent(anchors.Center).
	nso.Append(els.NewAText("One", "/one"), els.NewAText("Two", "/two"))

	cdo.Append(nso)
	s.Append(cdo)

	cdc := details_toggle.NewClosed(p, "Screenshots").
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

	links := map[string]string{
		"Achievements":       "/achievements",
		"Controller support": "/controller-support",
		"Overlay":            "/overlay",
		"Single-player":      "/single-player",
	}
	tv := title_values.NewText(p, "Features", maps.Keys(links)...)

	s.Append(tv)

	p.Append(s)

	//p.Append(tv)

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
