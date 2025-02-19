package cap16

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Nome      string
	Sobrenome string
	Idade     int
	Profissao string
	AccBanco  float64
}

func clas02() {

	pedro := Person{"Pedro", "Araujo", 26, "Sec Engineer", 30754.33}
	karol := Person{
		Nome:      "Karol",
		Sobrenome: "Sousa",
		Idade:     27,
		Profissao: "Data Analyst",
		AccBanco:  345821.21,
	}

	p, err := json.Marshal(pedro)
	if err != nil {
		fmt.Println(err)
	}
	k, err := json.Marshal(karol)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(p))
	fmt.Println(string(k))

}
