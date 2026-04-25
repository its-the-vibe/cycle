package main

import (
	"os"
	"path/filepath"
	"testing"
)

func writeFile(t *testing.T, dir, name, content string) string {
	t.Helper()
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writeFile: %v", err)
	}
	return path
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("readFile: %v", err)
	}
	return string(data)
}

func TestCycleFile_Normal(t *testing.T) {
	dir := t.TempDir()
	path := writeFile(t, dir, "test.txt", "line1\nline2\nline3\n")

	if err := cycleFile(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := readFile(t, path)
	want := "line2\nline3\nline1\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCycleFile_TwoLines(t *testing.T) {
	dir := t.TempDir()
	path := writeFile(t, dir, "test.txt", "first\nsecond\n")

	if err := cycleFile(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := readFile(t, path)
	want := "second\nfirst\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCycleFile_SingleLine(t *testing.T) {
	dir := t.TempDir()
	original := "onlyone\n"
	path := writeFile(t, dir, "test.txt", original)

	if err := cycleFile(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := readFile(t, path)
	if got != original {
		t.Errorf("single-line file should be unchanged: got %q, want %q", got, original)
	}
}

func TestCycleFile_Empty(t *testing.T) {
	dir := t.TempDir()
	path := writeFile(t, dir, "test.txt", "")

	if err := cycleFile(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := readFile(t, path)
	if got != "" {
		t.Errorf("empty file should remain empty: got %q", got)
	}
}

func TestCycleFile_NoTrailingNewline(t *testing.T) {
	dir := t.TempDir()
	// File with no trailing newline on the last line.
	path := writeFile(t, dir, "test.txt", "line1\nline2\nline3")

	if err := cycleFile(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	got := readFile(t, path)
	want := "line2\nline3\nline1\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCycleFile_NonexistentFile(t *testing.T) {
	err := cycleFile("/nonexistent/path/file.txt")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

func TestCycleFile_MultipleRotations(t *testing.T) {
	dir := t.TempDir()
	path := writeFile(t, dir, "test.txt", "a\nb\nc\n")

	for i := 0; i < 3; i++ {
		if err := cycleFile(path); err != nil {
			t.Fatalf("rotation %d: unexpected error: %v", i, err)
		}
	}

	// After 3 full rotations the file should be back to original.
	got := readFile(t, path)
	want := "a\nb\nc\n"
	if got != want {
		t.Errorf("after 3 rotations got %q, want %q", got, want)
	}
}
