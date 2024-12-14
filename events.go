package hbxml

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

// BreakLength returns the total length of breaks in the beatmap in seconds
func (b *Beatmap) BreakLength() float64 {
	var length float64
	for _, brk := range b.Events.Breaks {
		length += float64(brk.EndOffset-brk.Offset) / 1000
	}
	return length
}
