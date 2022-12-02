package main

import (
	"errors"
	"os"
	"strings"
	"text/template"

	"github.com/charmbracelet/lipgloss"
	"github.com/erikgeiser/promptkit"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/manifoldco/promptui"
	"github.com/muesli/termenv"
)

func prompt(label string) (string, error) {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | green }} ",
	}

	validate := func(input string) error {
		if input != "" {
			return nil
		}
		return errors.New(label + " cannot be empty")
	}

	prompt := promptui.Prompt{
		Label:     label + ":",
		Validate:  validate,
		Templates: templates,
	}
	return prompt.Run()
}

func selectItem(items Items, size int) (int, error) {

	var itemsNames []string

	for _, item := range items {
		itemsNames = append(itemsNames, item.Name)
	}

	sel := &selection.Selection{
		Choices:                     selection.Choices(itemsNames),
		FilterPrompt:                selection.DefaultFilterPrompt,
		Template:                    selection.DefaultTemplate,
		ResultTemplate:              selection.DefaultResultTemplate,
		Filter:                      filter,
		FilterInputPlaceholderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
		SelectedChoiceStyle:         selection.DefaultSelectedChoiceStyle,
		FinalChoiceStyle: func(c *selection.Choice) string {
			return termenv.String(items[c.Index].Cmd).Foreground(termenv.ANSI256Color(32)).String()
		},
		KeyMap:                selection.NewDefaultKeyMap(),
		FilterPlaceholder:     selection.DefaultFilterPlaceholder,
		ExtendedTemplateFuncs: template.FuncMap{},
		WrapMode:              promptkit.Truncate,
		Output:                os.Stdout,
		Input:                 os.Stdin,
		PageSize:              10,
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
