package main

import "testing"

func TestSoma(t *testing.T) {
	var result = Soma(1, 2, 3, 4)
	if result != 10 {
		t.Errorf("Expected: %v. Got: %v", 10, result)
	}
}

func TestMultiply(t *testing.T) {
	var result = multiply(1, 2, 3, 4)
	if result != 24 {
		t.Errorf("Expected %v. Got %v", 24, result)
	}
}
