package hbxml

import (
	"encoding/xml"
	"io"
)

// Read a beatmap from an io.Reader
func NewBeatmap(r io.Reader) (beatmap *Beatmap, err error) {
	beatmap = new(Beatmap)
	err = xml.NewDecoder(r).Decode(beatmap)
	return beatmap, err
}
