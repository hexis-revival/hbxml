package hbxml

// Meta is a struct that represents the metadata of a beatmap
type Meta struct {
	Title        string   `xml:"title,attr"`
	Artist       string   `xml:"artist,attr"`
	Creator      string   `xml:"creator,attr"`
	Version      string   `xml:"version,attr"`
	BaseVersion  string   `xml:"baseVersion,attr"`
	Source       string   `xml:"source,attr"`
	MapSetID     int      `xml:"mapSetId,attr"`
	MapVersionID int      `xml:"mapVersionId,attr"`
	Tags         []string `xml:"tags>tag"`
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

// Combo is a struct that represents the combo colors of a beatmap
type Combo struct {
	Red   int `xml:"red,attr"`
	Green int `xml:"green,attr"`
	Blue  int `xml:"blue,attr"`
}
