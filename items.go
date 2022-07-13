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

func (is Items) get(file os.File) (Items, error) {
	byteValue, err := ioutil.ReadAll(&file)
	if err != nil {
		return is, err
	}

	if len(byteValue) == 0 {
		emptyJson, err := json.Marshal(is)
		if err != nil {
			return is, err
		}

		return is, ioutil.WriteFile(configFilePath, emptyJson, 0600)
	}

	return is, json.Unmarshal(byteValue, &is)
}
