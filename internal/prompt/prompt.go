package prompt

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"github.com/mazzz1y/mcli/internal/shortcuts"
)

func InputPromptInt(prompt string, initValue int) (string, error) {
	input := textinput.New(prompt + ":")
	input.Placeholder = intPlaceholderText
	input.Validate = intValidator
	input.InitialValue = strconv.Itoa(initValue)

	return input.RunPrompt()
}

func InputPromptString(prompt string, initValue string) (string, error) {
	input := textinput.New(prompt + ":")
	input.Placeholder = stringPlaceholderText
	input.InitialValue = initValue

	return input.RunPrompt()
}

func InputPromptEnv(prompt string, initValues []shortcuts.Env) ([]shortcuts.Env, error) {
	input := textinput.New("")
	input.Placeholder = envPlaceholderText
	input.Validate = envValidator

	var envs []shortcuts.Env
	var count int
	for {
		input.Prompt = fmt.Sprintf("%s[%d]:", prompt, count)
		if len(initValues) > count {
			input.InitialValue = fmt.Sprintf("%s=%s", initValues[count].Key, initValues[count].Value)
		} else {
			input.InitialValue = ""
		}

		res, err := input.RunPrompt()
		if err != nil {
			return nil, err
		}

		if res == "" {
			break
		}

		envKV := strings.Split(res, "=")
		if len(envKV) == 2 {
			envs = append(envs, shortcuts.Env{Key: envKV[0], Value: envKV[1]})
		}
		count++
	}

	return envs, nil
}

func SelectionPrompt(ss shortcuts.Shortcuts, size int) (int, error) {
	sel := selection.New("", ss)
	sel.Template = selectionSelectTemplate
	sel.ResultTemplate = selectionResultTemplate
	sel.FilterPlaceholder = filterPlaceholderText
	sel.Filter = selectionFilter
	sel.SelectedChoiceStyle = selectedChoiceStyle
	sel.UnselectedChoiceStyle = unselectedChoiceStyle
	sel.FinalChoiceStyle = finalChoiceStyle
	sel.ExtendedTemplateFuncs = map[string]interface{}{
		"CommandStyle":       commandStyle,
		"CommandPromptStyle": commandPromptStyle,
		"FilterPromptStyle":  filterPromptStyle,
		"UpDownSymbolStyle":  upDownSymbolStyle,
		"FinalSymbolStyle":   finalSymbolStyle,
	}
	sel.PageSize = size

	choice, err := sel.RunPrompt()
	if err != nil {
		return 0, err
	}

	return choice.Index, err
}

func envValidator(s string) error {
	if s == "" {
		return nil
	}

	parts := strings.Split(s, "=")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return textinput.ErrInputValidation
	}

	return nil
}

func intValidator(s string) error {
	if i, err := strconv.Atoi(s); err != nil || i < 0 {
		return textinput.ErrInputValidation
	}

	return nil
}
