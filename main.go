package main

import (
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var configFilePath = expandHomedir("~/.mcli.json")
var version = "git"

func main() {
	items, err := Items{}.get()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "mcli"
	app.Usage = "Shell command shortcut menu"
	app.Version = version

	app = &cli.App{
		Action: func(c *cli.Context) error {
			index, err := selectItem(items)
			if err != nil {
				return err
			}

			return subprocess(items[index].Cmd)
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

					return items.add(Item{Name: nameField, Cmd: commandField})
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "Remove item",
				Action: func(c *cli.Context) error {
					index, err := selectItem(items)
					if err != nil {
						return err
					}

					return items.delete(Item{Index: index})
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func expandHomedir(dir string) string {
	dir, err := homedir.Expand(dir)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
