package hbxml

type HitObjectType int

const (
	HitCircle HitObjectType = 1 << iota
	Slider    HitObjectType = 1 << iota
	Spinner   HitObjectType = 1 << iota
	Hold      HitObjectType = 1 << iota
)

type HitSoundType int

const (
	Normal  HitSoundType = 1 << iota
	Whistle HitSoundType = 1 << iota
	Finish  HitSoundType = 1 << iota
	Clap    HitSoundType = 1 << iota
)

type CurveType string

const (
	LinearCurve      CurveType = "l"
	CatmullCurve     CurveType = "c"
	BezierCurve      CurveType = "b"
	PassThroughCurve CurveType = "p"
)
