package main

import (
	"errors"
	"strings"

	"github.com/manifoldco/promptui"
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
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "> {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "$ {{ .Cmd }}",
		Details:  "$ {{ .Cmd }}",
	}

	searcher := func(input string, index int) bool {
		return searchMatch(items[index], input)
	}

	prompt := promptui.Select{
		Label:             "Select Command:",
		Items:             items,
		Templates:         templates,
		Size:              size,
		Searcher:          searcher,
		StartInSearchMode: true,
	}

	i, _, err := prompt.Run()
	return i, err
}

func searchMatch(item Item, input string) bool {
	name := strings.ToLower(item.Name)
	cmd := strings.ToLower(item.Cmd)
	input = strings.ToLower(input)

	for _, in := range strings.Split(input, " ") {
		if strings.Contains(name, in) || strings.Contains(cmd, in) {
			name = strings.ReplaceAll(name, in, "")
			cmd = strings.ReplaceAll(cmd, in, "")
		} else {
			return false
		}
	}

	return true
}
