package math

import (
	"testing"
)

func TestSum(t *testing.T) {
	x := []int{2, 3}
	actual := sum(x)
	expected := 5
	if actual != expected {
		t.Errorf("got %v\nwant%v", actual, expected)
	}
}

func TestSum2(t *testing.T) {
	x := []int{1, 2, 3, 4, 5}
	actual := sum(x)
	expected := 16
	if actual != expected {
		t.Errorf("got %v\nwant%v", actual, expected)
	}
}
