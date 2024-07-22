package main

import (
	_ "embed"
	"fmt"
	human_resource "github.com/boggydigital/compton/human-resource"
	"github.com/boggydigital/compton/page"
	"os"
	"path/filepath"
)

var ()

func main() {

	p := page.New("test")

	p.Add(human_resource.New(p, "John", "Smith", "Sales"))
	p.Add(human_resource.New(p, "Mike", "Jones", "Marketing"))
	p.Add(human_resource.New(p, "Brian", "Paul", "Security"))
	p.Add(human_resource.New(p, "Fiona", "Apple", "Capital"))

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
