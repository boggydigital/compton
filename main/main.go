package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/heading"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/table"
	"os"
	"path/filepath"
)

func main() {

	p := page.New("test", "🤔")

	p.Append(heading.NewText("Heading 1", 1, ""))

	p.Append(
		table.New("test-table").
			AppendHead("Property", "Value", "Another one").
			AppendRow("Name", "John", "two").
			AppendRow("Last Name", "Smith", "three"))

	//thead := table.NewHead().Append(
	//	table.NewTh().Append(text.New("Property")),
	//	table.NewTh().Append(text.New("Value")),
	//)

	//tbody := table.NewBody().Append(
	//	table.NewTr().Append(
	//		table.NewTd().Append(text.New("Name")),
	//		table.NewTd().Append(text.New("John")),
	//	),
	//	table.NewTr().Append(
	//		table.NewTd().Append(text.New("Last Name")),
	//		table.NewTd().Append(text.New("Smith")),
	//	),
	//)
	//t.Append(tbody)

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
