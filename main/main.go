package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/anchor"
	c_details "github.com/boggydigital/compton/c-details"
	c_stack "github.com/boggydigital/compton/c-stack"
	"github.com/boggydigital/compton/heading"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/table"
	"os"
	"path/filepath"
)

//go:embed "styles.css"
var appStyles []byte

func main() {

	p := page.New("test", "🤔")
	p.SetCustomStyles(appStyles)

	s := c_stack.New(p).SetRowGap(measures.Large)

	h1 := heading.NewText("Success", 1)
	h1.SetClass("success")
	h1.SetAttr("data-test", "test-val")
	s.Append(h1)

	t := table.New().
		AppendHead("Property", "Value", "Another one").
		AppendRow("Name", "John", "two").
		AppendRow("Last Name", "Smith", "three").
		AppendFoot("Summary", "123", "456")
	s.Append(t)

	cdo := c_details.New(p, "Open").SetSummaryMarginBlockEnd(measures.Large).Open()

	nso := c_stack.New(p)
	nso.Append(anchor.NewText("One", "/one"), anchor.NewText("Two", "/two"))
	cdo.Append(nso)
	s.Append(cdo)

	cdc := c_details.New(p, "Closed").SetSummaryMarginBlockEnd(measures.Large)

	nsc := c_stack.New(p)
	nsc.Append(anchor.NewText("One", "/one"), anchor.NewText("Two", "/two"))
	cdc.Append(nsc)
	s.Append(cdc)

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
