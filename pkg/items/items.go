package items

import (
	"sort"
)

type Item struct {
	Name  string `json:"name"`
	Cmd   string `json:"cmd"`
	Index int    `json:"-"`
}

type Items []Item

func (it *Items) Add(i Item) {
	*it = append(*it, i)
	it.sort()
}

func (it *Items) Delete(i Item) {
	a := *it
	a[i.Index] = a[len(a)-1]
	*it = a[:len(a)-1]
	it.sort()
}

func (it *Items) sort() {
	sort.Slice(*it, func(i, j int) bool {
		a := *it
		return a[i].Name < a[j].Name
	})
}
