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
