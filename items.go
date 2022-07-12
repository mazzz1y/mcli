package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Item struct {
	Name  string `json:"Name"`
	Cmd   string `json:"Cmd"`
	Index int    `json:"-"`
}

type Items []Item

func (is Items) add(i Item) error {
	is = append(is, i)
	file, err := json.Marshal(is)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath, file, 0600)
}

func (is Items) delete(item Item) error {
	is[item.Index] = is[len(is)-1]
	is = is[:len(is)-1]

	file, err := json.Marshal(is)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath, file, 0600)
}

func (is Items) get() (Items, error) {
	jsonFile, err := os.Open(configFilePath)
	if err != nil {
		return is, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return is, err
	}

	return is, json.Unmarshal(byteValue, &is)
}
