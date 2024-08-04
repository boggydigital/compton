package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/heading"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/stack"
	"github.com/boggydigital/compton/table"
	"os"
	"path/filepath"
)

//go:embed "styles.css"
var appStyles []byte

func main() {

	p := page.New("test", "🤔")
	p.SetCustomStyles(appStyles)

	s := stack.New(p, stack.Large)

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
