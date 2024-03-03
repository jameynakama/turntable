package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArgs(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			"GardenPath",
			"play \"This\" \"Song\"",
			[]string{"This", "Song"},
		},
		{
			"NoArgs",
			"play",
			nil,
		},
		{
			"NoInput",
			"",
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			args, err := GetArgs(tc.input)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.expected, args)
		})
	}
}
