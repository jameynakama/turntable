package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/jameynakama/turntable/collections"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		albums   collections.Collection
		input    []string
		expected []collections.Album
		expErr   error
	}{
		{
			"AddsAlbum",
			collections.New(),
			[]string{"Hoobastank", "Hoobastank"},
			[]collections.Album{{Name: "Hoobastank", IsPlayed: false}},
			nil,
		},
		{
			"ErrWrongNumberOfArgsOne",
			collections.New(),
			[]string{"Hoobastank"},
			[]collections.Album(nil),
			ErrWrongNumberOfArgs(2, "an album and an artist"),
		},
		{
			"ErrWrongNumberOfArgsThree",
			collections.New(),
			[]string{"Hoobastank", "Hoobastank", "Hoobastank"},
			[]collections.Album(nil),
			ErrWrongNumberOfArgs(2, "an album and an artist"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := add(io.Discard, tc.input, tc.albums)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}
			assert.Equal(t, tc.expected, tc.albums["Hoobastank"])
		})
	}
}

func TestShow(t *testing.T) {
	albums := collections.New()
	albums["Imagine Dragons"] = []collections.Album{
		{Name: "Night Visions", IsPlayed: false},
		{Name: "Evolve", IsPlayed: true},
	}
	albums["Nickelback"] = []collections.Album{
		{Name: "Curb", IsPlayed: true},
	}
	albums["KISS"] = []collections.Album{
		{Name: "Hotter than Hell", IsPlayed: false},
	}

	testCases := []struct {
		name     string
		fn       CollectionCmd
		filter   string
		expected string
	}{
		{
			"All", showAll, "", `
"Night Visions" by Imagine Dragons (unplayed)
"Evolve" by Imagine Dragons (played)
"Hotter than Hell" by KISS (unplayed)
"Curb" by Nickelback (played)

`,
		},
		{
			"Unplayed", showUnplayed, "", `
"Night Visions" by Imagine Dragons
"Hotter than Hell" by KISS

`,
		},
		{
			"By", showAllBy, "Imagine Dragons", `
"Night Visions" by Imagine Dragons (unplayed)
"Evolve" by Imagine Dragons (played)

`,
		},
		{
			"UnplayedBy", showUnplayedBy, "Imagine Dragons", `
"Night Visions" by Imagine Dragons

`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tc.fn(&buf, nil, albums, tc.filter)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expected, buf.String())
		})
	}

}

func TestPlay(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected bool
		expErr   error
	}{
		{
			"SetsPlayed",
			[]string{"Hoobastank"},
			true,
			nil,
		},
		{
			"ErrAlbumNotFound",
			[]string{"Love Gun"},
			false,
			ErrAlbumNotFound("Love Gun"),
		},
		{
			"ErrWrongNumberOfArgsZero",
			nil,
			false,
			ErrWrongNumberOfArgs(1, "an album name"),
		},
		{
			"ErrWrongNumberOfArgsTwo",
			[]string{"Love Gun", "KISS"},
			false,
			ErrWrongNumberOfArgs(1, "an album name"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			albums := collections.New()
			albums["Hoobastank"] = []collections.Album{
				{Name: "Hoobastank", IsPlayed: false},
			}

			err := play(io.Discard, tc.input, albums)
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}

			assert.Equal(t, tc.expected, albums["Hoobastank"][0].IsPlayed)
		})
	}
}
