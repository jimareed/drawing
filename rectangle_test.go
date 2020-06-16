package drawing

import (
	"testing"
)

func TestRectangleFromString(t *testing.T) {

	rect1 := Rectangle{100, 80, 90, 60}

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

	if rect2.X != 100 || rect2.Y != 80 {
		t.Log("rectangle To/From string error")
		t.Fail()
	}
}
