package main

import (
	"errors"
	"log"
	"os"

	"github.com/dmirubtsov/mcli/internal/config"
	"github.com/erikgeiser/promptkit"
	"github.com/urfave/cli/v2"
)

var version = defaultVersionText
var cfg = config.Config{}

func init() {
	cfg = config.Config{}
	if err := cfg.Read(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := &cli.App{
		Name:    "mcli",
		Version: version,
		Usage:   usageText,
		Action:  run,
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   addShortcutText,
				Action:  add,
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   editShortcutText,
				Action:  edit,
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   removeShortcutText,
				Action:  delete,
			},
			{
				Name:    "prompt-size",
				Aliases: []string{"p"},
				Usage:   setPromptSizeText,
				Action:  setPromptSize,
			},
		},
	}

	if err := app.Run(os.Args); err != nil && !errors.Is(err, promptkit.ErrAborted) {
		log.Fatal(err)
	}
}
