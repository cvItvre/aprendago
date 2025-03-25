package main

import "testing"

type Test struct {
	data   []int
	answer int
}

func TestSoma(t *testing.T) {
	var result = Soma(1, 2, 3, 4)
	if result != 10 {
		t.Errorf("Expected: %v. Got: %v", 10, result)
	}
}

func TestSomaTabela(t *testing.T) {
	var tests = []Test{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 1}, 2},
		{[]int{10, 20, 30}, 60},
		{[]int{-10, 10}, 0},
		{[]int{}, 0},
	}

	var result int

	for _, test := range tests {
		result = Soma(test.data...)
		if result != test.answer {
			t.Errorf("Expected %v. Got %v.", test.answer, result)
		}
	}
}

func TestMultiply(t *testing.T) {
	var result = multiply(1, 2, 3, 4)
	if result != 24 {
		t.Errorf("Expected %v. Got %v", 24, result)
	}
}

func TestMultiplyTabela(t *testing.T) {
	var tests = []Test{
		{[]int{1, 1}, 1},
		{[]int{}, 0},
		{[]int{2, 2, 2, 2, 2}, 32},
		{[]int{28, 36, 4, 0, 17, 44}, 0},
	}

	var result int

	for _, test := range tests {
		result = multiply(test.data...)
		if result != test.answer {
			t.Errorf("Expected %v. Got %v", test.answer, result)
		}
	}
}
