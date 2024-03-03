package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jameynakama/turntable/collections"
	"github.com/jameynakama/turntable/internal"
)

type config struct {
	albums collections.Collection
	in     io.Reader
	out    io.Writer
}

func main() {
	var albums = collections.New()
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
		command, args, err := promptUser(cfg.in, cfg.out, cfg.albums)

		var commandFn CollectionCmd
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
			fmt.Fprintf(cfg.out, "\n%q is not a valid command\n\n", command)
			continue
		}

		err = commandFn(cfg.out, args, cfg.albums)
		if err != nil {
			fmt.Fprintf(cfg.out, "\nError processing command: %s\n\n", err.Error())
			continue
		}
	}
}

func promptUser(in io.Reader, out io.Writer, albums collections.Collection) (string, []string, error) {
	fmt.Fprint(out, "> ")

	s := bufio.NewScanner(in)
	input, err := internal.ScanString(s)
	if err != nil {
		return "", nil, err
	}

	var command = input
	if idx := strings.IndexByte(input, '"'); idx >= 0 {
		command = input[:idx]
	}

	args, err := getArgs(input)
	if err != nil {
		return "", nil, err
	}

	return strings.TrimSpace(command), args, nil
}

func getArgs(input string) ([]string, error) {
	r := csv.NewReader(strings.NewReader(input))
	r.Comma = ' '
	fields, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}
	return fields, nil
}
