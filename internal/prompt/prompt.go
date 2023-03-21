package prompt

import (
	"strconv"

	"github.com/dmirubtsov/mcli/internal/shortcuts"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
)

func InputPromptInt(label string, initValue int) (string, error) {
	input := textinput.New(label + ":")
	input.Placeholder = IntPlaceholderText
	input.Validate = intValidator
	input.InitialValue = strconv.Itoa(initValue)

	return input.RunPrompt()
}

func InputPromptString(label string, initValue string) (string, error) {
	input := textinput.New(label + ":")
	input.Placeholder = StringPlaceholderText
	input.InitialValue = initValue

	return input.RunPrompt()
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

func intValidator(s string) error {
	if i, err := strconv.Atoi(s); err != nil || i < 0 {
		return textinput.ErrInputValidation
	}

	return nil
}
