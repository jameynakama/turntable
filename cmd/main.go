package main

import (
	"fmt"
	"io"
	"os"

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
		command, args, err := internal.PromptUser(cfg.in, cfg.out, cfg.albums)

		var commandFn CollectionCmd
		switch command {
		case "add":
			commandFn = add
		case "play":
			commandFn = play
		case "show all":
			commandFn = showAll
		case "show unplayed":
			commandFn = showUnplayed
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
