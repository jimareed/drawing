package drawing

import (
	"testing"
)

func TestFromString(t *testing.T) {

	_, err := FromString("test")
	if err != nil {
		t.Log("Error testing stubbed implementation")
		t.Fail()
	}
}

func TestToSvg(t *testing.T) {

	drawing, err := FromString("test")
	if err != nil {
		t.Log("Error testing stubbed implementation")
		t.Fail()
	}
	_, err = ToSvg(drawing)
	if err != nil {
		t.Log("Error testing stubbed implementation")
		t.Fail()
	}
}

func TestAddRectangle(t *testing.T) {

	d, err := FromString("test")

	d, err = AddRectangle(d, 20, 40)
	if err != nil {
		t.Log("Error adding rectangle")
		t.Fail()
	}

	count := RectangleCount(d)
	if count != 1 {
		t.Log("Error invalid rectangle count")
		t.Fail()
	}
}
