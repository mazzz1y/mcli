package main

import (
	"github.com/mazzz1y/mcli/internal/prompt"
	"github.com/mazzz1y/mcli/internal/shortcuts"
	"github.com/mazzz1y/mcli/internal/subprocess"
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

	return subprocess.Exec(cfg.Shortcuts[index].Cmd, cfg.Shortcuts[index].Envs)
}

func add(*cli.Context) error {
	name, err := prompt.InputPromptString(namePromptText, "")
	if err != nil {
		return err
	}

	command, err := prompt.InputPromptString(commandPromptText, "")
	if err != nil {
		return err
	}

	envs, err := prompt.InputPromptEnv(envPromptText, nil)
	if err != nil {
		return err
	}

	cfg.Shortcuts.Add(shortcuts.Shortcut{Name: name, Cmd: command, Envs: envs})
	return cfg.Write()
}

func edit(*cli.Context) error {
	index, err := prompt.SelectionPrompt(cfg.Shortcuts, cfg.PromptSize)
	if err != nil {
		return err
	}

	name, err := prompt.InputPromptString(namePromptText, cfg.Shortcuts[index].Name)
	if err != nil {
		return err
	}

	command, err := prompt.InputPromptString(commandPromptText, cfg.Shortcuts[index].Cmd)
	if err != nil {
		return err
	}

	envs, err := prompt.InputPromptEnv(envPromptText, cfg.Shortcuts[index].Envs)
	if err != nil {
		return err
	}

	cfg.Shortcuts.Delete(shortcuts.Shortcut{Index: index})
	cfg.Shortcuts.Add(shortcuts.Shortcut{Name: name, Cmd: command, Envs: envs})
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
	size, err := prompt.InputPromptInt(sizePromptText, cfg.PromptSize)
	if err != nil {
		return err
	}

	err = cfg.SetPromptSize(size)
	if err != nil {
		return err
	}

	return cfg.Write()
}
