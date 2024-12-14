package hbxml

// TimingPoint is a struct that represents the timing points of a beatmap
type TimingPoint struct {
	Inherited        bool    `xml:"inherited,attr"`
	Offset           int     `xml:"offset,attr"`
	BPM              float64 `xml:"bpm,attr"`
	SliderMultiplier float64 `xml:"sliderMultiplier,attr"`
	TimeSignature    int     `xml:"timeSignature,attr"`
	SampleSet        int     `xml:"sampleSet,attr"`
	CustomSampleSet  int     `xml:"customSampleSet,attr"`
	Volume           int     `xml:"volume,attr"`
	Special          bool    `xml:"special,attr"`
}

func (b *Beatmap) GetTimingPoint(offset int) *TimingPoint {
	for _, tp := range b.TimingPoints {
		if tp.Offset == offset {
			return &tp
		}
	}
	return nil
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
