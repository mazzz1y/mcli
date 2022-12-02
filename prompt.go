package main

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/erikgeiser/promptkit"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/muesli/termenv"
)

const template = `
{{- if .Prompt -}}
  {{ Bold .Prompt }}
{{ end -}}
{{ if .IsFiltered }}
  {{- print "Filter: " " " .FilterInput }}
{{ end }}
{{- if not (eq (len .Choices) 0)}}
{{- print "Command: " (index .Choices $.SelectedIndex).Value}}
{{- print "\n"}}
{{- end }}

{{- range  $i, $choice := .Choices }}
  {{- if IsScrollUpHintPosition $i }}
	{{- "⇡ " -}}
  {{- else if IsScrollDownHintPosition $i -}}
	{{- "⇣ " -}}
  {{- else -}}
	{{- "  " -}}
  {{- end -}}

  {{- if eq $.SelectedIndex $i }}
   {{- print (Foreground "32" (Bold "~ ")) (Selected $choice) "\n" }}
  {{- else }}
	{{- print "  " (Unselected $choice) "\n" }}
  {{- end }}
{{- end}}
`

func prompt(label string) (string, error) {
	input := textinput.New(label)
	input.Placeholder = "cannot be empty"
	return input.RunPrompt()
}

func selectItem(items Items, size int) (int, error) {
	var ch []*selection.Choice

	for _, item := range items {
		ch = append(ch, &selection.Choice{
			String: item.Name,
			Value:  termenv.String(item.Cmd).Foreground(termenv.ANSI256Color(240)).String(),
		})
	}

	sel := &selection.Selection{
		Choices:                     ch,
		Template:                    template,
		ResultTemplate:              selection.DefaultResultTemplate,
		Filter:                      filter,
		FilterInputPlaceholderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
		SelectedChoiceStyle:         selection.DefaultSelectedChoiceStyle,
		FinalChoiceStyle: func(c *selection.Choice) string {
			return termenv.String(items[c.Index].Cmd).Foreground(termenv.ANSI256Color(32)).String()
		},
		KeyMap:            selection.NewDefaultKeyMap(),
		FilterPlaceholder: selection.DefaultFilterPlaceholder,
		WrapMode:          promptkit.Truncate,
		Output:            os.Stdout,
		Input:             os.Stdin,
		PageSize:          config.PromptSize,
	}

	choice, err := sel.RunPrompt()
	if err != nil {
		return 0, err
	}

	return choice.Index, err
}

func filter(filter string, choice *selection.Choice) bool {
	name := strings.ToLower(config.Items[choice.Index].Name)
	cmd := strings.ToLower(config.Items[choice.Index].Cmd)
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
