package main

import "fmt"

func main() {

	fmt.Println(Soma(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(Multiply(1, 2, 3, 4, 5))

}

func Soma(numbers ...int) int {
	var sum = 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Multiply(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	var result = 1
	for _, num := range numbers {
		result *= num
	}
	return result
}
