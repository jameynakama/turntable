package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		albums   collection
		input    string
		expected []album
		expErr   error
	}{
		{
			"AddsAlbum",
			newCollection(),
			"add \"Moose\" \"Moose Men\"",
			[]album{{"Moose", false}},
			nil,
		},
		{
			"ErrInvalidArgsOne",
			newCollection(),
			"add \"Moose\"",
			[]album(nil),
			fmt.Errorf("Please provide 2 quoted args: an album and an artist"),
		},
		{
			"ErrInvalidArgsThree",
			newCollection(),
			"add \"Moose\" \"Moose Men\" \"Bogus\"",
			[]album(nil),
			fmt.Errorf("Please provide 2 quoted args: an album and an artist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// err := add(tc.input, tc.albums)
			// assert.Equal(t, tc.expected, tc. albums["Moose Men"])
			// if err != nil {
			// 	assert.Equal(t, tc.expErr, err)
			// }
		})
	}
}

func TestShowAll(t *testing.T) {
	var buf bytes.Buffer

	albums := newCollection()
	albums["Imagine Dragons"] = []album{
		{"Night Visions", false},
		{"Evolve", true},
	}
	albums["Nickelback"] = []album{
		{"Curb", true},
	}
	albums["KISS"] = []album{
		{"Hotter than Hell", false},
	}

	err := showAll(&buf, "", albums)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, `
"Night Visions" by Imagine Dragons (unplayed)
"Evolve" by Imagine Dragons (played)
"Curb" by Nickelback (played)
"Hotter than Hell" by KISS (unplayed)

`, buf.String())
}
