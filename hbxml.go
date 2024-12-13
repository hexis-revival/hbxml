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

// TotalCircles returns the total number of circles in the beatmap
func (b *Beatmap) TotalCircles() int {
	var circles int
	for _, obj := range b.HitObjects {
		if obj.IsCircle() {
			circles++
		}
	}
	return circles
}

// TotalSliders returns the total number of sliders in the beatmap
func (b *Beatmap) TotalSliders() int {
	var sliders int
	for _, obj := range b.HitObjects {
		if obj.IsSlider() {
			sliders++
		}
	}
	return sliders
}

// TotalSpinners returns the total number of spinners in the beatmap
func (b *Beatmap) TotalSpinners() int {
	var spinners int
	for _, obj := range b.HitObjects {
		if obj.IsSpinner() {
			spinners++
		}
	}
	return spinners
}

func (b *Beatmap) TotalHolds() int {
	var holds int
	for _, obj := range b.HitObjects {
		if obj.IsHold() {
			holds++
		}
	}
	return holds
}

// TotalObjects returns the total number of hit objects in the beatmap
func (b *Beatmap) TotalObjects() int {
	return len(b.HitObjects)
}

// BPMs returns a slice of all the BPMs in the beatmap
func (b *Beatmap) BPMs() []float64 {
	bpms := make([]float64, 0)
	for _, tp := range b.TimingPoints {
		if !tp.Inherited {
			bpms = append(bpms, tp.BPM)
		}
	}
	return bpms
}

// MedianBPM returns the median BPM of the beatmap
func (b *Beatmap) MedianBPM() float64 {
	bpms := b.BPMs()
	if len(bpms) == 0 {
		return 0
	}
	return bpms[len(bpms)/2]
}

// AverageBPM returns the average BPM of the beatmap
func (b *Beatmap) AverageBPM() float64 {
	bpms := b.BPMs()
	if len(bpms) == 0 {
		return 0
	}
	var sum float64
	for _, bpm := range bpms {
		sum += bpm
	}
	return sum / float64(len(bpms))
}

// HighestBPM returns the highest BPM of the beatmap
func (b *Beatmap) HighestBPM() float64 {
	bpms := b.BPMs()
	if len(bpms) == 0 {
		return 0
	}
	highest := bpms[0]
	for _, bpm := range bpms {
		if bpm > highest {
			highest = bpm
		}
	}
	return highest
}

// LowestBPM returns the lowest BPM of the beatmap
func (b *Beatmap) LowestBPM() float64 {
	bpms := b.BPMs()
	if len(bpms) == 0 {
		return 0
	}
	lowest := bpms[0]
	for _, bpm := range bpms {
		if bpm < lowest {
			lowest = bpm
		}
	}
	return lowest
}

// Read a beatmap from an io.Reader
func NewBeatmap(r io.Reader) (beatmap *Beatmap, err error) {
	beatmap = new(Beatmap)
	err = xml.NewDecoder(r).Decode(beatmap)
	return beatmap, err
}
