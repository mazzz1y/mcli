package main

import (
	"errors"
	"github.com/dmirubtsov/mcli/pkg/items"
	"github.com/dmirubtsov/mcli/pkg/prompt"
	"github.com/dmirubtsov/mcli/pkg/subprocess"
	"log"
	"os"

	"github.com/dmirubtsov/mcli/pkg/config"
	"github.com/erikgeiser/promptkit"
	"github.com/urfave/cli/v2"
)

var Version = "git"

func main() {
	var config = config.Config{}
	if err := config.Read(); err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:    "mcli",
		Version: Version,
		Usage:   "Shell command shortcut menu",
		Action: func(c *cli.Context) error {
			index, err := prompt.SelectionPrompt(config.Items, config.PromptSize)
			if err != nil {
				return err
			}

			return subprocess.Exec(config.Items[index].Cmd)
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add item",
				Action: func(c *cli.Context) error {
					nameField, err := prompt.InputPrompt("Name")
					if err != nil {
						return err
					}

					commandField, err := prompt.InputPrompt("Command")
					if err != nil {
						return err
					}

					config.Items.Add(items.Item{Name: nameField, Cmd: commandField})
					return config.Write()
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "Edit item",
				Action: func(c *cli.Context) error {
					index, err := prompt.SelectionPrompt(config.Items, config.PromptSize)
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

					config.Items.Delete(items.Item{Index: index})
					config.Items.Add(items.Item{Name: nameField, Cmd: commandField})
					return config.Write()
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Remove item",
				Action: func(c *cli.Context) error {
					index, err := prompt.SelectionPrompt(config.Items, config.PromptSize)
					if err != nil {
						return err
					}

					config.Items.Delete(items.Item{Index: index})
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
