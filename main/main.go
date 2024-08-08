package main

import (
	_ "embed"
	"fmt"
	c_details "github.com/boggydigital/compton/c-details"
	c_stack "github.com/boggydigital/compton/c-stack"
	"github.com/boggydigital/compton/elements"
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

	s := c_stack.New(p).SetRowGap(measures.Large)

	s.Append(elements.NewHeadingText("Success", 1).
		SetClass("success"))

	t := table.New().
		AppendHead("Property", "Value", "Another one").
		AppendRow("Name", "John", "two").
		AppendRow("Last Name", "Smith", "three").
		AppendFoot("Summary", "123", "456").
		SetClass("red")
	s.Append(t)

	cdo := c_details.New(p, "Open").SetSummaryMarginBlockEnd(measures.Large).Open()

	nso := c_stack.New(p).
		Append(elements.NewAText("One", "/one"), elements.NewAText("Two", "/two"))

	cdo.Append(nso)
	s.Append(cdo)

	cdc := c_details.New(p, "Closed").SetSummaryMarginBlockEnd(measures.Large)

	nsc := c_stack.New(p).Append(elements.NewAText("One", "/one"), elements.NewAText("Two", "/two"))
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
