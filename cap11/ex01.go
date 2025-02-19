package cap11

import (
	"fmt"
)

type Pessoa struct {
	nome        string
	sobrenome   string
	favIceCream []string
}

func ex01() {

	pessoas := make([]Pessoa, 2, 4)

	pessoas[0].nome = "Pedro"
	pessoas[0].sobrenome = "Araujo"
	pessoas[0].favIceCream = []string{"Nata Goiaba", "Chocolate"}

	user := Pessoa{
		nome:      "Karol",
		sobrenome: "Souza",
		favIceCream: []string{
			"Ninho com Nutela",
			"Baunilha",
		},
	}

	pessoas[1] = user

	for _, pessoa := range pessoas {
		fmt.Printf("Nome: %v %v\t-\tSabores Favoritos: ", pessoa.nome, pessoa.sobrenome)
		for _, sabor := range pessoa.favIceCream {
			fmt.Printf("%v. ", sabor)
		}
		fmt.Println()
	}

}
