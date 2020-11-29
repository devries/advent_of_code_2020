package utils

// 2-D Point
type Point struct {
	X int
	Y int
}

// +Y direction
var North = Point{0, 1}

// -Y direction
var South = Point{0, -1}

// +X direction
var East = Point{1, 0}

// -X direction
var West = Point{-1, 0}

// Slice of cardinal directions
var Directions = []Point{North, South, East, West}

func (p Point) Add(p2 Point) Point {
	r := Point{p.X + p2.X, p.Y + p2.Y}

	return r
}
