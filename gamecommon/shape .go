package gamecommon

type Shape struct {
	X, Y     int
	geometry any
}

type Rectangle struct {
	Width, Height int
}

type Circle struct {
	Radius int
}

// x and y represent the top-left corner for Rectangle
func NewRectangle(x, y, width, height int) Shape {
	return Shape{
		X: x,
		Y: y,
		geometry: Rectangle{
			Width:  width,
			Height: height,
		},
	}
}

// x and y represent the center for Circle
func NewCircle(x, y, radius int) Shape {
	return Shape{
		X: x,
		Y: y,
		geometry: Circle{
			Radius: radius,
		},
	}
}

func (s Shape) Contains(px, py int) bool {
	switch g := s.geometry.(type) {
	case Rectangle:
		return px >= s.X && px <= s.X+g.Width && py >= s.Y && py <= s.Y+g.Height
	case Circle:
		dx := px - s.X
		dy := py - s.Y
		return dx*dx+dy*dy <= g.Radius*g.Radius
	default:
		return false
	}
}
