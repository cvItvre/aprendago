package cap10

import (
	"fmt"
)

func class04() {

	x := struct { // struct anonimos
		nome  string
		idade int
	}{
		nome:  "Pedro",
		idade: 27,
	}

	fmt.Println(x)

}
