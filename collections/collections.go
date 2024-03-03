package collections

import (
	"errors"
	"fmt"
	"io"
	"slices"
)

var (
	ErrWrongNumberOfArtists = errors.New("Please provide 1 quoted artist name")
	ErrAlbumAlreadyExists   = errors.New("Album already exists")
)

type Collection map[string][]Album

// Returns a new empty collection
func New() Collection {
	return make(map[string][]Album)
}

func (c Collection) sortArtists() []string {
	artists := make([]string, len(c))
	i := 0
	for a := range c {
		artists[i] = a
		i++
	}
	slices.Sort(artists)
	return artists
}

// Adds an album to the collection
func (c Collection) Add(album Album, artist string) error {
	var albumNames []string
	for _, albums := range c {
		for _, album := range albums {
			albumNames = append(albumNames, album.Name)
		}
	}
	if slices.Index(albumNames, album.Name) != -1 {
		return ErrAlbumAlreadyExists
	}

	c[artist] = append(c[artist], album)

	return nil
}

// Prints stored albums according to query, sorted by artist name
func (c Collection) Show(out io.Writer, subcommand string, filter ...string) error {
	artists := c.sortArtists()

	if _, err := fmt.Fprintln(out); err != nil {
		return err
	}

	for _, artist := range artists {
		for _, album := range c[artist] {
			switch subcommand {
			case "all":
				album.show(out, artist, true)
			case "unplayed":
				if !album.IsPlayed {
					album.show(out, artist, false)
				}
			case "all by":
				// TODO: Handle no artist
				if len(filter) != 1 {
					return ErrWrongNumberOfArtists
				}
				if filter[0] == artist {
					album.show(out, artist, true)
				}
			case "unplayed by":
				// TODO: Handle no artist
				if len(filter) != 1 {
					return ErrWrongNumberOfArtists
				}
				if filter[0] == artist && !album.IsPlayed {
					album.show(out, artist, false)
				}
			default:
				// TODO: Unrecognized command
				// album.show(out, artist)
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

func (a *Album) show(out io.Writer, artist string, showPlayedStatus bool) error {
	var playedStatus string
	if showPlayedStatus {
		if a.IsPlayed {
			playedStatus = " (played)"
		} else {
			playedStatus = " (unplayed)"
		}
	}
	if _, err := fmt.Fprintf(out, "%q by %s%s\n", a.Name, artist, playedStatus); err != nil {
		return err
	}
	return nil
}

// Sets album to "played"
func (a *Album) Play() {
	a.IsPlayed = true
}
