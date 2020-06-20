package drawing

import (
	"testing"
)

func TestTransitionFromString(t *testing.T) {

	t1 := Transition{15}

	str, err := transitionToString(t1)
	if err != nil {
		t.Log("transitionToString error")
		t.Fail()
	}

	t2, err := transitionFromString(str)
	if err != nil {
		t.Log("transitionFromString error")
		t.Fail()
	}

	if t1.Duration != 15 || t2.Duration != 15 {
		t.Log("transition To/From string error")
		t.Fail()
	}
}
