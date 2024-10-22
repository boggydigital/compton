package main

import (
	_ "embed"
	"fmt"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/input_types"
	"github.com/boggydigital/compton/consts/size"
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
	p := compton.Page("test")
	p.AppendStyle("app-style", appStyles)

	s := compton.FlexItems(p, direction.Column)

	topNavLinks := map[string]string{
		"Updates": "/updates",
		"Search":  "/search",
	}

	topNavIcons := map[string]compton.Symbol{
		"Updates": compton.Sparkle,
		"Search":  compton.Search,
	}

	targets := compton.TextLinks(
		topNavLinks,
		"Search",
		"Updates", "Search")
	compton.SetIcons(targets, topNavIcons)

	topNav := compton.NavLinksTargets(p, targets...)

	navLinks := map[string]string{
		"New":      "/new",
		"Owned":    "/owned",
		"Wishlist": "/wishlist",
		"Sale":     "/sale",
		"All":      "/all",
	}

	nav := compton.NavLinksTargets(p,
		compton.TextLinks(
			navLinks,
			"New",
			"New",
			"Owned",
			"Wishlist",
			"Sale",
			"All")...)

	s.Append(compton.FICenter(p, topNav, nav))

	filterSearchTitle := compton.Fspan(p, "Filter & Search").
		FontWeight(font_weight.Bolder).
		FontSize(size.Large)

	dsFilterSearch := compton.DSLarge(p, filterSearchTitle, true).
		BackgroundColor(color.Highlight)

	form := compton.Form("/action", "GET")

	formStack := compton.FlexItems(p, direction.Column)

	qf := createQueryFragment(p)

	formStack.Append(qf)

	submitRow := compton.FlexItems(p, direction.Row).JustifyContent(align.Center)

	submit := compton.InputValue(p, input_types.Submit, "Submit Query")
	submitRow.Append(submit)
	formStack.Append(submitRow)

	tiGrid := compton.GridItems(p).JustifyContent(align.Center)

	ti1 := compton.TISearchValue(p, "Title", "title", "Hello")

	tiList := map[string]string{
		"true":  "True",
		"false": "False",
		"maybe": "Maybe",
	}

	ti2 := compton.TISearch(p, "Description", "description").
		SetDatalist(tiList, "")

	ti3 := compton.TISearch(p, "Descending", "desc")
	ti4 := compton.TISearch(p, "Sort", "sort")
	ti5 := compton.TISearchValue(p, "Data Type", "data-type", "Account Products")

	tiGrid.Append(ti1, ti2, ti3, ti4, ti5)

	formStack.Append(tiGrid)
	form.Append(formStack)

	dsFilterSearch.Append(form)
	s.Append(dsFilterSearch)

	s.Append(qf)

	tvsTitle := compton.Fspan(p, "Title Values").
		FontWeight(font_weight.Bolder).
		FontSize(size.Large)
	dsTitleValues := compton.DSLarge(p, tvsTitle, true).
		BackgroundColor(color.Purple).
		ForegroundColor(color.Background).
		MarkerColor(color.Yellow)

	tvGrid := compton.GridItems(p).JustifyContent(align.Center)

	tvLinks := map[string]string{
		"Achievements":       "/achievements",
		"Controller support": "/controller-support",
		"Overlay":            "/overlay",
		"Single-player":      "/single-player",
	}
	tv1 := compton.TitleValues(p, "Features").AppendTextValues(maps.Keys(tvLinks)...)
	tv2 := compton.TitleValues(p, "Feature Links").AppendLinkValues(tvLinks)
	tv3 := compton.TitleValues(p, "Features").AppendTextValues(maps.Keys(tvLinks)...)
	tv4 := compton.TitleValues(p, "Feature Links").AppendLinkValues(tvLinks)
	tv5 := compton.TitleValues(p, "Features").AppendTextValues(maps.Keys(tvLinks)...)
	tv6 := compton.TitleValues(p, "Feature Links").AppendLinkValues(tvLinks)
	tv7 := compton.TitleValues(p, "Lots of values")
	dsTitle := compton.Fspan(p, "Expand all...").
		FontWeight(font_weight.Bolder)
	dsValues := compton.DSSmall(p, dsTitle, false)
	for ii := range 10 {
		element := compton.Fspan(p, "Element "+strconv.Itoa(ii)+"&nbsp;").ForegroundColor(color.Gray)
		dsValues.Append(element)
	}
	tv7.Append(dsValues)

	tvGrid.Append(tv1, tv2, tv3, tv4, tv5, tv6, tv7)
	dsTitleValues.Append(tvGrid)
	s.Append(dsTitleValues)

	switchesTitle := compton.Fspan(p, "Switches").
		FontWeight(font_weight.Bolder).
		FontSize(size.Large)
	dsSwitches := compton.DSLarge(p, switchesTitle, true).
		BackgroundColor(color.Highlight).
		SummaryRowGap(size.XSmall)

	subTitle := compton.Fspan(p, "This section has subtitle").
		FontSize(size.XSmall).
		FontWeight(font_weight.Normal).
		ForegroundColor(color.Gray)

	dsSwitches.AppendSummary(subTitle)

	swColumn := compton.FlexItems(p, direction.Column).AlignContent(align.Center)

	sw1 := switchLabel(p, "test1", "Some very important switch")
	sw2 := switchLabel(p, "test2", "Another, equally important switch")
	sw3 := switchLabel(p, "test3", "The last important switch")

	swColumn.Append(sw1, sw2, sw3)

	dsSwitches.Append(swColumn)

	s.Append(dsSwitches)

	footer := compton.FlexItems(p, direction.Row).JustifyContent(align.Center)

	div := compton.Fspan(p, "").ForegroundColor(color.Gray).FontSize(size.Small)

	div.Append(compton.Text("Last updated: "),
		compton.TimeText(time.Now().Format("2006-01-02 15:04:05")))

	footer.Append(div)

	s.Append(footer)

	p.Append(s)

	testPath := filepath.Join(os.TempDir(), "test.html")
	testFile, err := os.Create(testPath)
	if err != nil {
		panic(err)
	}
	defer testFile.Close()

	if err := p.Write(testFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + testPath)
}

func switchLabel(r compton.Registrar, id, label string) compton.Element {
	row := compton.FlexItems(r, direction.Row)

	switchElement := compton.Switch(r)
	switchElement.SetId(id)

	labelElement := compton.Label(id)
	labelElement.Append(compton.Text(label))

	row.Append(switchElement, labelElement)

	return row
}

func writeIframeContent() {

	c := compton.Page("content")
	ifec := compton.IframeExpandContent("test", "whatever")
	c.Append(ifec)

	for i := range 1000 {
		c.Append(compton.DivText(strconv.Itoa(i)))
	}

	contentPath := filepath.Join(os.TempDir(), "content.html")
	contentFile, err := os.Create(contentPath)
	if err != nil {
		panic(err)
	}
	defer contentFile.Close()

	if err := c.Write(contentFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + contentPath)

	p := compton.Page("iframe")

	dc := compton.DSLarge(p, compton.HeadingText("Description", 2), true)

	ife := compton.IframeExpandHost(p, "test", "content.html")
	dc.Append(ife)

	p.Append(dc)

	iframePath := filepath.Join(os.TempDir(), "iframe.html")
	iframeFile, err := os.Create(iframePath)
	if err != nil {
		panic(err)
	}

	if err := p.Write(iframeFile); err != nil {
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

	p := compton.Page("issa page")
	p.AppendStyle("app-style", appStyles)

	issaImage := compton.IssaImageDehydrated(p, dehydratedSrc, imageSrc)
	p.Append(issaImage)

	issaPath := filepath.Join(os.TempDir(), "issa.html")
	issaFile, err := os.Create(issaPath)
	if err != nil {
		panic(err)
	}
	defer issaFile.Close()

	if err := p.Write(issaFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + issaPath)
}

func createQueryFragment(r compton.Registrar) compton.Element {

	shStack := compton.FlexItems(r, direction.Row).FontSize(size.Small)

	sp1 := compton.Span()
	pt1 := compton.Fspan(r, "Descending: ").
		ForegroundColor(color.Gray)
	pv1 := compton.Fspan(r, "True").
		FontWeight(font_weight.Bolder)
	sp1.Append(pt1, pv1)
	shStack.Append(sp1)

	sp2 := compton.Span()
	pt2 := compton.Fspan(r, "Sort: ").
		ForegroundColor(color.Gray)
	pv2 := compton.Fspan(r, "GOG Order Date").
		FontWeight(font_weight.Bolder)
	sp2.Append(pt2, pv2)
	shStack.Append(sp2)

	sp3 := compton.Span()
	pt3 := compton.Fspan(r, "Data Type: ").
		ForegroundColor(color.Gray)
	pv3 := compton.Fspan(r, "Account Products").
		FontWeight(font_weight.Bolder)
	sp3.Append(pt3, pv3)
	shStack.Append(sp3)

	clearAction := compton.AText("Clear", "/clear")
	clearAction.AddClass("action")
	shStack.Append(clearAction)

	return compton.FICenter(r, shStack)
}

func writeSvgUsePage() {
	p := compton.Page("svg use page")

	p.Append(compton.SvgUse(p, compton.MacOS))

	svgUsePath := filepath.Join(os.TempDir(), "svg_use.html")
	svgUseFile, err := os.Create(svgUsePath)
	if err != nil {
		panic(err)
	}
	defer svgUseFile.Close()

	if err := p.Write(svgUseFile); err != nil {
		panic(err)
	}

	fmt.Println("file://" + svgUsePath)

}
