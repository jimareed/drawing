package drawing

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Shape
type Transition struct {
	Duration int `json:"duration"`
}

func transitionFromString(input string) (Transition, error) {

	r := strings.NewReader(input)
	t := Transition{}
	err := json.NewDecoder(r).Decode(&t)

	return t, err
}

func transitionToString(t Transition) (string, error) {

	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func transitionToSvg(t Transition, transitionId int) string {

	svg := ""

	svg += fmt.Sprintf(
		".transition%d {"+
			"	animation-name: transitionOpacity;"+
			"	animation-duration: %ds;"+
			"	animation-iteration-count: 1;"+
			"}", transitionId, t.Duration)

	return svg
}
