package cap10

import "fmt"

type Cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func class01() {

	cliente1 := Cliente{
		nome:      "Pedro",
		sobrenome: "Araujo",
		fumante:   true,
	}

	cliente2 := Cliente{"Karol", "Souza", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

}
