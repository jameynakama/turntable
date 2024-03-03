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
	}{
		{
			"All", "all", "", `
"Night Visions" by Imagine Dragons (unplayed)
"Evolve" by Imagine Dragons (played)
"Hotter than Hell" by KISS (unplayed)
"Curb" by Nickelback (played)

`,
		},
		{
			"Unplayed", "unplayed", "", `
"Night Visions" by Imagine Dragons
"Hotter than Hell" by KISS

`,
		},
		{
			"AllBy", "all by", "KISS", `
"Hotter than Hell" by KISS (unplayed)

`,
		},
		{
			"UnplayedBy", "unplayed by", "Imagine Dragons", `
"Night Visions" by Imagine Dragons

`,
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
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tc.expected, buf.String())
		})
	}
}
