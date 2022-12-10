package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/dmirubtsov/mcli/pkg/shortcuts"
)

const (
	filePerm          = 0600
	dirPerm           = 0700
	defaultPromptSize = 10
	appDirName        = "mcli"
	appConfigFileName = "config.json"
)

type Config struct {
	PromptSize int                 `json:"promptSize"`
	Shortcuts  shortcuts.Shortcuts `json:"shortcuts,omitempty"`
}

var configFilePath string

func init() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	appDirPath := fmt.Sprintf("%s/%s", userConfigDir, appDirName)

	_, err = os.Stat(appDirPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(appDirPath, dirPerm)
	}

	if err != nil {
		log.Fatal(err)
	}

	configFilePath = fmt.Sprintf("%s/%s", appDirPath, appConfigFileName)
}

func (c *Config) SetPromptSize(size string) error {
	s, err := strconv.ParseInt(size, 10, 0)
	if err != nil {
		return err
	}

	c.PromptSize = int(s)
	return err
}

func (c *Config) Read() error {
	file, err := os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, filePerm)
	if err != nil {
		return err
	}

	cBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(cBytes) == 0 {
		return c.Write()
	}

	err = json.Unmarshal(cBytes, &c)
	if err != nil {
		return err
	}

	for i := range c.Shortcuts {
		c.Shortcuts[i].Index = i
	}

	return nil
}

func (c *Config) Write() error {
	if c.PromptSize == 0 {
		c.PromptSize = defaultPromptSize
	}

	file, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, file, filePerm)
}

func (c *Config) WriteDemo() error {
	c.Shortcuts = []shortcuts.Shortcut{
		{
			Name: "Demo Command 1",
			Cmd:  "echo test1",
		},
		{
			Name: "Demo Command 2",
			Cmd:  "echo test2",
		},
		{
			Name: "Demo Command 3",
			Cmd:  "echo test3",
		},
	}

	return c.Write()
}
