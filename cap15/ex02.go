package cap15

import (
	"fmt"
)

type Pessoa struct {
	nome  string
	idade int
}

func mudeMe(p *Pessoa) {
	p.nome = "Karol" // (*p).nome
	p.idade = 27     // (*p).idade
}

func ex02() {

	pessoa := Pessoa{"Pedro", 26}
	fmt.Println(pessoa)

	mudeMe(&pessoa)
	fmt.Println(pessoa)

}
