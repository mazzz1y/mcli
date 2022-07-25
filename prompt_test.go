package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatch(t *testing.T) {
	name := []string{"one", "two", "three"}
	cmd := []string{"four", "five", "six"}

	item := Item{
		Name: strings.Join(name, " "),
		Cmd:  strings.Join(cmd, " "),
	}

	tt := []struct {
		item   Item
		input  string
		result bool
	}{
		{
			item:   item,
			input:  name[0],
			result: true,
		},
		{
			item:   item,
			input:  cmd[0],
			result: true,
		},
		{
			item:   item,
			input:  item.Cmd,
			result: true,
		},
		{
			item:   item,
			input:  fmt.Sprintf("%s %s", name[1], name[2]),
			result: true,
		},
		{
			item:   item,
			input:  fmt.Sprintf("%s %s", name[2], name[1]),
			result: true,
		},
		{
			item:   item,
			input:  fmt.Sprintf("%s %s", name[2][:2], cmd[1][:2]),
			result: true,
		},
		{
			item:   item,
			input:  fmt.Sprintf("%s %s %s", name[2], name[1], "false"),
			result: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			assert.Equal(t, tc.result, searchMatch(tc.item, tc.input))
		})
	}
}
