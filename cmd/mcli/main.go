package main

import (
	"errors"
	"log"
	"os"

	"github.com/dmirubtsov/mcli/pkg/prompt"
	"github.com/dmirubtsov/mcli/pkg/shortcuts"
	"github.com/dmirubtsov/mcli/pkg/subprocess"

	"github.com/dmirubtsov/mcli/pkg/config"
	"github.com/erikgeiser/promptkit"
	"github.com/urfave/cli/v2"
)

var version = "git"

func main() {
	var config = config.Config{}
	if err := config.Read(); err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:    "mcli",
		Version: version,
		Usage:   "Simple shortcut menu",
		Action: func(c *cli.Context) error {
			index, err := prompt.SelectionPrompt(config.Shortcuts, config.PromptSize)
			if err != nil {
				return err
			}

			return subprocess.Exec(config.Shortcuts[index].Cmd)
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add shortcut",
				Action: func(c *cli.Context) error {
					nameField, err := prompt.InputPrompt("Name")
					if err != nil {
						return err
					}

					commandField, err := prompt.InputPrompt("Command")
					if err != nil {
						return err
					}

					config.Shortcuts.Add(shortcuts.Shortcut{Name: nameField, Cmd: commandField})
					return config.Write()
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "Edit shortcut",
				Action: func(c *cli.Context) error {
					index, err := prompt.SelectionPrompt(config.Shortcuts, config.PromptSize)
					if err != nil {
						return err
					}

					nameField, err := prompt.InputPrompt("Name")
					if err != nil {
						return err
					}

					commandField, err := prompt.InputPrompt("Command")
					if err != nil {
						return err
					}

					config.Shortcuts.Delete(shortcuts.Shortcut{Index: index})
					config.Shortcuts.Add(shortcuts.Shortcut{Name: nameField, Cmd: commandField})
					return config.Write()
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Remove shortcut",
				Action: func(c *cli.Context) error {
					index, err := prompt.SelectionPrompt(config.Shortcuts, config.PromptSize)
					if err != nil {
						return err
					}

					config.Shortcuts.Delete(shortcuts.Shortcut{Index: index})
					return config.Write()
				},
			},
			{
				Name:    "prompt-size",
				Aliases: []string{"p"},
				Usage:   "Set prompt size",
				Action: func(c *cli.Context) error {
					size, err := prompt.InputPrompt("Size")
					if err != nil {
						return err
					}

					err = config.SetPromptSize(size)
					if err != nil {
						return err
					}

					return config.Write()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil && !errors.Is(err, promptkit.ErrAborted) {
		log.Fatal(err)
	}
}
