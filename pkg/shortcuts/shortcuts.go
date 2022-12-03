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

func (ss *Shortcuts) Add(s Shortcut) {
	*ss = append(*ss, s)
	ss.sort()
}

func (ss *Shortcuts) Delete(i Shortcut) {
	a := *ss
	a[i.Index] = a[len(a)-1]
	*ss = a[:len(a)-1]
	ss.sort()
}

func (ss *Shortcuts) sort() {
	sort.Slice(*ss, func(i, j int) bool {
		a := *ss
		return a[i].Name < a[j].Name
	})
}
