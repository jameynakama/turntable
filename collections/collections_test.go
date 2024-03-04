package collections

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShow(t *testing.T) {
	testCases := []struct {
		name       string
		subcommand string
		filter     string
		expected   string
		expErr     error
	}{
		{
			"All", "all", "", `
"Night Visions" by Imagine Dragons (unplayed)
"Evolve" by Imagine Dragons (played)
"Hotter than Hell" by KISS (unplayed)
"Curb" by Nickelback (played)

`, nil,
		},
		{
			"Unplayed", "unplayed", "", `
"Night Visions" by Imagine Dragons
"Hotter than Hell" by KISS

`, nil,
		},
		{
			"AllBy", "all by", "KISS", `
"Hotter than Hell" by KISS (unplayed)

`, nil,
		},
		{
			"AllBy-ErrWrongNumberOfArtists", "all by", "", "\n", ErrWrongNumberOfArtists,
		},
		{
			"UnplayedBy", "unplayed by", "Imagine Dragons", `
"Night Visions" by Imagine Dragons

`, nil,
		},
		{
			"UnplayedBy-ErrWrongNumberOfArtists", "unplayed by", "", "\n", ErrWrongNumberOfArtists,
		},
		{
			"ErrInvalidSubcommand", "boost mobile", "Imagine Dragons", "\n", ErrInvalidSubcommand,
		},
	}

	c := New()
	c["Imagine Dragons"] = []Album{
		{Name: "Night Visions", IsPlayed: false},
		{Name: "Evolve", IsPlayed: true},
	}
	c["Nickelback"] = []Album{
		{Name: "Curb", IsPlayed: true},
	}
	c["KISS"] = []Album{
		{Name: "Hotter than Hell", IsPlayed: false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer

			err := c.Show(&buf, tc.subcommand, tc.filter)
			if tc.expErr != nil && err == nil {
				t.Fatalf("Was expecting error %q but got none", tc.expErr)
			}
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}

			assert.Equal(t, tc.expected, buf.String())
		})
	}
}

func TestAdd(t *testing.T) {
	testCases := []struct {
		name     string
		toAdd    *Album
		expected string
		expErr   error
	}{
		{"GardenPath", nil, "Hoobastank", nil},
		{"GardenPath", &Album{Name: "Hoobastank"}, "Hoobastank", ErrAlbumAlreadyExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := New()
			err := c.Add(Album{Name: "Hoobastank", IsPlayed: false}, "Hoobastank")
			if err != nil {
				t.Fatal(err)
			}

			if tc.toAdd != nil {
				err = c.Add(*tc.toAdd, "Hoobastank")
			}

			if tc.expErr != nil && err == nil {
				t.Fatalf("Was expecting error %q but got none", tc.expErr)
			}
			if err != nil {
				assert.Equal(t, tc.expErr, err)
			}

			assert.Equal(t, tc.expected, c["Hoobastank"][0].Name)
		})
	}
}

func TestPlay(t *testing.T) {
	c := New()
	c.Add(Album{Name: "Hoobastank"}, "Hoobastank")
	assert.False(t, c["Hoobastank"][0].IsPlayed)
	c["Hoobastank"][0].Play()
	assert.True(t, c["Hoobastank"][0].IsPlayed)
}
