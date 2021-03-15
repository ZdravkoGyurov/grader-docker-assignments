package calc

import (
	"assignment1/calc"
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 3
	actual := assignment1.Add(1, 2)

	if expected != actual {
		t.Errorf("actual=%d, expected=%d", actual, expected)
	}
}
