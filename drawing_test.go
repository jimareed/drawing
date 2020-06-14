package drawing

import (
	"testing"
)

func TestToString(t *testing.T) {

	d1 := Drawing{}

	d1, err := AddRectangle(d1, 40, 30)
	if err != nil {
		t.Log("AddRectangle error")
		t.Fail()
	}

	s, err := ToString(d1)
	if err != nil {
		t.Log("ToString error")
		t.Fail()
	}

	d2, err := FromString(s)
	if err != nil {
		t.Log("FromString error")
		t.Fail()
	}

	if RectangleCount(d2) != RectangleCount(d1) {
		t.Log("To/FromString error")
		t.Fail()
	}
}

func TestToSvg(t *testing.T) {

	d, err := AddRectangle(Drawing{}, 40, 30)
	if err != nil {
		t.Log("AddRectangle error")
		t.Fail()
	}
	_, err = ToSvg(d)
	if err != nil {
		t.Log("ToSvg error")
		t.Fail()
	}
}
