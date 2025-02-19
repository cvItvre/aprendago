package cap13

import (
	"fmt"
)

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) toString() {
	fmt.Printf("Olá, sou %v %v e tenho %v anos.\n", p.nome, p.sobrenome, p.idade)
}

func ex04() {

	pessoa := Pessoa{"Pedro", "Araujo", 27}
	pessoa.toString()

}
