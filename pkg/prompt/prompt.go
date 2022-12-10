package prompt

import (
	"github.com/dmirubtsov/mcli/pkg/shortcuts"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	"strconv"
)

func InputPromptInt(label string) (string, error) {
	input := textinput.New(label + ":")
	input.Placeholder = IntPlaceholderText
	input.Validate = intValidator

	return input.RunPrompt()
}

func InputPromptString(label string) (string, error) {
	input := textinput.New(label + ":")
	input.Placeholder = StringPlaceholderText

	return input.RunPrompt()
}

func SelectionPrompt(ss shortcuts.Shortcuts, size int) (int, error) {
	sel := selection.New("", ss)
	sel.Template = SelectionSelectTemplate
	sel.ResultTemplate = SelectionResultTemplate
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

func intValidator(s string) error {
	if i, err := strconv.Atoi(s); err != nil || i < 0 {
		return textinput.ErrInputValidation
	}

	return nil
}
