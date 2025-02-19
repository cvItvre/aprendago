package cap03

import (
	"fmt"
)

type hotdog int

var salsicha hotdog

func ex04() {

	fmt.Println(salsicha)
	fmt.Printf("%T\n", salsicha)

	salsicha = 42

	fmt.Println(salsicha)

}
