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

type command struct {
	Name string `json:"name"`
	Cmd  string `json:"cmd"`
}

func main() {
	os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	app := cli.NewApp()
	app.Name = "mcli"
	app.Usage = "cmd shortcut menu"
	app.Action = func(c *cli.Context) error {
		index := selectCommand()
		subprocess(getCommands()[index].Cmd)
		return nil
	}
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add new command",
			Action: func(c *cli.Context) error {
				addCommand(command{
					prompt("Name"),
					prompt("Command"),
				})
				return nil
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Remove command",
			Action: func(c *cli.Context) error {
				i := selectCommand()
				delCommand(i)
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func prompt(label string) string {
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
	result, err := prompt.Run()

	if err != nil {
		log.Fatalln(err.Error())
		return ""
	}
	return result
}

func selectCommand() int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "> {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "> {{ .Cmd }}",
		Details:  `{{ "Command: " | faint }}{{ .Cmd }}`,
	}

	searcher := func(input string, index int) bool {
		command := getCommands()[index]
		name := strings.Replace(strings.ToLower(command.Name), " ", "", -1)
		cmd := strings.Replace(strings.ToLower(command.Cmd), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(cmd, input)
	}

	prompt := promptui.Select{
		Label:     "Select command:",
		Items:     getCommands(),
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}
	prompt.StartInSearchMode = true

	i, _, err := prompt.Run()

	if err != nil {
		log.Fatalln(err.Error())
		return 0
	}
	return i
}

func subprocess(command string) {
	cmdArray := strings.Fields(command)
	cmd := exec.Command(cmdArray[0], cmdArray[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	if err := cmd.Start(); nil != err {
		log.Fatalln(err.Error())
	}
	cmd.Wait()
}

func expandHomedir(dir string) string {
	dir, err := homedir.Expand(dir)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return dir
}

func getCommands() []command {
	var c []command
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalln(err.Error())
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &c)
	if err != nil {
		log.Fatalln("Emply list, please add commands by run `mcli add`")
	}
	return c
}

func addCommand(cmd command) {
	commandList := append(getCommands(), cmd)
	file, _ := json.Marshal(commandList)
	err := ioutil.WriteFile(configFilePath, file, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func delCommand(index int) {
	commandList := getCommands()
	commandList[index] = commandList[len(commandList)-1]
	commandList = commandList[:len(commandList)-1]

	file, _ := json.Marshal(commandList)
	err := ioutil.WriteFile(configFilePath, file, 0644)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
