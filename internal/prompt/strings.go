package prompt

import (
	"github.com/dmirubtsov/mcli/internal/shortcuts"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/muesli/termenv"
)

const (
	StringPlaceholderText = "cannot be empty"
	IntPlaceholderText    = "should be positive integer, 0 is unlimited"
	filterPlaceholderText = "type to filter choices"
	upDownSymbolText      = "•"
	finalSymbolText       = "$"
	CommandText           = "Command"
	FilterText            = "Filter"
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
	return termenv.String(CommandText).Faint().String()
}

func filterPromptStyle() string {
	return termenv.String(FilterText).Faint().String()
}

func upDownSymbolStyle() string {
	return termenv.String(upDownSymbolText).Faint().String()
}

func finalSymbolStyle() string {
	return termenv.String(finalSymbolText).Faint().String()
}
