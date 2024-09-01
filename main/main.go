package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/alignment"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/details-toggle"
	els2 "github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/flex-items"
	"github.com/boggydigital/compton/elements/grid-items"
	iframe_expand2 "github.com/boggydigital/compton/elements/iframe-expand"
	nav_links2 "github.com/boggydigital/compton/elements/nav-links"
	"github.com/boggydigital/compton/elements/page"
	"github.com/boggydigital/compton/elements/section-highlight"
	svg_inline "github.com/boggydigital/compton/elements/svg-inline"
	title_values2 "github.com/boggydigital/compton/elements/title-values"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//go:embed "styles.css"
var appStyles []byte

func main() {
	writeTestPage()
}

func writeTestPage() {
	p := page.New("test").SetFavIconEmoji("🤔")
	p.SetCustomStyles(appStyles)

	s := flex_items.New(p, direction.Column)

	topNavLinks := map[string]string{
		"Updates": "/updates",
		"Search":  "/search",
	}

	topNavIcons := map[string]svg_inline.Symbol{
		"Updates": svg_inline.Sparkle,
		"Search":  svg_inline.Search,
	}

	targets := nav_links2.TextLinks(
		topNavLinks,
		"Search",
		"Updates", "Search")
	nav_links2.SetIcons(targets, topNavIcons)

	topNav := nav_links2.NewLinks(p, targets...)

	s.Append(topNav)

	navLinks := map[string]string{
		"New":      "/new",
		"Owned":    "/owned",
		"Wishlist": "/wishlist",
		"Sale":     "/sale",
		"All":      "/all",
	}

	nav := nav_links2.NewLinks(p,
		nav_links2.TextLinks(
			navLinks,
			"New",
			"New",
			"Owned",
			"Wishlist",
			"Sale",
			"All")...)

	s.Append(nav)

	cdc := details_toggle.NewOpen(p, "Filter & Search")

	form := els2.NewForm("/action", "GET")

	formStack := flex_items.New(p, direction.Column)

	qf := createQueryFragment(p)

	formStack.Append(qf)

	submitRow := flex_items.New(p, direction.Row).
		JustifyContent(alignment.Center)

	submit := els2.NewInputValue(input_types.Submit, "Submit Query")
	submitRow.Append(submit)
	formStack.Append(submitRow)

	tiGrid := grid_items.New(p)

	ti1 := title_values2.NewSearchValue(p, "Title", "title", "Hello")

	tiList := map[string]string{
		"true":  "True",
		"false": "False",
		"maybe": "Maybe",
	}

	ti2 := title_values2.NewSearch(p, "Description", "description").
		SetDataList(tiList)
	tiGrid.Append(ti1, ti2)

	formStack.Append(tiGrid)
	form.Append(formStack)

	cdc.Append(form)
	s.Append(cdc)

	s.Append(qf)

	cdo := details_toggle.NewOpen(p, "Title Values")

	tvGrid := grid_items.New(p)

	tvLinks := map[string]string{
		"Achievements":       "/achievements",
		"Controller support": "/controller-support",
		"Overlay":            "/overlay",
		"Single-player":      "/single-player",
	}
	tv1 := title_values2.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv2 := title_values2.NewLinks(p, "Feature Links", tvLinks)
	tv3 := title_values2.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv4 := title_values2.NewLinks(p, "Feature Links", tvLinks)
	tv5 := title_values2.NewText(p, "Features", maps.Keys(tvLinks)...)
	tv6 := title_values2.NewLinks(p, "Feature Links", tvLinks)

	tvGrid.Append(tv1, tv2, tv3, tv4, tv5, tv6)
	cdo.Append(tvGrid)
	s.Append(cdo)

	footer := flex_items.New(p, direction.Row).
		JustifyContent(alignment.Center)

	div := els2.NewDiv()
	div.SetClass("fg-subtle", "fs-xs")

	div.Append(els2.NewText("Last updated: "),
		els2.NewTimeText(time.Now().Format("2006-01-02 15:04:05")))

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

func writeIframeContent() {

	c := page.New("content")
	ifec := iframe_expand2.NewContent("test", "whatever")
	c.Append(ifec)

	for i := range 1000 {
		c.Append(els2.NewDivText(strconv.Itoa(i)))
	}

	tempPath := filepath.Join(os.TempDir(), "content.html")
	tempFile, err := os.Create(tempPath)
	if err != nil {
		panic(err)
	}

	if err := c.WriteContent(tempFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + tempPath)

	p := page.New("iframe")

	dc := details_toggle.NewClosed(p, "Description")

	ife := iframe_expand2.New(p, "test", "content.html")
	dc.Append(ife)

	p.Append(dc)

	tempPath = filepath.Join(os.TempDir(), "test.html")
	tempFile, err = os.Create(tempPath)
	if err != nil {
		panic(err)
	}

	if err := p.WriteContent(tempFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + tempPath)
}

func createQueryFragment(r compton.Registrar) compton.Element {
	sh := section_highlight.New(r)
	sh.SetClass("fs-xs")

	shStack := flex_items.New(r, direction.Row).SetColumnGap(size.Normal)
	sh.Append(shStack)

	sp1 := els2.NewSpan()
	pt1 := els2.NewSpanText("Descending: ")
	pt1.SetClass("fg-subtle")
	pv1 := els2.NewSpanText("True")
	pv1.SetClass("fw-b")
	sp1.Append(pt1, pv1)
	shStack.Append(sp1)

	sp2 := els2.NewSpan()
	pt2 := els2.NewSpanText("Sort: ")
	pt2.SetClass("fg-subtle")
	pv2 := els2.NewSpanText("GOG Order Date")
	pv2.SetClass("fw-b")
	sp2.Append(pt2, pv2)
	shStack.Append(sp2)

	sp3 := els2.NewSpan()
	pt3 := els2.NewSpanText("Data Type: ")
	pt3.SetClass("fg-subtle")
	pv3 := els2.NewSpanText("Account Products")
	pv3.SetClass("fw-b")
	sp3.Append(pt3, pv3)
	shStack.Append(sp3)

	clearAction := els2.NewAText("Clear", "/clear")
	clearAction.SetClass("action")
	shStack.Append(clearAction)

	return sh
}
