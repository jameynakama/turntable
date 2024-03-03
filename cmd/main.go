package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jameynakama/turntable/internal"
)

type album struct {
	name     string
	isPlayed bool
}

type collection map[string][]album

type config struct {
	albums collection
	in     io.Reader
	out    io.Writer
}

func newCollection() collection {
	return make(map[string][]album)
}

func main() {
	var albums = newCollection()
	var cfg = config{
		albums: albums,
		in:     os.Stdin,
		out:    os.Stdout,
	}
	err := run(cfg)
	if err != nil {
		fmt.Printf("Encountered error: %s\n", err)
		os.Exit(1)
	}
}

func run(cfg config) error {
	_, err := fmt.Fprint(cfg.out, "\nWelcome to your music collection!\n\n")
	if err != nil {
		return err
	}

	for {
		input, command, err := promptUser(cfg.in, cfg.out, cfg.albums)

		var commandFn collectionCmd
		switch command {
		case "add":
			commandFn = add
		case "play":
			commandFn = play
		case "show all":
			commandFn = showAll
		case "quit":
			quit(cfg.out)
			return nil
		default:
			fmt.Fprintf(cfg.out, "\n%q is not a valid command\n\n", input)
			continue
		}

		err = commandFn(cfg.out, input, cfg.albums)
		if err != nil {
			fmt.Fprintf(cfg.out, "\nError processing command: %s\n\n", err.Error())
			continue
		}
	}
}

func promptUser(in io.Reader, out io.Writer, albums collection) (string, string, error) {
	fmt.Fprint(out, "> ")

	s := bufio.NewScanner(in)
	input, err := internal.ScanString(s)
	if err != nil {
		return "", "", err
	}

	var command = input
	if idx := strings.IndexByte(input, '"'); idx >= 0 {
		command = input[:idx]
	}

	return input, strings.TrimSpace(command), nil
}
