package drawing

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Shape
type Shape struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
	Type   string  `json:"type"`
	Desc   string  `json:"desc"`
	Size   int     `json:"size"`
	Style  string  `json:"style"`
	Slide  string  `json:"slide"`
}

func shapeFromString(input string) (Shape, error) {

	r := strings.NewReader(input)
	rect := Shape{}
	err := json.NewDecoder(r).Decode(&rect)

	return rect, err
}

func shapeToString(rect Shape) (string, error) {

	b, err := json.Marshal(rect)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func shapeToSvg(rect Shape, transitionId int) string {

	svg := ""

	if rect.Type == "text" {
		svg += fmt.Sprintf(
			"<text class=\"transition%d\" x=\"%f\" y=\"%f\" fill=\"black\" font-size=\"%dpx\">%s</text>\n",
			transitionId, rect.X, rect.Y, rect.Size, rect.Desc)
	} else {
		strokeWidth := 4
		if rect.Style == "hidden" {
			strokeWidth = 0
		}
		onClick := ""
		if rect.Slide != "" {
			onClick = fmt.Sprintf("onclick=\"location.href='%s'\"", rect.Slide)
		}
		svg += fmt.Sprintf(
			"<rect class=\"transition%d\" x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" id=\"1\" stroke=\"black\" fill=\"transparent\" stroke-width=\"%d\" %s\"></rect>\n",
			transitionId, rect.X, rect.Y, rect.Width, rect.Height, strokeWidth, onClick)
	}

	return svg
}
