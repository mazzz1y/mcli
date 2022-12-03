package shortcuts

import (
	"sort"
)

type Shortcut struct {
	Name  string `json:"name"`
	Cmd   string `json:"cmd"`
	Index int    `json:"-"`
}

type Shortcuts []Shortcut

func (sh *Shortcuts) Add(i Shortcut) {
	*sh = append(*sh, i)
	sh.sort()
}

func (sh *Shortcuts) Delete(i Shortcut) {
	a := *sh
	a[i.Index] = a[len(a)-1]
	*sh = a[:len(a)-1]
	sh.sort()
}

func (sh *Shortcuts) sort() {
	sort.Slice(*sh, func(i, j int) bool {
		a := *sh
		return a[i].Name < a[j].Name
	})
}
