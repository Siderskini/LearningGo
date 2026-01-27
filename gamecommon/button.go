package gamecommon

type Button struct {
	Name  string
	shape Shape
	input *Input
	Text  string
}

func NewRectangleButton(name string, x, y, width, height int, text string) Button {
	return Button{
		Name:  name,
		shape: NewRectangle(x, y, width, height),
		Text:  text,
	}
}

func NewCircleButton(name string, x, y, radius int, text string) Button {
	return Button{
		Name:  name,
		shape: NewCircle(x, y, radius),
		Text:  text,
	}
}

func IsPressed(button Button, input *Input) bool {
	pressedLocation, ok := input.PressedLocation()
	if !ok {
		return false
	}
	return button.shape.Contains(pressedLocation[0], pressedLocation[1])
}
