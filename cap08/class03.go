package cap08

import (
	"fmt"
)

func class03() {
	slice := []string{"banana", "maça", "jaca"}

	for k, v := range slice {
		fmt.Println("No indice", k, "temos o valor:", v)
	}

	slice = append(slice, "melancia")

	for k, v := range slice {
		fmt.Println("No indice", k, "tem o valor:", v)
	}

	number_collection := []int{3, 6, 12, 7, 3, 1, 73, 1}
	total := 0

	for _, v := range number_collection {
		total += v
	}

	fmt.Println("O valor total é:", total)

}
