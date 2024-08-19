package nav_links

import (
	"github.com/boggydigital/compton/svg_icons"
	"golang.org/x/exp/maps"
	"sort"
)

type Target struct {
	Title   string
	Href    string
	Icon    svg_icons.Symbol
	Current bool
}

func TextLinks(links map[string]string, selected string, order ...string) []*Target {
	if len(order) == 0 {
		order = maps.Keys(links)
		sort.Strings(order)
	}

	targets := make([]*Target, 0, len(links))

	for _, key := range order {
		t := &Target{
			Title:   key,
			Href:    links[key],
			Current: key == selected,
		}
		targets = append(targets, t)
	}

	return targets
}

func SetIcons(targets []*Target, icons map[string]svg_icons.Symbol) []*Target {
	for _, t := range targets {
		t.Icon = icons[t.Title]
	}
	return targets
}
