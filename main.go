package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cycleFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read file %q: %w", path, err)
	}

	if len(data) == 0 {
		// Empty file: nothing to do.
		return nil
	}

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file %q: %w", path, err)
	}

	if len(lines) <= 1 {
		// Single line (or only newlines): nothing to move.
		return nil
	}

	// Rotate: move first line to end.
	rotated := append(lines[1:], lines[0])

	var sb strings.Builder
	for _, l := range rotated {
		sb.WriteString(l)
		sb.WriteByte('\n')
	}

	if err := os.WriteFile(path, []byte(sb.String()), 0o644); err != nil {
		return fmt.Errorf("cannot write file %q: %w", path, err)
	}
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: cycle <file>\n")
		os.Exit(1)
	}

	path := os.Args[1]
	if err := cycleFile(path); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
