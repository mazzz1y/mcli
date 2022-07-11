package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

var configFilePath = expandHomedir("~/.mcli.json")
var version = "git"

type command struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
}

func main() {
	os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	app := cli.NewApp()
	app.Name = "mcli"
	app.Usage = "cmd shortcut menu"
	app.Version = version
	app.Action = func(c *cli.Context) error {
		index, err := selectCommand()
		if err != nil {
			return err
		}

		commands, err := getCommands()
		if err != nil {
			return err
		}

		return subprocess(commands[index].Cmd)
	}
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add new command",
			Action: func(c *cli.Context) error {
				nameField, err := prompt("Name")
				if err != nil {
					return err
				}

				commandField, err := prompt("Command")
				if err != nil {
					return err
				}

				return addCommand(command{
					nameField,
					commandField,
				})
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Remove command",
			Action: func(c *cli.Context) error {
				index, err := selectCommand()
				if err != nil {
					return err
				}

				return delCommand(index)
			},
		},
	}

	_ = app.Run(os.Args)

}

func prompt(label string) (string, error) {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | green }} ",
	}

	validate := func(input string) error {
		if input != "" {
			return nil
		}
		return errors.New(label + " cannot be empty")
	}

	prompt := promptui.Prompt{
		Label:     label + ":",
		Validate:  validate,
		Templates: templates,
	}
	return prompt.Run()
}

func selectCommand() (int, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "> {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "> {{ .Cmd }}",
	}

	commands, err := getCommands()
	if err != nil {
		return 0, err
	}

	searcher := func(input string, index int) bool {
		command := commands[index]
		name := strings.Replace(strings.ToLower(command.Name), " ", "", -1)
		cmd := strings.Replace(strings.ToLower(command.Cmd), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(cmd, input)
	}

	prompt := promptui.Select{
		Label:     "Select command:",
		Items:     commands,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}
	prompt.StartInSearchMode = true

	i, _, err := prompt.Run()

	if err != nil {
		return 0, err
	}

	return i, nil
}

func subprocess(command string) error {
	commandArr := strings.Fields(command)
	cmd := exec.Command(commandArr[0], commandArr[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); nil != err {
		return err
	}

	return cmd.Wait()
}

func expandHomedir(dir string) string {
	dir, err := homedir.Expand(dir)
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func getCommands() ([]command, error) {
	var c []command
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		return []command{}, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return []command{}, err
	}

	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		return []command{}, err
	}

	return c, err
}

func addCommand(cmd command) error {
	commands, err := getCommands()
	if err != nil {
		return err
	}

	commandList := append(commands, cmd)
	file, err := json.Marshal(commandList)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath, file, 0644)
}

func delCommand(index int) error {
	commandList, err := getCommands()
	if err != nil {
		return err
	}

	commandList[index] = commandList[len(commandList)-1]
	commandList = commandList[:len(commandList)-1]

	file, err := json.Marshal(commandList)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath, file, 0644)
}
