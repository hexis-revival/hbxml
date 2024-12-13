package hbxml

// Meta is a struct that represents the metadata of a beatmap
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

// General is a struct that represents the general settings of a beatmap
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

// Difficulty is a struct that represents the difficulty settings of a beatmap
type Difficulty struct {
	HPDrainRate       float64 `xml:"hpDrainRate,attr"`
	CircleSize        float64 `xml:"circleSize,attr"`
	OverallDifficulty float64 `xml:"overallDifficulty,attr"`
	ApproachRate      float64 `xml:"approachRate,attr"`
	SliderMultiplier  float64 `xml:"sliderMultiplier,attr"`
	SliderTickRate    float64 `xml:"sliderTickRate,attr"`
}

// Events is a struct that represents the events of a beatmap.
type Events struct {
	Backgrounds []Background `xml:"backgrounds>background"`
	Breaks      []Break      `xml:"breaks>break"`
}

// Background is a struct that represents the background settings of a beatmap
type Background struct {
	Offset   int    `xml:"offset,attr"`
	Filename string `xml:"filename,attr"`
}

// Break is a struct that represents the break settings of a beatmap
type Break struct {
	Offset    int `xml:"offset,attr"`
	EndOffset int `xml:"endOffset,attr"`
}

// TimingPoint is a struct that represents the timing points of a beatmap
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

// Combo is a struct that represents the combo colors of a beatmap
type Combo struct {
	Red   int `xml:"red,attr"`
	Green int `xml:"green,attr"`
	Blue  int `xml:"blue,attr"`
}

// HitObject is a struct that represents the hit objects of a beatmap
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

// Point is a struct that represents the points of a hit object
type Point struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}
