package drawing

import (
	"encoding/json"
	"strings"
)

// Drawing
type Drawing struct {
	Width      float64     `json:"width"`
	Height     float64     `json:"height"`
	RectWidth  float64     `json:"rectWidth"`
	RectHeight float64     `json:"rectHeight"`
	Rects      []Rectangle `json:"rects"`
}

func FromString(input string) (Drawing, error) {

	r := strings.NewReader(input)
	drawing := Drawing{}
	err := json.NewDecoder(r).Decode(&drawing)

	return drawing, err
}

func ToString(d Drawing) (string, error) {

	b, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func ToSvg(drawing Drawing) (string, error) {

	return "Hello World!", nil
}

func AddRectangle(drawing Drawing, x float64, y float64) (Drawing, error) {

	drawing.Rects = append(drawing.Rects, Rectangle{x, y})
	return drawing, nil
}

func RectangleCount(drawing Drawing) int {

	return len(drawing.Rects)
}
