package prompt

import (
	"errors"
	"os"
	"strings"

	"github.com/dmirubtsov/mcli/pkg/shortcuts"

	"github.com/charmbracelet/lipgloss"
	"github.com/dmirubtsov/mcli/pkg/templates"
	"github.com/erikgeiser/promptkit"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/muesli/termenv"
)

func InputPrompt(label string) (string, error) {
	input := textinput.New(label + ":")
	input.Placeholder = templates.SelectionInputPlaceholderText
	return input.RunPrompt()
}

func SelectionPrompt(shortcuts shortcuts.Shortcuts, size int) (int, error) {
	var choices []*selection.Choice

	if len(shortcuts) == 0 {
		return 0, errors.New("please add your shortcuts first")
	}

	for _, shortcut := range shortcuts {
		choices = append(choices, &selection.Choice{
			String: shortcut.Name,
			Value:  termenv.String(shortcut.Cmd).Foreground(termenv.ANSI256Color(240)).String(),
		})
	}

	sel := &selection.Selection{
		Choices:                     choices,
		Template:                    templates.SelectionSelectTemplate,
		ResultTemplate:              templates.SelectionResultTemplate,
		Filter:                      selectionFilter,
		FilterInputPlaceholderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
		SelectedChoiceStyle:         selection.DefaultSelectedChoiceStyle,
		FinalChoiceStyle: func(c *selection.Choice) string {
			return termenv.String(shortcuts[c.Index].Cmd).String()
		},
		KeyMap:            selection.NewDefaultKeyMap(),
		FilterPlaceholder: templates.SelectionFilterPlaceholderText,
		WrapMode:          promptkit.Truncate,
		Output:            os.Stdout,
		Input:             os.Stdin,
		PageSize:          size,
	}

	choice, err := sel.RunPrompt()
	if err != nil {
		return 0, err
	}

	return choice.Index, err
}

func selectionFilter(filter string, choice *selection.Choice) bool {
	name := strings.ToLower(choice.String)
	cmd := strings.ToLower(choice.Value.(string))
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
