package prompt

import (
	"github.com/dmirubtsov/mcli/internal/shortcuts"
	"github.com/erikgeiser/promptkit/selection"
	"strings"
)

func selectionFilter(filter string, c *selection.Choice[shortcuts.Shortcut]) bool {
	name := strings.ToLower(c.Value.Name)
	cmd := strings.ToLower(c.Value.Cmd)
	filter = strings.ToLower(filter)

	for _, in := range strings.Split(filter, " ") {
		match := false

		if strings.Contains(name, in) {
			name = strings.ReplaceAll(name, in, "")
			match = true
		} else if strings.Contains(cmd, in) {
			cmd = strings.ReplaceAll(cmd, in, "")
			match = true
		}

		if !match {
			return match
		}
	}

	return true
}
