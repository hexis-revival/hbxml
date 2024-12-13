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

		beatmap, err := NewBeatmap(file)
		if err != nil {
			t.Error(err)
		}

		if !contains(file.Name(), beatmap.Meta.Title) {
			t.Errorf("Title not found in %s", file.Name())
		}

		if !contains(file.Name(), beatmap.Meta.Artist) {
			t.Errorf("Artist not found in %s", file.Name())
		}
	}
}

func contains(s string, substr string) bool {
	before, after, found := strings.Cut(s, substr)
	return found && before != after
}
