package prompt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"

	"github.com/dmirubtsov/mcli/internal/shortcuts"
	"github.com/erikgeiser/promptkit/selection"
)

func TestSelectionFilter(t *testing.T) {
	name := []string{"one", "two", "three"}
	cmd := []string{"four", "five", "six"}

	shortcut := shortcuts.Shortcut{
		Name: strings.Join(name, " "),
		Cmd:  strings.Join(cmd, " "),
	}

	tt := []struct {
		shortcut shortcuts.Shortcut
		input    string
		result   bool
	}{
		{
			shortcut: shortcut,
			input:    name[0],
			result:   true,
		},
		{
			shortcut: shortcut,
			input:    cmd[0],
			result:   true,
		},
		{
			shortcut: shortcut,
			input:    shortcut.Cmd,
			result:   true,
		},
		{
			shortcut: shortcut,
			input:    fmt.Sprintf("%s %s", name[1], name[2]),
			result:   true,
		},
		{
			shortcut: shortcut,
			input:    fmt.Sprintf("%s %s", name[2], name[1]),
			result:   true,
		},
		{
			shortcut: shortcut,
			input:    fmt.Sprintf("%s %s", name[2][:2], cmd[1][:2]),
			result:   true,
		},
		{
			shortcut: shortcut,
			input:    fmt.Sprintf("%s %s %s", name[2], name[1], "false"),
			result:   false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, selectionFilter(tc.input, &selection.Choice[shortcuts.Shortcut]{
				Value: shortcut,
			}), tc.result)
		})
	}
}
