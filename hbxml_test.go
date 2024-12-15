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

		_, err = ParseFromFile(fmt.Sprintf("examples/%s", file.Name()))
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestCombo(t *testing.T) {
	comboMap := map[string]int{
		"LeaF - NANO DEATH!!!!! (NerdNerdTewt) [Death].hbxml":                                   87,
		"Drop - Granat (TheHowl) [Pro].hbxml":                                                   214,
		"RhythmHolic - GLP Dubstep Intro (ProPlayer08) [Pro].hbxml":                             163,
		"RhythmHolic - GLP Dubstep Intro (ProPlayer08) [EZ].hbxml":                              132,
		"NJK Record - Dirty Sexy Girl (Nicholas) [BASIC].hbxml":                                 702,
		"NJK Record - Dirty Sexy Girl (Nicholas) [NO HOLDS].hbxml":                              548,
		"Hige Driver join. SELEN - Dadadadadadadadadada (Nicholas) [EXPERT].hbxml":              470,
		"Kuroma - Feelings beloved overheat (Takuya) [Expert].hbxml":                            1378,
		"shisotex - Final Fantasy III Sailing Enterprise ~ The Invincible (Takuya) [Pro].hbxml": 1198,
	}

	for file, combo := range comboMap {
		b, err := ParseFromFile("examples/" + file)
		if err != nil {
			t.Error(err)
			return
		}
		caclulatedCombo := b.ComputeMaxCombo()
		difference := combo - caclulatedCombo

		if difference != 0 {
			t.Errorf(
				"WARNING: Expected %d, got %d (%d): %s\n",
				combo, caclulatedCombo, difference, b.FormatName(),
			)
		}
	}
}

func contains(s string, substr string) bool {
	before, after, found := strings.Cut(s, substr)
	return found && before != after
}
