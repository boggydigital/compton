package main

import (
	_ "embed"
	"github.com/boggydigital/compton/page"
	"github.com/boggydigital/compton/text"
	"os"
)

var ()

func main() {

	p := page.New("test")
	p.Add(text.New("testing"))

	if err := p.Write(os.Stdout); err != nil {
		panic(err)
	}
}
