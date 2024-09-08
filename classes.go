package compton

import (
	"golang.org/x/exp/maps"
	"strings"
)

type ClassList struct {
	classList map[string]any
}

func (cl *ClassList) AddClass(classes ...string) {
	if cl.classList == nil {
		cl.classList = make(map[string]any)
	}
	for _, class := range classes {
		cl.classList[class] = nil
	}
}

func (cl *ClassList) RemoveClass(classes ...string) {
	if cl.classList == nil {
		cl.classList = make(map[string]any)
	}
	for _, class := range classes {
		delete(cl.classList, class)
	}
}

func (cl *ClassList) HasClass(classes ...string) bool {
	for _, class := range classes {
		if _, ok := cl.classList[class]; !ok {
			return false
		}
	}
	return true
}

func (cl *ClassList) ToggleClass(classes ...string) {
	for _, class := range classes {
		if cl.HasClass(class) {
			cl.RemoveClass(class)
		} else {
			cl.AddClass(class)
		}
	}
}

func (cl *ClassList) String() string {
	return strings.Join(maps.Keys(cl.classList), " ")
}
