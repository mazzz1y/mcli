package shortcuts

import (
	"sort"
)

type Shortcut struct {
	Name  string `json:"name"`
	Cmd   string `json:"cmd"`
	Envs  []Env  `json:"env,omitempty"`
	Index int    `json:"-"`
}

type Env struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Shortcuts []Shortcut

func (ss *Shortcuts) Add(s Shortcut) {
	*ss = append(*ss, s)
	ss.sort()
}

func (ss *Shortcuts) Delete(s Shortcut) {
	a := *ss
	a[s.Index] = a[len(a)-1]
	*ss = a[:len(a)-1]
	ss.sort()
}

func (ss *Shortcuts) sort() {
	sort.Slice(*ss, func(i, j int) bool {
		a := *ss
		return a[i].Name < a[j].Name
	})
}
