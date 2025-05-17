package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "value")
	fmt.Println(ctx.Value("key"))
}
