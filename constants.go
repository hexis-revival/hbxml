package hbxml

type HitObjectType int

const (
	HitCircle HitObjectType = 1 << iota
	Slider    HitObjectType = 1 << iota
	Spinner   HitObjectType = 1 << iota
	Hold      HitObjectType = 1 << iota
)
