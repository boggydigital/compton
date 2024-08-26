package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton/anchors"
	details_toggle "github.com/boggydigital/compton/details-toggle"
	"github.com/boggydigital/compton/directions"
	"github.com/boggydigital/compton/els"
	flex_items "github.com/boggydigital/compton/flex-items"
	grid_items "github.com/boggydigital/compton/grid-items"
	"github.com/boggydigital/compton/input_types"
	nav_links "github.com/boggydigital/compton/nav-links"
	"github.com/boggydigital/compton/page"
	section_highlight "github.com/boggydigital/compton/section-highlight"
	"github.com/boggydigital/compton/svg_inline"
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

	s := flex_items.New(p, directions.Column)

	topNavLinks := map[string]string{
		"Updates": "/updates",
		"Search":  "/search",
	}

	topNavIcons := map[string]svg_inline.Symbol{
		"Updates": svg_inline.Sparkle,
		"Search":  svg_inline.Search,
	}

	targets := nav_links.TextLinks(
		topNavLinks,
		"Updates",
		"Updates", "Search")
	nav_links.SetIcons(targets, topNavIcons)

	topNav := nav_links.NewLinks(p, targets...)

	s.Append(topNav)

	//h1 := els.NewHeadingText("Success", 1)
	//h1.SetClass("success")
	//s.Append(h1)
	//
	//t := table.New().
	//	AppendHead("Property", "Value", "Another one").
	//	AppendRow("Name", "John", "two").
	//	AppendRow("Last Name", "Smith", "three").
	//	AppendFoot("Summary", "123", "456")
	//t.SetClass("red")
	//s.Append(t)

	navLinks := map[string]string{
		"Description":   "#description",
		"Screenshots":   "#screenshots",
		"Videos":        "#videos",
		"Steam News":    "#steam_news",
		"Steam Reviews": "#steam_reviews",
		"Steam Deck":    "#steam_deck",
		"Downloads":     "#download",
	}

	nav := nav_links.NewLinks(p,
		nav_links.TextLinks(
			navLinks,
			"",
			"Description",
			"Screenshots",
			"Videos",
			"Steam News",
			"Steam Reviews",
			"Steam Deck",
			"Downloads")...)

	s.Append(nav)

	cdc := details_toggle.NewOpen(p, "Title Inputs")

	form := els.NewForm("/action", "GET")

	formStack := flex_items.New(p, directions.Column)

	sh := section_highlight.New(p)
	sh.SetClass("fs-x-smaller")

	clearAction := els.NewAText("Clear", "/clear")
	clearAction.SetClass("action")
	sh.Append(clearAction)

	formStack.Append(sh)

	submitRow := flex_items.New(p, directions.Row).
		JustifyContent(anchors.Center)

	submit := els.NewInputValue(input_types.Submit, "Submit Query")
	submitRow.Append(submit)
	formStack.Append(submitRow)

	tiGrid := grid_items.New(p)

	ti1 := title_values.NewSearchValue(p, "Title", "title", "Hello")

	tiList := map[string]string{
		"true":  "True",
		"false": "False",
		"maybe": "Maybe",
	}

	ti2 := title_values.NewSearch(p, "Description", "description").
		SetDataList(tiList)
	tiGrid.Append(ti1, ti2)

	formStack.Append(tiGrid)
	form.Append(formStack)

	cdc.Append(form)
	s.Append(cdc)

	cdo := details_toggle.NewOpen(p, "Title Values")

	tvGrid := grid_items.New(p)

	tvLinks := map[string]string{
		"Achievements":       "/achievements",
		"Controller support": "/controller-support",
		"Overlay":            "/overlay",
		"Single-player":      "/single-player",
	}
	tv1 := title_values.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv2 := title_values.NewLinks(p, "Feature Links", tvLinks)
	tv3 := title_values.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv4 := title_values.NewLinks(p, "Feature Links", tvLinks)
	tv5 := title_values.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv6 := title_values.NewLinks(p, "Feature Links", tvLinks)

	tvGrid.Append(tv1, tv2, tv3, tv4, tv5, tv6)
	cdo.Append(tvGrid)
	s.Append(cdo)

	footer := flex_items.New(p, directions.Row).
		JustifyContent(anchors.Center)

	div := els.NewDiv()
	div.SetClass("fg-subtle", "fs-x-smaller")

	div.Append(els.NewText("Last updated: "),
		els.NewTimeText(time.Now().Format("2006-01-02 15:04:05")))

	footer.Append(div)

	s.Append(footer)

	p.Append(s)

	tempPath := filepath.Join(os.TempDir(), "test.html")
	tempFile, err := os.Create(tempPath)
	if err != nil {
		panic(err)
	}

	if err := p.WriteContent(tempFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + tempPath)
}
