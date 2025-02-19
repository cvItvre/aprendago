package cap11

import (
	"fmt"
)

func ex04() {

	pessoa := struct {
		nome     string
		idade    int
		certs    []string
		rapAlbum map[int]string
	}{
		nome:  "Pedro",
		idade: 27,
		certs: []string{
			"security+",
			"dcpt",
			"ceh",
			"oscp",
		},
		rapAlbum: map[int]string{
			10: "To Pimp A Butterfly",
			9:  "IGOR",
			8:  "Gigantes",
		},
	}

	fmt.Println(pessoa)

}
