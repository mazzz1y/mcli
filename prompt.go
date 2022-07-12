package main

import (
	"errors"
	"github.com/manifoldco/promptui"
	"strings"
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

func selectItem(items Items) (int, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "> {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "> {{ .Cmd }}",
	}

	searcher := func(input string, index int) bool {
		item := items[index]
		name := strings.Replace(strings.ToLower(item.Name), " ", "", -1)
		cmd := strings.Replace(strings.ToLower(item.Cmd), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(cmd, input)
	}

	prompt := promptui.Select{
		Label:     "Select Command:",
		Items:     items,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}
	prompt.StartInSearchMode = true

	i, _, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	return i, nil
}
