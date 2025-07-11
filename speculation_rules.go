package compton

import (
	"bytes"
	"encoding/json"
)

const speculationRulesName = "speculationrules"

type SpeculationRulesEagerness int

const (
	SpeculationRulesImmediateEagerness SpeculationRulesEagerness = iota
	SpeculationRulesEagerEagerness
	SpeculationRulesModerateEagerness
	SpeculationRulesConservativeEagerness
)

var speculationRulesEagernessStrings = map[SpeculationRulesEagerness]string{
	SpeculationRulesImmediateEagerness:    "immediate",
	SpeculationRulesEagerEagerness:        "eager",
	SpeculationRulesModerateEagerness:     "moderate",
	SpeculationRulesConservativeEagerness: "conservative",
}

type UriMatch struct {
	HrefMatches     string    `json:"href_matches,omitempty"`
	SelectorMatches string    `json:"selector_matches,omitempty"`
	Not             *UriMatch `json:"not,omitempty"`
}

type SpeculationRules struct {
	Prerender []SpeculationRulesPrerender `json:"prerender"`
}

type SpeculationRulesPrerender struct {
	Source string `json:"source"`
	Where  struct {
		And []*UriMatch `json:"and"`
	} `json:"where"`
	Eagerness string `json:"eagerness"`
}

func SpeculationRulesBytes(eagerness SpeculationRulesEagerness, hrefMatches ...string) []byte {

	var srEagerness string
	if sre, ok := speculationRulesEagernessStrings[eagerness]; ok {
		srEagerness = sre
	} else {
		srEagerness = speculationRulesEagernessStrings[SpeculationRulesModerateEagerness]
	}

	sr := new(SpeculationRules)
	srp := SpeculationRulesPrerender{
		Source:    "document",
		Eagerness: srEagerness,
	}

	for _, hr := range hrefMatches {
		srp.Where.And = append(srp.Where.And, &UriMatch{HrefMatches: hr})
	}

	sr.Prerender = append(sr.Prerender, srp)

	var bts []byte
	buf := bytes.NewBuffer(bts)

	if err := json.NewEncoder(buf).Encode(sr); err != nil {
		return nil
	}

	return buf.Bytes()
}
