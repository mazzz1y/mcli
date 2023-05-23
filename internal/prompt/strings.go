package prompt

import (
	"github.com/erikgeiser/promptkit/selection"
	"github.com/mazzz1y/mcli/internal/shortcuts"
	"github.com/muesli/termenv"
)

const (
	envPlaceholderText    = "key=value, empty string to end input"
	stringPlaceholderText = "cannot be empty"
	intPlaceholderText    = "should be positive integer, 0 is unlimited"
	filterPlaceholderText = "type to filter choices"
	commandText           = "Command"
	filterText            = "Filter"
	upDownSymbolText      = "•"
	finalSymbolText       = "$"
)

func selectedChoiceStyle(c *selection.Choice[shortcuts.Shortcut]) string {
	return termenv.String(c.Value.Name).Foreground(termenv.ANSI256Color(32)).Bold().String()
}

func finalChoiceStyle(c *selection.Choice[shortcuts.Shortcut]) string {
	return termenv.String(c.Value.Cmd).Faint().String()
}

func unselectedChoiceStyle(c *selection.Choice[shortcuts.Shortcut]) string {
	return termenv.String(c.Value.Name).Faint().String()
}

func commandStyle(c *selection.Choice[shortcuts.Shortcut]) string {
	return termenv.String(c.Value.Cmd).Foreground(termenv.ANSI256Color(240)).String()
}

func commandPromptStyle() string {
	return termenv.String(commandText).Faint().String()
}

func filterPromptStyle() string {
	return termenv.String(filterText).Faint().String()
}

func upDownSymbolStyle() string {
	return termenv.String(upDownSymbolText).Faint().String()
}

func finalSymbolStyle() string {
	return termenv.String(finalSymbolText).Faint().String()
}
