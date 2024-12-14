package hbxml

import (
	"fmt"
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
		if !strings.HasSuffix(file.Name(), ".hbxml") {
			continue
		}

		_, err = ParseFile(fmt.Sprintf("examples/%s", file.Name()))
		if err != nil {
			t.Error(err)
		}
	}
}

func contains(s string, substr string) bool {
	before, after, found := strings.Cut(s, substr)
	return found && before != after
}
