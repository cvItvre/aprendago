package cap12

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade int
}

func (p Pessoa) bomdia() { // metodos
	fmt.Println("Olá,", p.nome, "bom dia!")
}

func class04() {

	pessoa := Pessoa{"Pedro", 27}
	pessoa.bomdia()

}
