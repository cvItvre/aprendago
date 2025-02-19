package cap03

import (
	"fmt"
)

//type hotdog int

// var salsicha hotdog
var vina int

func ex05() {

	fmt.Println(salsicha)
	fmt.Printf("%T\n", salsicha)

	salsicha = 42

	fmt.Println(salsicha)

	vina = int(salsicha)

	fmt.Printf("%v, %T", vina, vina)

}
