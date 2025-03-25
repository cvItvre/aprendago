package main

import "testing"

func TestAverage(t *testing.T) {
	numbers := []float64{1, 2}
	v := Average(numbers)

	if v != 1.5 {
		t.Error("Expected 1.5, got", v)
	}
}

func TestAverageZeroValue(t *testing.T) {
	numbers := []float64{}
	v := Average(numbers)

	if v != 0 {
		t.Error("Expected 0, got", v)
	}
}

func TestMin(t *testing.T) {
	var numbers []float64
	for num := range 100 {
		numbers = append(numbers, float64(num))
	}
	min := Min(numbers)
	if min != 0 {
		t.Error("Expected 0, got", min)
	}
}

func TestMax(t *testing.T) {
	var numbers []float64
	for num := range 100 {
		numbers = append(numbers, float64(num))
	}
	max := Max(numbers)
	if max != 99 {
		t.Error("Expected 0, got", max)
	}
}
