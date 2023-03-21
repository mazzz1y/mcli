package shortcuts

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortcutAdd(t *testing.T) {
	ss := Shortcuts{}
	name := "one"
	cmd := "two"

	ss.Add(Shortcut{Name: name, Cmd: cmd})

	assert.Equal(t, 1, len(ss))
	assert.Equal(t, name, ss[0].Name)
	assert.Equal(t, cmd, ss[0].Cmd)
}

func TestShortcutDelete(t *testing.T) {
	name := "one"
	cmd := "two"

	ss := Shortcuts{
		{Name: name, Cmd: cmd},
	}

	ss.Delete(Shortcut{Index: 0})

	assert.Equal(t, 0, len(ss))
}

func TestShortcutSort(t *testing.T) {
	ss := Shortcuts{
		{Name: "c"},
		{Name: "b"},
		{Name: "a"},
	}

	ss.sort()

	assert.Equal(t, "a", ss[0].Name)
	assert.Equal(t, "b", ss[1].Name)
	assert.Equal(t, "c", ss[2].Name)
}
