package main

import (
	"bytes"
	"fmt"
	"io"
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
			"add \"Hoobastank\" \"Hoobastank\"",
			[]album{{"Hoobastank", false}},
			nil,
		},
		{
			"ErrInvalidArgsOne",
			newCollection(),
			"add \"Hoobastank\"",
			[]album(nil),
			fmt.Errorf("Please provide 2 quoted args: an album and an artist"),
		},
		{
			"ErrInvalidArgsThree",
			newCollection(),
			"add \"Hoobastank\" \"Hoobastank\" \"Hoobastank\"",
			[]album(nil),
			fmt.Errorf("Please provide 2 quoted args: an album and an artist"),
		},
		{
			"ErrParsingARgs",
			newCollection(),
			"add \"Hoobastank",
			[]album(nil),
			fmt.Errorf("Err: parse error on line 1, column 16: extraneous or missing \" in quoted-field"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := add(io.Discard, tc.input, tc.albums)
			assert.Equal(t, tc.expected, tc.albums["Hoobastank"])
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
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
"Hotter than Hell" by KISS (unplayed)
"Curb" by Nickelback (played)

`, buf.String())
}
