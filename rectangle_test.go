package drawing

import (
	"testing"
)

func TestRectangleFromString(t *testing.T) {

	rect1 := Rectangle{60, 40}

	str, err := rectangleToString(rect1)
	if err != nil {
		t.Log("rectangleToString error")
		t.Fail()
	}

	rect2, err := rectangleFromString(str)
	if err != nil {
		t.Log("rectangleFromString error")
		t.Fail()
	}

	if rect2.X != 60 || rect2.Y != 40 {
		t.Log("rectangle To/From string error")
		t.Fail()
	}
}
