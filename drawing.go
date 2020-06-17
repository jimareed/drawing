package drawing

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Drawing
type Drawing struct {
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	RectWidth  float64 `json:"rectWidth"`
	RectHeight float64 `json:"rectHeight"`
	Shapes     []Shape `json:"shapes"`
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

func ToSvg(d Drawing) (string, error) {

	rects := ""
	i := 0
	for _, r := range d.Shapes {
		rects += shapeToSvg(r, i)
	}

	connectors := ""
	transitions := ""

	s := fmt.Sprintf(
		"<svg width=\"%f\" height=\"%f\" align=\"center\">"+
			" <rect x=\"0\" y=\"0\" id=\"editor-canvas\" width=\"%f\" height=\"%f\" stroke=\"white\" fill=\"transparent\" stroke-width=\"0\"></rect>"+
			"%s\n"+
			"<defs>\n"+
			"<marker id=\"arrowhead\" markerWidth=\"5\" markerHeight=\"3.5\" refX=\"0\" refY=\"1.75\" orient=\"auto\">\n"+
			"    <polygon points=\"0 0, 5 1.75 0 3.5\"></polygon>\n"+
			"</marker>\n"+
			"</defs>\n"+
			"%s\n"+
			"<style>\n"+
			"%s\n"+
			"@keyframes transitionOpacity {\n"+
			"    0%%   { opacity: 0; }\n"+
			"    50%%   { opacity: 0; }\n"+
			"    100%% { opacity: 1; }\n"+
			"}\n"+
			"</style>\n"+
			"</svg>\n",
		d.Width, d.Height, d.Width, d.Height, rects, connectors, transitions)

	return s, nil
}

func AddRectangle(drawing Drawing, x float64, y float64) (Drawing, error) {

	drawing.Shapes = append(drawing.Shapes, Shape{x, y, drawing.RectWidth, drawing.RectHeight, "rect", "", 0})
	return drawing, nil
}

func RectangleCount(drawing Drawing) int {

	return len(drawing.Shapes)
}
