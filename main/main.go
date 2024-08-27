package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/alignment"
	details_toggle "github.com/boggydigital/compton/details-toggle"
	"github.com/boggydigital/compton/direction"
	"github.com/boggydigital/compton/els"
	flex_items "github.com/boggydigital/compton/flex-items"
	grid_items "github.com/boggydigital/compton/grid-items"
	"github.com/boggydigital/compton/input_types"
	nav_links "github.com/boggydigital/compton/nav-links"
	"github.com/boggydigital/compton/page"
	section_highlight "github.com/boggydigital/compton/section-highlight"
	"github.com/boggydigital/compton/size"
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

	s := flex_items.New(p, direction.Column)

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
		"Search",
		"Updates", "Search")
	nav_links.SetIcons(targets, topNavIcons)

	topNav := nav_links.NewLinks(p, targets...)

	s.Append(topNav)

	navLinks := map[string]string{
		"New":      "/new",
		"Owned":    "/owned",
		"Wishlist": "/wishlist",
		"Sale":     "/sale",
		"All":      "/all",
	}

	nav := nav_links.NewLinks(p,
		nav_links.TextLinks(
			navLinks,
			"New",
			"New",
			"Owned",
			"Wishlist",
			"Sale",
			"All")...)

	s.Append(nav)

	cdc := details_toggle.NewOpen(p, "Filter & Search")

	form := els.NewForm("/action", "GET")

	formStack := flex_items.New(p, direction.Column)

	qf := createQueryFragment(p)

	formStack.Append(qf)

	submitRow := flex_items.New(p, direction.Row).
		JustifyContent(alignment.Center)

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

	s.Append(qf)

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

	footer := flex_items.New(p, direction.Row).
		JustifyContent(alignment.Center)

	div := els.NewDiv()
	div.SetClass("fg-subtle", "fs-xs")

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

func createQueryFragment(r compton.Registrar) compton.Element {
	sh := section_highlight.New(r)
	sh.SetClass("fs-xs")

	shStack := flex_items.New(r, direction.Row).SetColumnGap(size.Normal)
	sh.Append(shStack)

	sp1 := els.NewSpan()
	pt1 := els.NewSpanText("Descending:")
	pt1.SetClass("fg-subtle")
	pv1 := els.NewSpanText("True")
	pv1.SetClass("fw-b")
	sp1.Append(pt1, pv1)
	shStack.Append(sp1)

	sp2 := els.NewSpan()
	pt2 := els.NewSpanText("Sort:")
	pt2.SetClass("fg-subtle")
	pv2 := els.NewSpanText("GOG Order Date")
	pv2.SetClass("fw-b")
	sp2.Append(pt2, pv2)
	shStack.Append(sp2)

	sp3 := els.NewSpan()
	pt3 := els.NewSpanText("Data Type:")
	pt3.SetClass("fg-subtle")
	pv3 := els.NewSpanText("Account Products")
	pv3.SetClass("fw-b")
	sp3.Append(pt3, pv3)
	shStack.Append(sp3)

	clearAction := els.NewAText("Clear", "/clear")
	clearAction.SetClass("action")
	shStack.Append(clearAction)

	return sh
}
