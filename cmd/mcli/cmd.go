package main

import (
	"github.com/dmirubtsov/mcli/pkg/prompt"
	"github.com/dmirubtsov/mcli/pkg/shortcuts"
	"github.com/dmirubtsov/mcli/pkg/subprocess"
	"github.com/urfave/cli/v2"
)

func run(*cli.Context) error {
	if len(cfg.Shortcuts) == 0 {
		err := cfg.WriteDemo()
		if err != nil {
			return err
		}
	}

	index, err := prompt.SelectionPrompt(cfg.Shortcuts, cfg.PromptSize)
	if err != nil {
		return err
	}

	return subprocess.Exec(cfg.Shortcuts[index].Cmd)
}

func add(*cli.Context) error {
	nameField, err := prompt.InputPromptString(nameFieldText)
	if err != nil {
		return err
	}

	commandField, err := prompt.InputPromptString(commandFieldText)
	if err != nil {
		return err
	}

	cfg.Shortcuts.Add(shortcuts.Shortcut{Name: nameField, Cmd: commandField})
	return cfg.Write()
}

func edit(*cli.Context) error {
	index, err := prompt.SelectionPrompt(cfg.Shortcuts, cfg.PromptSize)
	if err != nil {
		return err
	}

	nameField, err := prompt.InputPromptString(nameFieldText)
	if err != nil {
		return err
	}

	commandField, err := prompt.InputPromptString(commandFieldText)
	if err != nil {
		return err
	}

	cfg.Shortcuts.Delete(shortcuts.Shortcut{Index: index})
	cfg.Shortcuts.Add(shortcuts.Shortcut{Name: nameField, Cmd: commandField})
	return cfg.Write()
}

func delete(*cli.Context) error {
	index, err := prompt.SelectionPrompt(cfg.Shortcuts, cfg.PromptSize)
	if err != nil {
		return err
	}

	cfg.Shortcuts.Delete(shortcuts.Shortcut{Index: index})
	return cfg.Write()
}

func setPromptSize(*cli.Context) error {
	size, err := prompt.InputPromptInt(sizeFieldText)
	if err != nil {
		return err
	}

	err = cfg.SetPromptSize(size)
	if err != nil {
		return err
	}

	return cfg.Write()
}
