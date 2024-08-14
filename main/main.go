package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/colors"
	details_toggle "github.com/boggydigital/compton/details-toggle"
	"github.com/boggydigital/compton/elements"
	"github.com/boggydigital/compton/flex"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/table"
	"os"
	"path/filepath"
	"time"
)

//go:embed "styles.css"
var appStyles []byte

func main() {

	p := page.New("test", "🤔")
	p.SetCustomStyles(appStyles)

	s := flex.NewColumn(p).SetRowGap(measures.Large)

	s.Append(elements.NewHeadingText("Success", 1).
		SetClass("success"))

	t := table.New().
		AppendHead("Property", "Value", "Another one").
		AppendRow("Name", "John", "two").
		AppendRow("Last Name", "Smith", "three").
		AppendFoot("Summary", "123", "456").
		SetClass("red")
	s.Append(t)

	cdo := details_toggle.New(p, "Open").
		SetSummaryMargin(measures.Large).
		SetBackgroundColor(colors.Red).
		SetForegroundColor(colors.Background).
		Open()

	nso := flex.NewRow(p).
		SetColumnGap(measures.Normal)

	nso.Append(elements.NewAText("One", "/one"), elements.NewAText("Two", "/two"))

	cdo.Append(nso)
	s.Append(cdo)

	cdc := details_toggle.New(p, "Closed").SetSummaryMargin(measures.Large)

	nsc := flex.NewColumn(p).
		Append(elements.NewAText("One", "/one"), elements.NewAText("Two", "/two"))
	cdc.Append(nsc)
	s.Append(cdc)

	dv := elements.NewDiv().
		SetClass("subtle").
		Append(
			elements.NewText("Last updated: "),
			elements.NewTimeText(time.Now().Format("2006-01-02 15:04:05")))
	s.Append(dv)

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
