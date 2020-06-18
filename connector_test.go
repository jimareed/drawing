package drawing

import (
	"testing"
)

func TestConnectorFromString(t *testing.T) {

	c1 := Connector{1, 2}

	str, err := connectorToString(c1)
	if err != nil {
		t.Log("connectorToString error")
		t.Fail()
	}

	c2, err := connectorFromString(str)
	if err != nil {
		t.Log("connectorFromString error")
		t.Fail()
	}

	if c2.Shape1 != 1 || c2.Shape2 != 2 {
		t.Log("connector To/From string error")
		t.Fail()
	}

}

func TestConnectorBasics(t *testing.T) {
	d1 := Drawing{}
	d1.Width = 600
	d1.Height = 400
	d1.RectWidth = 30
	d1.RectHeight = 20

	d1, _ = AddRectangle(d1, 10, 10)
	d1, _ = AddRectangle(d1, 60, 10)

	c1 := Connector{0, 1}

	d1 = AddConnector(d1, c1)

	if connectorSlope(d1, c1) != 0 {
		t.Log("c1 slope error")
		t.Fail()
	}

	expectedP1 := Point{40, 20}
	p1 := connectorP1(d1, c1)

	if p1 != expectedP1 {
		t.Log("c1 p1 error")
		t.Log(p1)
		t.Fail()
	}

	d1, _ = AddRectangle(d1, 20, 20)
	d1, _ = AddRectangle(d1, 60, 40)

	c2 := Connector{2, 3}

	if connectorSlope(d1, c2) != 0.5 {
		t.Log("c2 slope error")
		t.Log(connectorSlope(d1, c2))
		t.Fail()
	}

}
