package internal

import (
	"bufio"
	"strings"
)

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
