package hbxml

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ParseFromFile parses a beatmap file from the given filepath.
func ParseFromFile(filepath string) (*Beatmap, error) {
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

	return ParseFromBytes(bytes)
}

// ParseFromBytes parses a beatmap file from the given byte slice.
func ParseFromBytes(bytes []byte) (*Beatmap, error) {
	return ParseFromString(string(bytes))
}

// ParseFromString parses a beatmap file from the given string.
func ParseFromString(str string) (*Beatmap, error) {
	return NewBeatmap(strings.NewReader(str))
}
