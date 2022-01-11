package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestEqual[T comparable](t *testing.T, expect, result T, message ...string) {
	t.Helper()
	var msg string
	if len(message) > 0 {
		msg = message[0]
	}
	t.Run(msg, func(t *testing.T) {
		t.Helper()
		if result != expect {
			t.Fatalf("expected %v, got %v", expect, result)
		}
	})
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Atoi(input string) int {
	i := Must(strconv.Atoi(input))
	return i
}

const inputDir = "."

func inputPath(day int, suffix ...string) string {
	var s string
	if len(suffix) > 0 {
		s = strings.Join(suffix, "")
	}
	return Must(filepath.Abs(fmt.Sprintf("%s/day%02d%s.txt", inputDir, day, s)))
}

func Input(day int, suffix ...string) string {
	data := Must(os.ReadFile(inputPath(day, suffix...)))
	return strings.TrimSpace(string(data))
}

func InputReader(day int, suffix ...string) io.ReadSeekCloser {
	return Must(os.Open(inputPath(day, suffix...)))
}

func InputScanner(day int, suffix ...string) (s *bufio.Scanner, close func() error) {
	f := InputReader(day, suffix...)
	s = bufio.NewScanner(f)
	return s, f.Close
}

func Lines(day int, suffix ...string) []string {
	b, close := InputScanner(day, suffix...)
	defer close()
	var lines []string
	for b.Scan() {
		lines = append(lines, b.Text())
	}
	return lines
}
