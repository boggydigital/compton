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
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/flex-items"
	"github.com/boggydigital/compton/elements/grid-items"
	iframe_expand "github.com/boggydigital/compton/elements/iframe-expand"
	issa_image "github.com/boggydigital/compton/elements/issa-image"
	nav_links "github.com/boggydigital/compton/elements/nav-links"
	"github.com/boggydigital/compton/elements/page"
	"github.com/boggydigital/compton/elements/section-highlight"
	svg_inline "github.com/boggydigital/compton/elements/svg-inline"
	title_values "github.com/boggydigital/compton/elements/title-values"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//go:embed "styles.css"
var appStyles []byte

func main() {
	//writeTestPage()
	//writeIframeContent()
	writeIssaPage()
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

	testPath := filepath.Join(os.TempDir(), "test.html")
	testFile, err := os.Create(testPath)
	if err != nil {
		panic(err)
	}
	defer testFile.Close()

	if err := p.WriteContent(testFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + testPath)
}

func writeIframeContent() {

	c := page.New("content")
	ifec := iframe_expand.NewContent("test", "whatever")
	c.Append(ifec)

	for i := range 1000 {
		c.Append(els.NewDivText(strconv.Itoa(i)))
	}

	contentPath := filepath.Join(os.TempDir(), "content.html")
	contentFile, err := os.Create(contentPath)
	if err != nil {
		panic(err)
	}
	defer contentFile.Close()

	if err := c.WriteContent(contentFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + contentPath)

	p := page.New("iframe")

	dc := details_toggle.NewClosed(p, "Description")

	ife := iframe_expand.New(p, "test", "content.html")
	dc.Append(ife)

	p.Append(dc)

	iframePath := filepath.Join(os.TempDir(), "iframe.html")
	iframeFile, err := os.Create(iframePath)
	if err != nil {
		panic(err)
	}

	if err := p.WriteContent(iframeFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + iframePath)
}

func writeIssaPage() {
	//hydratedSrc := "data:image/gif;base64,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"
	//imageSrc := "https://gaugin.frmnt.io/image?id=eaad5d1fce93e36d33b9983fba7edc623e23793e138667aafff6b7a305717c84"

	dehydratedSrc := "21=30=I/wABCBTYZmCkFUkEJhwI4ABDSZKStElia46tSJIutYlUgiFBXpLa2LIlKdKcSxVN2uooMABFSRdHloyEcQ7GSAMDBICI8pKkJxAh2rpEFCeAFS5tgTsgwwodcFVgPnlS4slJSXPUBVghj45QWytKpLvkhufIOSRWHAgQaaS8KhZBSpKHydaTSAsBOGhwo+5TSew+zEkX44mtNm0CDNQAAgRMeeZspUsHLgPCwU8OOAQw4UOGS5E+fDA3B9yEJ0lipJuDVDEJBKop32ARzgEFhElKmJjDooTiFQ1izJlDp40SEyYElEiSJIAJdUpKAG/gIMOE4ElM5A5AYjuLPSXCr/9lAUJ0hgwHSggIgKS9gHDgwwc44KBxY+vrS7QnEQB+eBOa1WdfBg0EoB9zJRyQQ3wsNLCCA4zdtx4JSLSRjmYIuKNEEiwYAEAAEYKQAQK5LbcCOOqEsw4UVQnUQIR8TaAZCyxAQc8eV7jDom8NUQCCBiZURqB0JkBhQjiC7KhWAPUBeQOBOgWAwAE4iNGOOlCwEBaTIFBQwpMTRDnBBF3ocZwJ6Rh5AAUfeAlOAw3QiKE8x+kExRNK0AcCAtKtQGNrJtBhQpQBPAEFAA2MmCCNuiFnwhMG1tlnFTEY6KBvjkKBlBJKRNHCoMuRAAAJT1RqoE4AOucociyNmlqUBqwVtRYChALQETa2lqDCWgB4mJNOAwUEADs"
	//hydratedSrc := "data:image/gif;base64,R0lGODlhFQAeAAAAACwAAAAAFQAeAIcqKioqKlQqKn4qKqgqKtIqKvwqVCoqVFQqVH4qVKgqVNIqVPwqfioqflQqfn4qfqgqftIqfvwqqCoqqFQqqH4qqKgqqNIqqPwq0ioq0lQq0n4q0qgq0tIq0vwq/Coq/FQq/H4q/Kgq/NIq/PxUKipUKlRUKn5UKqhUKtJUKvxUVCpUVFRUVH5UVKhUVNJUVPxUfipUflRUfn5UfqhUftJUfvxUqCpUqFRUqH5UqKhUqNJUqPxU0ipU0lRU0n5U0qhU0tJU0vxU/CpU/FRU/H5U/KhU/NJU/Px+Kip+KlR+Kn5+Kqh+KtJ+Kvx+VCp+VFR+VH5+VKh+VNJ+VPx+fip+flR+fn5+fqh+ftJ+fvx+qCp+qFR+qH5+qKh+qNJ+qPx+0ip+0lR+0n5+0qh+0tJ+0vx+/Cp+/FR+/H5+/Kh+/NJ+/PyoKiqoKlSoKn6oKqioKtKoKvyoVCqoVFSoVH6oVKioVNKoVPyofiqoflSofn6ofqioftKofvyoqCqoqFSoqH6oqKioqNKoqPyo0iqo0lSo0n6o0qio0tKo0vyo/Cqo/FSo/H6o/Kio/NKo/PzSKirSKlTSKn7SKqjSKtLSKvzSVCrSVFTSVH7SVKjSVNLSVPzSfirSflTSfn7SfqjSftLSfvzSqCrSqFTSqH7SqKjSqNLSqPzS0irS0lTS0n7S0qjS0tLS0vzS/CrS/FTS/H7S/KjS/NLS/Pz8Kir8KlT8Kn78Kqj8KtL8Kvz8VCr8VFT8VH78VKj8VNL8VPz8fir8flT8fn78fqj8ftL8fvz8qCr8qFT8qH78qKj8qNL8qPz80ir80lT80n780qj80tL80vz8/Cr8/FT8/H78/Kj8/NL8/PwAAAAAAP8A/wAA////AAD/AP///wD///9sbGxsbJZsbMBslmxslpZslsBswGxswJZswMCWbGyWbJaWbMCWlmyWlpaWlsCWwGyWwJaWwMDAbGzAbJbAbMDAlmzAlpbAlsDAwGzAwJbAwMAAAAAAAAAAAAAAAAAAAAAI/wABCBTYZmCkFUkEJhwI4ABDSZKStElia46tSJIutYlUgiFBXpLa2LIlKdKcSxVN2uooMABFSRdHloyEcQ7GSAMDBICI8pKkJxAh2rpEFCeAFS5tgTsgwwodcFVgPnlS4slJSXPUBVghj45QWytKpLvkhufIOSRWHAgQaaS8KhZBSpKHydaTSAsBOGhwo+5TSew+zEkX44mtNm0CDNQAAgRMeeZspUsHLgPCwU8OOAQw4UOGS5E+fDA3B9yEJ0lipJuDVDEJBKop32ARzgEFhElKmJjDooTiFQ1izJlDp40SEyYElEiSJIAJdUpKAG/gIMOE4ElM5A5AYjuLPSXCr/9lAUJ0hgwHSggIgKS9gHDgwwc44KBxY+vrS7QnEQB+eBOa1WdfBg0EoB9zJRyQQ3wsNLCCA4zdtx4JSLSRjmYIuKNEEiwYAEAAEYKQAQK5LbcCOOqEsw4UVQnUQIR8TaAZCyxAQc8eV7jDom8NUQCCBiZURqB0JkBhQjiC7KhWAPUBeQOBOgWAwAE4iNGOOlCwEBaTIFBQwpMTRDnBBF3ocZwJ6Rh5AAUfeAlOAw3QiKE8x+kExRNK0AcCAtKtQGNrJtBhQpQBPAEFAA2MmCCNuiFnwhMG1tlnFTEY6KBvjkKBlBJKRNHCoMuRAAAJT1RqoE4AOucociyNmlqUBqwVtRYChALQETa2lqDCWgB4mJNOAwUEADs"
	imageSrc := "https://gaugin.frmnt.io/image?id=0d9684e197ff3a8d34bddab41e2ef8c9f6d1050242b44b56dfab11ff69b670bb"

	p := page.New("issa page")
	p.SetCustomStyles(appStyles)

	issaImage := issa_image.NewDehydrated(p, dehydratedSrc, imageSrc)
	p.Append(issaImage)

	issaPath := filepath.Join(os.TempDir(), "issa.html")
	issaFile, err := os.Create(issaPath)
	if err != nil {
		panic(err)
	}
	defer issaFile.Close()

	if err := p.WriteContent(issaFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + issaPath)
}

func createQueryFragment(r compton.Registrar) compton.Element {
	sh := section_highlight.New(r)
	sh.SetClass("fs-xs")

	shStack := flex_items.New(r, direction.Row).SetColumnGap(size.Normal)
	sh.Append(shStack)

	sp1 := els.NewSpan()
	pt1 := els.NewSpanText("Descending: ")
	pt1.SetClass("fg-subtle")
	pv1 := els.NewSpanText("True")
	pv1.SetClass("fw-b")
	sp1.Append(pt1, pv1)
	shStack.Append(sp1)

	sp2 := els.NewSpan()
	pt2 := els.NewSpanText("Sort: ")
	pt2.SetClass("fg-subtle")
	pv2 := els.NewSpanText("GOG Order Date")
	pv2.SetClass("fw-b")
	sp2.Append(pt2, pv2)
	shStack.Append(sp2)

	sp3 := els.NewSpan()
	pt3 := els.NewSpanText("Data Type: ")
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
