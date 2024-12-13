package hbxml

import "encoding/xml"

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

type Meta struct {
	Title        string `xml:"title,attr"`
	Artist       string `xml:"artist,attr"`
	Creator      string `xml:"creator,attr"`
	Version      string `xml:"version,attr"`
	BaseVersion  string `xml:"baseVersion,attr"`
	Source       string `xml:"source,attr"`
	MapSetID     string `xml:"mapSetId,attr"`
	MapVersionID string `xml:"mapVersionId,attr"`
}

type General struct {
	AudioFilename         string  `xml:"audioFilename,attr"`
	AudioLeadIn           int     `xml:"audioLeadIn,attr"`
	PreviewOffset         int     `xml:"previewOffset,attr"`
	Countdown             int     `xml:"countdown,attr"`
	SampleSet             string  `xml:"sampleSet,attr"`
	StackLeniency         float64 `xml:"stackLeniency,attr"`
	Mode                  int     `xml:"mode,attr"`
	LetterboxDuringBreaks bool    `xml:"letterboxDuringBreaks,attr"`
}

type Difficulty struct {
	HPDrainRate       float64 `xml:"hpDrainRate,attr"`
	CircleSize        float64 `xml:"circleSize,attr"`
	OverallDifficulty float64 `xml:"overallDifficulty,attr"`
	ApproachRate      float64 `xml:"approachRate,attr"`
	SliderMultiplier  float64 `xml:"sliderMultiplier,attr"`
	SliderTickRate    float64 `xml:"sliderTickRate,attr"`
}

type Events struct {
	Backgrounds []Background `xml:"backgrounds>background"`
	Breaks      []Break      `xml:"breaks>break"`
}

type Background struct {
	Offset   int    `xml:"offset,attr"`
	Filename string `xml:"filename,attr"`
}

type Break struct {
	Offset    int `xml:"offset,attr"`
	EndOffset int `xml:"endOffset,attr"`
}

type TimingPoint struct {
	Inherited       bool    `xml:"inherited,attr"`
	Offset          int     `xml:"offset,attr"`
	BPM             float64 `xml:"bpm,attr"`
	TimeSignature   int     `xml:"timeSignature,attr"`
	SampleSet       int     `xml:"sampleSet,attr"`
	CustomSampleSet int     `xml:"customSampleSet,attr"`
	Volume          int     `xml:"volume,attr"`
	Special         bool    `xml:"special,attr"`
}

type Combo struct {
	Red   int `xml:"red,attr"`
	Green int `xml:"green,attr"`
	Blue  int `xml:"blue,attr"`
}

type HitObject struct {
	Type       int      `xml:"type,attr"`
	Offset     int      `xml:"offset,attr"`
	X          int      `xml:"x,attr"`
	Y          int      `xml:"y,attr"`
	NewCombo   bool     `xml:"newCombo,attr"`
	HitSound   string   `xml:"hitsound,attr"`
	Curve      string   `xml:"curve,attr,omitempty"`
	Length     float64  `xml:"length,attr,omitempty"`
	Backtracks int      `xml:"backtracks,attr,omitempty"`
	EndOffset  int      `xml:"endOffset,attr,omitempty"`
	TickRate   int      `xml:"tickRate,attr,omitempty"`
	Points     []Point  `xml:"point"`
	HitSounds  []string `xml:"hit-sound"`
}

type Point struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}
