package drawing

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Rectangle
type Rectangle struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func rectangleFromString(input string) (Rectangle, error) {

	r := strings.NewReader(input)
	rect := Rectangle{}
	err := json.NewDecoder(r).Decode(&rect)

	return rect, err
}

func rectangleToString(rect Rectangle) (string, error) {

	b, err := json.Marshal(rect)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func rectangleToSvg(rect Rectangle, transitionId int) string {

	svg := fmt.Sprintf(
		"<rect class=\"transition%d\" x=\"%f\" y=\"%f\" width=\"%f\" height=\"%f\" id=\"1\" stroke=\"black\" fill=\"transparent\" stroke-width=\"4\"></rect>\n",
		transitionId, rect.X, rect.Y, rect.Width, rect.Height)

	return svg
}
