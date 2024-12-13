package hbxml

import (
	"encoding/xml"
	"fmt"
	"io"
)

// Beatmap is a struct that represents the structure of a beatmap/hbxml file
type Beatmap struct {
	XMLName      xml.Name      `xml:"beatmap"`
	Version      string        `xml:"version,attr"`
	Meta         Meta          `xml:"meta"`
	General      General       `xml:"general"`
	Difficulty   Difficulty    `xml:"difficulty"`
	Editor       string        `xml:"editor"`
	Events       Events        `xml:"events"`
	TimingPoints []TimingPoint `xml:"timing-points>timing-point"`
	Colors       []Combo       `xml:"colors>combo"`
	HitObjects   []HitObject   `xml:"hit-objects>hit-object"`
}

// Serialize a beatmap to an io.Writer
func (b *Beatmap) Serialize(w io.Writer) error {
	return xml.NewEncoder(w).Encode(b)
}

// Deserialize a beatmap from an io.Reader
func (b *Beatmap) Deserialize(r io.Reader) error {
	return xml.NewDecoder(r).Decode(b)
}

// FormatName returns a formatted string of the beatmap's name
func (b *Beatmap) FormatName() string {
	return fmt.Sprintf("%s - %s [%s]", b.Meta.Artist, b.Meta.Title, b.Meta.Creator)
}

// Read a beatmap from an io.Reader
func NewBeatmap(r io.Reader) (beatmap *Beatmap, err error) {
	beatmap = new(Beatmap)
	err = xml.NewDecoder(r).Decode(beatmap)
	return beatmap, err
}
