package calc_test

import (
	"node-express-react/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 3
	actual := calc.Add(1, 2)

	if expected != actual {
		t.Errorf("actual=%d, expected=%d", actual, expected)
	}
}

func TestAdd2(t *testing.T) {
	expected := 9
	actual := calc.Add(4, 4)

	if expected != actual {
		t.Errorf("actual=%d, expected=%d", actual, expected)
	}
}
