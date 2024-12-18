package compton

import (
	"golang.org/x/exp/maps"
	"sort"
)

type Target struct {
	Title   string
	Href    string
	Icon    Symbol
	Current bool
}

func TextLinks(links map[string]string, selected string, order ...string) []*Target {
	if len(order) == 0 {
		order = maps.Keys(links)
		sort.Strings(order)
	}

	targets := make([]*Target, 0, len(links))

	for _, key := range order {
		if _, ok := links[key]; !ok {
			continue
		}
		t := &Target{
			Title:   key,
			Href:    links[key],
			Current: key == selected,
		}
		targets = append(targets, t)
	}

	return targets
}

func SetIcons(targets []*Target, icons map[string]Symbol) []*Target {
	for _, t := range targets {
		t.Icon = icons[t.Title]
	}
	return targets
}
