package cap11

import (
	"fmt"
)

type Person struct {
	nome        string
	sobrenome   string
	favIceCream []string
}

func ex02() {

	pessoa1 := Person{}
	pessoa1.nome = "Pedro"
	pessoa1.sobrenome = "Araujo"
	pessoa1.favIceCream = []string{"Nata Goiaba", "Chocolate"}

	pessoa2 := Person{
		nome:      "Karol",
		sobrenome: "Souza",
		favIceCream: []string{
			"Ninho com Nutela",
			"Baunilha",
		},
	}

	// pessoas := make(map[strin]Pessoa)
	pessoas := map[string]Person{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	fmt.Println(pessoas)

	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %v %v\t-\tSabores Favoritos: ", pessoa.nome, pessoa.sobrenome)
		for _, sabor := range pessoa.favIceCream {
			fmt.Printf("%v. ", sabor)
		}
		fmt.Println()
	}

}
