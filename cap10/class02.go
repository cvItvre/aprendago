package cap10

import (
	"fmt"
)

type Person struct {
	nome  string
	idade int
}

type Professional struct {
	pessoa  Person
	titulo  string
	salario float64
}

func class02() {

	pessoa1 := Person{
		nome:  "Pedro",
		idade: 27,
	}

	pessoa2 := Professional{
		pessoa: Person{
			nome:  "Karol",
			idade: 28,
		},
		titulo:  "Data Analyst",
		salario: 7500.00,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

}
