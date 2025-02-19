package cap16

import (
	"encoding/json"
	"fmt"
	"os"
)

type Pessoa struct {
	Nome      string  `json:"Nome"` // tags -> como deve ficar/estar no json
	Sobrenome string  `json:"Sobrenome"`
	Idade     int     `json:"Idade"`
	Profissao string  `json:"Profissao"`
	AccBanco  float64 `json:"AccBanco"`
}

func class03() {

	pedro := Pessoa{"Pedro", "Araujo", 26, "Sec Engineer", 30754.33}
	karol := Pessoa{
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

	jsonBlob := []byte(`
		{"Nome": "Pedro", "Sobrenome": "Araujo", "Idade": 26, "Profissao": "Sec Engineer", "AccBanco": 30754.33}
	`)

	data := Pessoa{}
	err = json.Unmarshal(jsonBlob, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	jsonBlob = []byte(`[
        {"Nome":"Pedro","Sobrenome":"Araujo","Idade":26,"Profissao":"Sec Engineer","AccBanco":30754.33},
        {"Nome":"Karol","Sobrenome":"Sousa","Idade":27,"Profissao":"Data Analyst","AccBanco":345821.21}
    ]`)

	var allData []Pessoa
	err = json.Unmarshal(jsonBlob, &allData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(allData)

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(pedro) // marshall -> variável | encoder -> write.io

}
