package cap11

import (
	"fmt"
)

type Veiculo struct {
	portas int
	cor    string
}

type Caminhonete struct {
	Veiculo
	quatroRodas bool
}

type Sedan struct {
	Veiculo
	modeloLuxo bool
}

func ex03() {

	saveiro := Caminhonete{
		Veiculo: Veiculo{
			portas: 2,
			cor:    "prata",
		},
		quatroRodas: true,
	}

	seal := Sedan{
		Veiculo: Veiculo{
			portas: 4,
			cor:    "preto",
		},
		modeloLuxo: true,
	}

	fmt.Println(saveiro)
	fmt.Println(seal)
	fmt.Println(saveiro.portas)
	fmt.Println(seal.cor)

}
