package main

import "fmt"

type Humanos interface {
	falar()
}

type Pessoa struct {
	nome  string
	idade int
}

func (p *Pessoa) falar() {
	fmt.Println("Olá, me chamo", p.nome, "e tenho", p.idade, "anos.")
}

func dizerAlgumaCoisa(h Humanos) {
	h.falar()
}

func main() {

	p1 := &Pessoa{nome: "Pedro", idade: 27}
	p2 := Pessoa{nome: "Karol", idade: 28}

	dizerAlgumaCoisa(p1)
	dizerAlgumaCoisa(&p2)

	pessoas := []Pessoa{
		Pessoa{nome: "Joao", idade: 21},
		Pessoa{nome: "Luiz", idade: 23},
	}

	for _, v := range pessoas {
		dizerAlgumaCoisa(&v)
	}

}
