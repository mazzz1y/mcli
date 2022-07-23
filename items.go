package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
)

type Item struct {
	Name  string `json:"name"`
	Cmd   string `json:"cmd"`
	Index int    `json:"-"`
}

type Items []Item

func (is Items) add(i Item) error {
	is = append(is, i).sort()

	file, err := json.MarshalIndent(is, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFilePath, file, 0600)
}

func (is Items) delete(item Item) error {
	is[item.Index] = is[len(is)-1]
	is = is[:len(is)-1].sort()

	file, err := json.MarshalIndent(is, "", "  ")
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

func (is Items) sort() Items {
	sort.Slice(is, func(i, j int) bool {
		return is[i].Name < is[j].Name
	})

	return is
}
