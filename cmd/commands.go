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

type CollectionCmd func(io.Writer, []string, collections.Collection, ...string) error

// Adds an album to the collection. The first quoted string is the album title,
// and the second is the band name.
func add(out io.Writer, args []string, albums collections.Collection, filters ...string) error {
	if len(args) != 2 {
		return ErrWrongNumberOfArgs(2, "an album and an artist")
	}

	artist := args[1]
	album := collections.Album{Name: args[0], IsPlayed: false}

	if err := albums.Add(album, artist); err != nil {
		return err
	}

	fmt.Fprintf(out, "\nAdded %q by %s\n\n", album.Name, artist)

	return nil
}

func showAll(out io.Writer, args []string, albums collections.Collection, filters ...string) error {
	if err := albums.Show(out, "all"); err != nil {
		return err
	}
	return nil
}

func showUnplayed(out io.Writer, args []string, albums collections.Collection, filters ...string) error {
	if err := albums.Show(out, "unplayed"); err != nil {
		return err
	}
	return nil
}

func showAllBy(out io.Writer, args []string, albums collections.Collection, filters ...string) error {
	if err := albums.Show(out, "all by", filters...); err != nil {
		return err
	}
	return nil
}

func showUnplayedBy(out io.Writer, args []string, albums collections.Collection, filters ...string) error {
	if err := albums.Show(out, "unplayed by", filters...); err != nil {
		return err
	}
	return nil
}

func play(out io.Writer, args []string, albums collections.Collection, filters ...string) error {
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
	fmt.Fprintf(out, "\nBye!\n\n")
}
