package main

import "fmt"

func main() {

	var numbers []float64
	for num := range 100 {
		numbers = append(numbers, float64(num))
	}

	fmt.Println("Average:", Average(numbers))
	fmt.Println("Min:", Min(numbers))
	fmt.Println("Max:", Max(numbers))

}

func Average(numbers []float64) float64 {
	var sum float64
	if len(numbers) == 0 {
		return 0
	}
	for _, v := range numbers {
		sum += v
	}
	return sum / float64(len(numbers))
}

func Min(numbers []float64) float64 {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func Max(numbers []float64) float64 {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}
