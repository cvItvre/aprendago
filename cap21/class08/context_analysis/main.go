package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("contexto err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("-------------------------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("contexto err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("-------------------------------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("contexto err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("-------------------------------")

}
