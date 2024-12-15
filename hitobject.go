package hbxml

import (
	"math"
)

// HitObject is a struct that represents the hit objects of a beatmap
type HitObject struct {
	Type        HitObjectType  `xml:"type,attr"`
	StartOffset int            `xml:"offset,attr"`
	X           int            `xml:"x,attr"`
	Y           int            `xml:"y,attr"`
	NewCombo    bool           `xml:"newCombo,attr"`
	HitSound    HitSoundType   `xml:"hitsound,attr"`
	Curve       CurveType      `xml:"curve,attr,omitempty"`
	Length      float64        `xml:"length,attr,omitempty"`
	Backtracks  int            `xml:"backtracks,attr,omitempty"`
	EndOffset   int            `xml:"endOffset,attr,omitempty"`
	TickRate    int            `xml:"tickRate,attr,omitempty"`
	Points      []Point        `xml:"point"`
	Additions   []HitSoundType `xml:"hit-sound"`
}

// IsCircle returns true if the hit object is a circle
func (h *HitObject) IsCircle() bool {
	return h.Type == HitCircle
}

// IsSlider returns true if the hit object is a slider
func (h *HitObject) IsSlider() bool {
	return h.Type == Slider
}

// IsSpinner returns true if the hit object is a spinner
func (h *HitObject) IsSpinner() bool {
	return h.Type == Spinner
}

// IsHold returns true if the hit object is a hold
func (h *HitObject) IsHold() bool {
	return h.Type == Hold
}

func (h *HitObject) ClosestTimingPoint(points []TimingPoint) *TimingPoint {
	var closest *TimingPoint
	for _, tp := range points {
		if tp.Offset <= h.StartOffset {
			if closest == nil || tp.Offset > closest.Offset {
				closest = &tp
			}
		}
	}
	return closest
}

func (h *HitObject) CurrentBPM(b *Beatmap) float64 {
	var bpm float64
	for _, tp := range b.TimingPoints {
		if !tp.Inherited {
			bpm = tp.BPM
		}
		if tp.Offset <= h.StartOffset {
			break
		}
	}
	return bpm
}

func (h *HitObject) SliderMultiplier(b *Beatmap) float64 {
	timingPoint := h.ClosestTimingPoint(b.TimingPoints)
	sliderMultiplier := timingPoint.SliderMultiplier

	if sliderMultiplier == 0 {
		sliderMultiplier = 1.0
	}

	return b.Difficulty.SliderMultiplier * sliderMultiplier
}

func (h *HitObject) SliderMultiplierPerBeat(b *Beatmap) float64 {
	return h.SliderMultiplier(b) * 100
}

func (h *HitObject) TickLength(b *Beatmap) float64 {
	return h.SliderMultiplierPerBeat(b) / float64(b.Difficulty.SliderTickRate)
}

func (h *HitObject) TickPerSide(b *Beatmap) int {
	length := math.Floor(float64(h.Length) / h.TickLength(b) * 100)
	return int(math.Ceil(length/100) - 1)
}

func (h *HitObject) HoldLength() float64 {
	return float64(h.EndOffset-h.StartOffset) / 1000.0
}

// ComputeMaxCombo returns the maximum combo of the beatmap
func (b *Beatmap) ComputeMaxCombo() int {
	if len(b.TimingPoints) == 0 || len(b.HitObjects) == 0 {
		return 0
	}
	maxCombo := 0

	for _, h := range b.HitObjects {
		switch h.Type {
		default:
			maxCombo++
		case Slider:
			maxCombo += (h.Backtracks+1)*(h.TickPerSide(b)+1) + 1
		case Hold:
			beatsPerSecond := h.CurrentBPM(b) / 60.0
			combo := h.HoldLength() * beatsPerSecond * float64(h.TickRate)
			maxCombo += int(math.Ceil(combo))
		}
	}

	return maxCombo
}

// TotalLength returns the total length of the beatmap in seconds
func (b *Beatmap) TotalLength() float64 {
	return float64(b.HitObjects[len(b.HitObjects)-1].StartOffset) / 1000
}

// DrainLength returns the total length of the drain time in the beatmap in seconds
func (b *Beatmap) DrainLength() float64 {
	return b.TotalLength() - b.BreakLength()
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
