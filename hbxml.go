package hbxml

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ParseFile parses a beatmap file from the given filepath.
func ParseFile(filepath string) (*Beatmap, error) {
	if stat, err := os.Stat(filepath); err != nil || stat.IsDir() {
		return nil, fmt.Errorf("hbxml: %s is not a valid file", filepath)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("hbxml: %s", err)
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("hbxml: %s", err)
	}

	return ParseBytes(bytes)
}

// ParseBytes parses a beatmap file from the given byte slice.
func ParseBytes(bytes []byte) (*Beatmap, error) {
	return ParseString(string(bytes))
}

// ParseString parses a beatmap file from the given string.
func ParseString(str string) (*Beatmap, error) {
	return NewBeatmap(strings.NewReader(str))
}
