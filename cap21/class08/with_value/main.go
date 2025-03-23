package main

import (
	"context"
	"fmt"
)

func main() {

	ctx := context.WithValue(context.Background(), "language", "go")
	foo(ctx, "language")

	ctx = context.WithValue(context.Background(), "cat", "Nina")
	foo(ctx, "cat")

	foo(ctx, "pedro")

}

func foo(ctx context.Context, key string) {
	if v := ctx.Value(key); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found", key)
}
