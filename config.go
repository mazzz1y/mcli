package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	filePerm          = 0600
	dirPerm           = 0700
	defaultPromptSize = 10
	appDirName        = "mcli"
	appConfigFileName = "config.json"
)

type Config struct {
	Version    string `json:"currentVersion"`
	PromptSize int    `json:"promptSize"`
	Items      Items  `json:"items,omitempty"`
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

func (c *Config) setPromptSize(size string) error {
	s, err := strconv.ParseInt(size, 10, 0)
	if err != nil {
		return err
	}

	c.PromptSize = int(s)
	return err
}

func (c *Config) read() error {
	file, err := os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, filePerm)
	if err != nil {
		return err
	}

	cBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(cBytes) == 0 {
		return c.write()
	}

	return json.Unmarshal(cBytes, &c)
}

func (c *Config) write() error {
	if c.PromptSize == 0 {
		c.PromptSize = defaultPromptSize
	}
	c.Version = version

	file, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configFilePath, file, filePerm)
}
