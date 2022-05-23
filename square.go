package square

import "math"

type Point struct {
	x, y int
}

type Square struct {
	start Point
	side  uint
}

func (s *Square) End() Point {
	// implement me

	var endPoint Point
	endPoint.x = s.start.x + int(s.side)
	endPoint.y = s.start.y + int(s.side)

	return endPoint

}

func (s *Square) Area() uint {
	return uint(math.Pow(float64(s.side), 2))
}
func (s *Square) Perimeter() uint {
	return s.side * 4
}
