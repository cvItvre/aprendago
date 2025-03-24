package main

import (
	"errors"
	"fmt"
	"log"
)

type TypeErrNorgateMath struct {
	lat  string
	long string
	err  error
}

func (e TypeErrNorgateMath) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", e.lat, e.long, e.err)
}

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {

	fmt.Println("------Exemplo01------")
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("------Exemplo02------")
	_, err = sqrtWithVarError(-10)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("------Exemplo03------")
	_, err = sqrtWithFormatError(-87.55)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("------Exemplo04------")
	_, err = sqrtWithVarFormatError(-95.40)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("------Exemplo04------")
	_, err = sqrtWithTypeError(-88.66)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number")
	}
	return 42, nil
}

func sqrtWithVarError(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}

func sqrtWithFormatError(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: square root of negative number: %v", f)
	}
	return 42, nil
}

func sqrtWithVarFormatError(f float64) (float64, error) {
	if f < 0 {
		NewErrNorgateMath := fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, NewErrNorgateMath
	}
	return 42, nil
}

func sqrtWithTypeError(f float64) (float64, error) {
	if f < 0 {
		NewErrNorgateMath := fmt.Errorf("norgate math: square root of negative number: %v", f)
		return 0, TypeErrNorgateMath{lat: "90.22N", long: "87.54W", err: NewErrNorgateMath}
	}
	return 42, nil
}
