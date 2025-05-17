package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	err := ctx.Err()
	fmt.Println(err)
}
