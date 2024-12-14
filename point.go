package hbxml

// Point is a struct that represents the points of a hit object
type Point struct {
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
}
