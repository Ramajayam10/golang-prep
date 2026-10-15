package unittesting

import "testing"

func TestAdd(t *testing.T) {
	a := 2
	b := 3
	expected := 5

	result := add(a, b)

	if result != expected {
		t.Errorf("Test failed: Expected %d, got %d", expected, result)
	}
}


func TestMultiply(t *testing.T){
	a := 2
	b := 3
	c := 4

	expected := 24 

	result := multiply(a,b,c)

	if result != expected{
		t.Errorf("Test failed: Expected %d, got %d",expected, result)
	}
}