package drawing

import (
)

// single slide
type Drawing struct {
	title string
}

func FromString(input string) (Drawing, error) {

	drawing := Drawing{}

    return drawing, nil
}

func ToSvg(drawing Drawing) (string, error) {

    return "Hello World!", nil
}
