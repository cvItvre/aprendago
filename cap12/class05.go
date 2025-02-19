package cap12

import (
	"fmt"
)

type Person struct {
	nome      string
	sobrenome string
	idade     int
}

type Dentista struct {
	Person
	dentesArrancados int
	salario          float64
}

type Arquiteto struct {
	Person
	tipoConstrucao string
	tamanhoLoucura string
}

func (d Dentista) oibomdia() {
	fmt.Println("Meu nome é", d.nome, "e eu ja arranquei", d.dentesArrancados, "e ouve só: Bom dia!")
}

func (a Arquiteto) oibomdia() {
	fmt.Println("Meu nome é", a.nome, "e ouve só: Bom dia!")
}

type Gente interface { //interfaces
	oibomdia()
}

func serHumano(g Gente) { //polimorfismo

	g.oibomdia()

	switch g.(type) {
	case Dentista:
		fmt.Println("Sou Rico? ->", g.(Dentista).salario)
	case Arquiteto:
		fmt.Println("Minha Loucura -> ", g.(Arquiteto).tamanhoLoucura)
	}

}

func class05() {

	mrpredio := Arquiteto{
		Person: Person{
			nome:      "Seu Dente",
			sobrenome: "Bonitao",
			idade:     50,
		},
		tipoConstrucao: "Hospícios",
		tamanhoLoucura: "Imensa",
	}
	mrdente := Dentista{
		Person: Person{
			nome:      "Seu Dente",
			sobrenome: "Amarelado",
			idade:     65,
		},
		dentesArrancados: 347,
		salario:          19000,
	}

	mrdente.oibomdia()
	mrpredio.oibomdia()

	serHumano(mrpredio)
	serHumano(mrdente)

}
