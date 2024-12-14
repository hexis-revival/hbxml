package hbxml

// HitObject is a struct that represents the hit objects of a beatmap
type HitObject struct {
	Type        HitObjectType  `xml:"type,attr"`
	StartOffset int            `xml:"offset,attr"`
	X           int            `xml:"x,attr"`
	Y           int            `xml:"y,attr"`
	NewCombo    bool           `xml:"newCombo,attr"`
	HitSound    HitSoundType   `xml:"hitsound,attr"`
	Curve       string         `xml:"curve,attr,omitempty"`
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

func (b *Beatmap) StartPosition() Point {
	if len(b.HitObjects) == 0 {
		return Point{}
	}
	return b.HitObjects[0].Points[0]
}

func (b *Beatmap) EndPosition() Point {
	if len(b.HitObjects) == 0 {
		return Point{}
	}
	return b.HitObjects[len(b.HitObjects)-1].Points[0]
}
