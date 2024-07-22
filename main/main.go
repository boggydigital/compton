package main

import (
	_ "embed"
	"fmt"
	human_resource "github.com/boggydigital/compton/human-resource"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/web-components"
	"os"
	"path/filepath"
)

var ()

func main() {

	cr := web_components.NewComponentRegistrar()

	p := page.New("test")

	p.Add(human_resource.New(cr, "John", "Smith", "Sales"))
	p.Add(human_resource.New(cr, "Mike", "Jones", "Marketing"))
	p.Add(human_resource.New(cr, "Brian", "Paul", "Security"))
	p.Add(human_resource.New(cr, "Fiona", "Apple", "Capital"))

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
