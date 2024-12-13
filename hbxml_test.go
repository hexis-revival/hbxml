package hbxml

import (
	"os"
	"strings"
	"testing"
)

func TestBeatmapDecoding(t *testing.T) {
	exampleFiles, err := os.ReadDir("examples")
	if err != nil {
		t.Error(err)
	}

	for _, file := range exampleFiles {
		file, err := os.Open("examples/" + file.Name())
		if err != nil {
			t.Error(err)
		}

		_, err = NewBeatmap(file)
		if err != nil {
			t.Error(err)
		}
	}
}

func contains(s string, substr string) bool {
	before, after, found := strings.Cut(s, substr)
	return found && before != after
}
