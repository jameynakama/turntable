package collections

import (
	"fmt"
	"io"
	"slices"
)

type Collection map[string][]Album

// Returns a new empty collection
func New() Collection {
	return make(map[string][]Album)
}

// Adds an album to the collection
func (c Collection) Add(album Album, artist string) {
	c[artist] = append(c[artist], album)
}

// Prints all stored albums, sorted by artist name
func (c Collection) ShowAll(out io.Writer) error {
	artists := make([]string, len(c))
	i := 0
	for a := range c {
		artists[i] = a
		i++
	}
	slices.Sort(artists)

	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}

	for _, artist := range artists {
		for _, album := range c[artist] {
			var played string
			if album.IsPlayed {
				played = "played"
			} else {
				played = "unplayed"
			}
			if _, err := fmt.Fprintf(out, "%q by %s (%s)\n", album.Name, artist, played); err != nil {
				return err
			}
		}
	}

	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}

	return nil
}

type Album struct {
	Name     string
	IsPlayed bool
}

// Sets album to "played"
func (a *Album) Play() {
	a.IsPlayed = true
}
