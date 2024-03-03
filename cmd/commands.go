package main

import (
	"fmt"
	"io"

	"github.com/jameynakama/turntable/collections"
)

func ErrAlbumNotFound(albumName string) error {
	return fmt.Errorf("Album %q not found", albumName)
}

func ErrWrongNumberOfArgs(n int, s string) error {
	return fmt.Errorf("Please provide %d quoted args: %s", n, s)
}

type CollectionCmd func(io.Writer, []string, collections.Collection) error

// Adds an album to the collection. The first quoted string is the album title,
// and the second is the band name.
func add(out io.Writer, args []string, albums collections.Collection) error {
	if len(args) != 2 {
		return ErrWrongNumberOfArgs(2, "an album and an artist")
	}

	artist := args[1]
	album := collections.Album{Name: args[0], IsPlayed: false}

	albums.Add(album, artist)

	fmt.Fprintf(out, "\nAdded %q by %s\n\n", album.Name, artist)

	return nil
}

func showAll(out io.Writer, args []string, albums collections.Collection) error {
	if err := albums.Show(out, collections.ALL); err != nil {
		return err
	}
	return nil
}

func showUnplayed(out io.Writer, args []string, albums collections.Collection) error {
	if err := albums.Show(out, collections.UNPLAYED); err != nil {
		return err
	}
	return nil
}

func play(out io.Writer, args []string, albums collections.Collection) error {
	if len(args) != 1 {
		return ErrWrongNumberOfArgs(1, "an album name")
	}

	for _, as := range albums {
		for i := range as {
			if as[i].Name == args[0] {
				as[i].Play()
				fmt.Fprintf(out, "\nYou're listening to %q\n\n", as[i].Name)
				return nil
			}
		}
	}

	return ErrAlbumNotFound(args[0])
}

func quit(out io.Writer) {
	fmt.Fprintln(out, "\nBye!")
}
