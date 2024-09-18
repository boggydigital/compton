package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/consts/weight"
	"github.com/boggydigital/compton/elements/details_summary"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/flex_items"
	"github.com/boggydigital/compton/elements/fspan"
	"github.com/boggydigital/compton/elements/grid_items"
	"github.com/boggydigital/compton/elements/iframe_expand"
	"github.com/boggydigital/compton/elements/inputs"
	"github.com/boggydigital/compton/elements/issa_image"
	"github.com/boggydigital/compton/elements/nav_links"
	"github.com/boggydigital/compton/elements/section"
	"github.com/boggydigital/compton/elements/svg_use"
	"github.com/boggydigital/compton/elements/title_values"
	"github.com/boggydigital/compton/page"
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
	//writeIframeContent()
	//writeIssaPage()
	//writeSvgUsePage()
}

func writeTestPage() {
	p := page.Page("test").SetFavIconEmoji("🤔")
	p.SetCustomStyles(appStyles)

	s := flex_items.FlexItems(p, direction.Column)

	topNavLinks := map[string]string{
		"Updates": "/updates",
		"Search":  "/search",
	}

	topNavIcons := map[string]svg_use.Symbol{
		"Updates": svg_use.Sparkle,
		"Search":  svg_use.Search,
	}

	targets := nav_links.TextLinks(
		topNavLinks,
		"Search",
		"Updates", "Search")
	nav_links.SetIcons(targets, topNavIcons)

	topNav := nav_links.NavLinksTargets(p, targets...)

	s.Append(topNav)

	navLinks := map[string]string{
		"New":      "/new",
		"Owned":    "/owned",
		"Wishlist": "/wishlist",
		"Sale":     "/sale",
		"All":      "/all",
	}

	nav := nav_links.NavLinksTargets(p,
		nav_links.TextLinks(
			navLinks,
			"New",
			"New",
			"Owned",
			"Wishlist",
			"Sale",
			"All")...)

	s.Append(nav)

	cdc := details_summary.Open(p, "Filter & Search").
		BackgroundColor(color.Highlight)

	form := els.Form("/action", "GET")

	formStack := flex_items.
		FlexItems(p, direction.Column)

	qf := createQueryFragment(p)

	formStack.Append(qf)

	submitRow := flex_items.FlexItems(p, direction.Row).JustifyContent(align.Center)

	submit := inputs.InputValue(p, input_types.Submit, "Submit Query")
	submitRow.Append(submit)
	formStack.Append(submitRow)

	tiGrid := grid_items.GridItems(p).JustifyContent(align.Center)

	ti1 := title_values.SearchValue(p, "Title", "title", "Hello")

	tiList := map[string]string{
		"true":  "True",
		"false": "False",
		"maybe": "Maybe",
	}

	ti2 := title_values.Search(p, "Description", "description").
		SetDataList(tiList, "")

	ti3 := title_values.Search(p, "Descending", "desc")
	ti4 := title_values.Search(p, "Sort", "sort")
	ti5 := title_values.SearchValue(p, "Data Type", "data-type", "Account Products")

	tiGrid.Append(ti1, ti2, ti3, ti4, ti5)

	formStack.Append(tiGrid)
	form.Append(formStack)

	cdc.Append(form)
	s.Append(cdc)

	s.Append(qf)

	cdo := details_summary.
		Open(p, "Title Values").
		BackgroundColor(color.Purple).
		ForegroundColor(color.Background)

	tvGrid := grid_items.GridItems(p).JustifyContent(align.Center)

	tvLinks := map[string]string{
		"Achievements":       "/achievements",
		"Controller support": "/controller-support",
		"Overlay":            "/overlay",
		"Single-player":      "/single-player",
	}
	tv1 := title_values.TitleValues(p, "Features").AppendTextValues(maps.Keys(tvLinks)...)
	tv2 := title_values.TitleValues(p, "Feature Links").AppendLinkValues(tvLinks)
	tv3 := title_values.TitleValues(p, "Features").AppendTextValues(maps.Keys(tvLinks)...)
	tv4 := title_values.TitleValues(p, "Feature Links").AppendLinkValues(tvLinks)
	tv5 := title_values.TitleValues(p, "Features").AppendTextValues(maps.Keys(tvLinks)...)
	tv6 := title_values.TitleValues(p, "Feature Links").AppendLinkValues(tvLinks)

	tvGrid.Append(tv1, tv2, tv3, tv4, tv5, tv6)
	cdo.Append(tvGrid)
	s.Append(cdo)

	footer := flex_items.FlexItems(p, direction.Row).JustifyContent(align.Center)

	div := fspan.Text(p, "").ForegroundColor(color.Subtle).FontSize(size.XSmall)

	div.Append(els.Text("Last updated: "),
		els.TimeText(time.Now().Format("2006-01-02 15:04:05")))

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

	c := page.Page("content")
	ifec := iframe_expand.IframeExpandContent("test", "whatever")
	c.Append(ifec)

	for i := range 1000 {
		c.Append(els.DivText(strconv.Itoa(i)))
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

	p := page.Page("iframe")

	dc := details_summary.Open(p, "Description")

	ife := iframe_expand.IframeExpandHost(p, "test", "content.html")
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

	p := page.Page("issa page")
	p.SetCustomStyles(appStyles)

	issaImage := issa_image.IssaImageDehydrated(p, dehydratedSrc, imageSrc)
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
	sh := section.Section(r).
		BackgroundColor(color.Highlight).
		FontSize(size.XSmall)

	shStack := flex_items.FlexItems(r, direction.Row)
	sh.Append(shStack)

	sp1 := els.Span()
	pt1 := fspan.Text(r, "Descending: ").
		ForegroundColor(color.Subtle)
	pv1 := fspan.Text(r, "True").
		FontWeight(weight.Bolder)
	sp1.Append(pt1, pv1)
	shStack.Append(sp1)

	sp2 := els.Span()
	pt2 := fspan.Text(r, "Sort: ").
		ForegroundColor(color.Subtle)
	pv2 := fspan.Text(r, "GOG Order Date").
		FontWeight(weight.Bolder)
	sp2.Append(pt2, pv2)
	shStack.Append(sp2)

	sp3 := els.Span()
	pt3 := fspan.Text(r, "Data Type: ").
		ForegroundColor(color.Subtle)
	pv3 := fspan.Text(r, "Account Products").
		FontWeight(weight.Bolder)
	sp3.Append(pt3, pv3)
	shStack.Append(sp3)

	clearAction := els.AText("Clear", "/clear")
	clearAction.AddClass("action")
	shStack.Append(clearAction)

	return sh
}

func writeSvgUsePage() {
	p := page.Page("svg use page")

	p.Append(svg_use.SvgUse(p, svg_use.MacOS))

	svgUsePath := filepath.Join(os.TempDir(), "svg_use.html")
	svgUseFile, err := os.Create(svgUsePath)
	if err != nil {
		panic(err)
	}
	defer svgUseFile.Close()

	if err := p.WriteContent(svgUseFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + svgUsePath)

}
