package cap10

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade int
}

type Profissional struct {
	Pessoa  //apenas nome da variável / struct
	titulo  string
	salario float64
}

func class03() {

	pessoa1 := Pessoa{
		nome:  "Pedro",
		idade: 27,
	}

	pessoa2 := Profissional{
		Pessoa: Pessoa{
			nome:  "Karol",
			idade: 28,
		},
		titulo:  "Data Analyst",
		salario: 7500.00,
	}

	fmt.Println(pessoa1.idade) // how to access
	fmt.Println(pessoa2.Pessoa.nome)
	fmt.Println(pessoa2.idade) // promoted field

	pessoa3 := Profissional{Pessoa{"Ellen", 34}, "Dev Senior", 14000.53}
	fmt.Println(pessoa3)

	// https://go.dev/ref/spec#Struct_types

}
