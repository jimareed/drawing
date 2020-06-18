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
