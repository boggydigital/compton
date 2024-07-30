package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/table"
	"os"
	"path/filepath"
)

func main() {

	p := page.New("test")

	t := table.New(p, "")
	t.AppendHead("Property", "Value")
	t.AppendRow("Name", "John")
	t.AppendRow("Last Name", "Smith")

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

	p.Append(t)

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
