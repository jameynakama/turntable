package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strings"
)

// TODO: Use "out" arg, not stdout

type collectionCmd func(io.Writer, string, collection) error

// Adds an album to the collection. The first quoted string is the album title,
// and the second is the band name.
func add(out io.Writer, input string, albums collection) error {
	r := csv.NewReader(strings.NewReader(input))
	r.Comma = ' '
	fields, err := r.Read()
	if err != nil {
		return fmt.Errorf("Err: %v", err)
	}
	if len(fields) != 3 {
		return fmt.Errorf("Please provide 2 quoted args: an album and an artist")
	}

	artist := fields[2]
	album := album{fields[1], false}

	albums[fields[2]] = append(albums[artist], album)

	fmt.Fprintf(out, "\nAdded %q by %s\n\n", album.name, artist)

	return nil
}

// Prints all stored albums, sorted by artist name
func showAll(out io.Writer, input string, albums collection) error {
	artists := make([]string, len(albums))
	i := 0
	for a := range albums {
		artists[i] = a
		i++
	}
	slices.Sort(artists)

	fmt.Fprintln(out)
	for _, artist := range artists {
		for _, song := range albums[artist] {
			var played string
			if song.isPlayed {
				played = "played"
			} else {
				played = "unplayed"
			}
			fmt.Fprintf(out, "%q by %s (%s)\n", song.name, artist, played)
		}
	}
	fmt.Fprintln(out)
	return nil
}

// STUB
func play(out io.Writer, input string, albums collection) error {
	fmt.Fprintln(out, "You chose \"play\"")
	return nil
}

func quit(out io.Writer) {
	fmt.Fprintln(out, "\nBye!")
}
