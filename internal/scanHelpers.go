package internal

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/jameynakama/turntable/collections"
)

func PromptUser(in io.Reader, out io.Writer, albums collections.Collection) (string, []string, error) {
	fmt.Fprint(out, "> ")

	s := bufio.NewScanner(in)
	input, err := ScanString(s)
	if err != nil {
		return "", nil, err
	}

	var command = input
	if idx := strings.IndexByte(input, '"'); idx >= 0 {
		command = input[:idx]
	}

	args, err := GetArgs(input)
	if err != nil {
		return "", nil, err
	}

	return strings.TrimSpace(command), args, nil
}

func ScanString(s *bufio.Scanner) (string, error) {
	var input string
	if s.Scan() {
		input = strings.TrimSpace(s.Text())
	}
	err := s.Err()
	if err != nil {
		return "", err
	}
	return input, nil
}

func GetArgs(input string) ([]string, error) {
	var rawArgs string
	if idx := strings.IndexByte(input, '"'); idx >= 0 {
		rawArgs = input[idx:]
	}
	if rawArgs == "" {
		return nil, nil
	}

	r := csv.NewReader(strings.NewReader(rawArgs))
	r.Comma = ' '
	fields, err := r.Read()
	if err != nil {
		return nil, fmt.Errorf("Err: %v", err)
	}
	return fields, nil
}
