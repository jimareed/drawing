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
