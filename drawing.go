package drawing

// single slide
type Drawing struct {
	width       float64
	height      float64
	blockWidth  float64
	blockHeight float64
	rects       []Rectangle
}

func FromString(input string) (Drawing, error) {

	drawing := Drawing{}

	return drawing, nil
}

func ToSvg(drawing Drawing) (string, error) {

	return "Hello World!", nil
}

func AddRectangle(drawing Drawing, x float64, y float64) (Drawing, error) {

	drawing.rects = append(drawing.rects, Rectangle{x, y})
	return drawing, nil
}

func RectangleCount(drawing Drawing) int {

	return len(drawing.rects)
}
