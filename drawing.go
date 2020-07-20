package drawing

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Next Page
type NextPage struct {
	Name  string `json:"name"`
	Delay int    `json:"delay"`
}

// Drawing
type Drawing struct {
	Width       float64      `json:"width"`
	Height      float64      `json:"height"`
	RectWidth   float64      `json:"rectWidth"`
	RectHeight  float64      `json:"rectHeight"`
	Shapes      []Shape      `json:"shapes"`
	Connectors  []Connector  `json:"connectors"`
	Transitions []Transition `json:"transitions"`
	Next        NextPage     `json:"nextPage"`
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

func ToHtml(d Drawing, autoPlay bool) (string, error) {

	s, err := ToSvg(d)
	if err != nil {
		return s, err
	}

	h := "<!DOCTYPE html>\n" +
		"<html>\n" +
		"<head>\n"

	if d.Next.Name != "" && d.Next.Delay > 0 {
		h += "<script>\n" +
			"function nextPage() {\n" +
			"  location.replace(\"" + d.Next.Name + "\");\n" +
			"}\n"
		if autoPlay {
			h += "function autoPage() {\n" +
				"  setTimeout(function(){ nextPage() }, " + fmt.Sprintf("%d", d.Next.Delay) + "000);\n" +
				"}\n"
		}
		h += "</script>\n"
	}

	h += "</head>\n"
	h += "<body"

	if d.Next.Name != "" && d.Next.Delay > 0 {
		h += " onload=\"autoPage()\""
	}

	h += ">\n"
	h += s
	h += "</body>\n</html>\n"

	return h, err
}

func ToSvg(d Drawing) (string, error) {

	rects := ""
	i := 0
	for _, r := range d.Shapes {
		rects += shapeToSvg(r, i)
		i++
	}

	connectors := ""
	i = 1
	for _, c := range d.Connectors {
		connectors += connectorToSvg(d, c, i)
		i++
	}

	transitions := ""
	i = 0
	for _, t := range d.Transitions {
		transitions += transitionToSvg(t, i)
		i++
	}

	s := fmt.Sprintf(
		"<svg width=\"%f\" height=\"%f\" align=\"center\">"+
			"<rect x=\"0\" y=\"0\" width=\"20\" height=\"20\" stroke=\"black\" fill=\"transparent\" stroke-width=\"0\" onclick=\"nextPage()\"></rect>"+
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
		d.Width, d.Height, rects, connectors, transitions)

	return s, nil
}

func AddRectangle(drawing Drawing, x float64, y float64) (Drawing, error) {

	drawing.Shapes = append(drawing.Shapes, Shape{x, y, drawing.RectWidth, drawing.RectHeight, "rect", "", 0, "", "", 0, 0})
	return drawing, nil
}

func AddText(drawing Drawing, x float64, y float64, text string, size int) (Drawing, error) {

	drawing.Shapes = append(drawing.Shapes, Shape{x, y, drawing.RectWidth, drawing.RectHeight, "text", text, size, "", "", 0, 0})
	return drawing, nil
}

func AddConnector(drawing Drawing, c Connector) Drawing {

	drawing.Connectors = append(drawing.Connectors, c)
	return drawing
}

func ShapeCount(drawing Drawing) int {

	return len(drawing.Shapes)
}
