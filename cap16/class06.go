package cap16

import (
	"fmt"
	"sort"
)

type Carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenarPorPotencia []Carro

func (x ordenarPorPotencia) Len() int { return len(x) }

func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }

func (x ordenarPorPotencia) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type ordenarPorConsumo []Carro

func (x ordenarPorConsumo) Len() int { return len(x) }

func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }

func (x ordenarPorConsumo) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type ordenarPorLucroProDonoDoPosto []Carro

func (x ordenarPorLucroProDonoDoPosto) Len() int { return len(x) }

func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }

func (x ordenarPorLucroProDonoDoPosto) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func class06() {

	carros := []Carro{
		{"Panamera", 650, 6},
		{"Passati", 300, 16},
		{"Song+", 330, 40},
		{"Ex90", 350, 45},
		{"Celta", 80, 30},
	}

	fmt.Println(carros)

	sort.Sort(ordenarPorPotencia(carros))
	fmt.Println(carros)

	sort.Sort(ordenarPorConsumo(carros))
	fmt.Println(carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
	fmt.Println(carros)

}
