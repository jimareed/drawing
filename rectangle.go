package drawing

import (
	"encoding/json"
	"strings"
)

// Rectangle
type Rectangle struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
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

func rectangleToSvg(rect Rectangle) (string, error) {

	return "Hello World!", nil
}
