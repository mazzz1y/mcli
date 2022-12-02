package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var version = "git"
var config = Config{}

func main() {
	if err := config.read(); err != nil {
		log.Fatal(err)
	}

	app := &cli.App{
		Name:    "mcli",
		Version: version,
		Usage:   "Shell command shortcut menu",
		Action: func(c *cli.Context) error {
			index, err := selectItem(config.Items, config.PromptSize)
			if err != nil {
				return err
			}

			return subprocess(config.Items[index].Cmd)
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add item",
				Action: func(c *cli.Context) error {
					nameField, err := prompt("Name")
					if err != nil {
						return err
					}

					commandField, err := prompt("Command")
					if err != nil {
						return err
					}

					config.Items.add(Item{Name: nameField, Cmd: commandField})
					return config.write()
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "Edit item",
				Action: func(c *cli.Context) error {
					index, err := selectItem(config.Items, config.PromptSize)
					if err != nil {
						return err
					}

					nameField, err := prompt("Name")
					if err != nil {
						return err
					}

					commandField, err := prompt("Command")
					if err != nil {
						return err
					}

					config.Items.delete(Item{Index: index})
					config.Items.add(Item{Name: nameField, Cmd: commandField})
					return config.write()
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Remove item",
				Action: func(c *cli.Context) error {
					index, err := selectItem(config.Items, config.PromptSize)
					if err != nil {
						return err
					}

					config.Items.delete(Item{Index: index})
					return config.write()
				},
			},
			{
				Name:    "prompt-size",
				Aliases: []string{"p"},
				Usage:   "Set prompt size",
				Action: func(c *cli.Context) error {
					size, err := prompt("Size")
					if err != nil {
						return err
					}

					err = config.setPromptSize(size)
					if err != nil {
						return err
					}

					return config.write()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil && err.Error() != "^C" {
		log.Fatal(err)
	}
}
