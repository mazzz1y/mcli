package prompt

import (
	"github.com/dmirubtsov/mcli/pkg/items"
	"os"
	"strings"

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

func SelectionPrompt(items items.Items, size int) (int, error) {
	var choices []*selection.Choice

	for _, item := range items {
		choices = append(choices, &selection.Choice{
			String: item.Name,
			Value:  termenv.String(item.Cmd).Foreground(termenv.ANSI256Color(240)).String(),
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
			return termenv.String(items[c.Index].Cmd).String()
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