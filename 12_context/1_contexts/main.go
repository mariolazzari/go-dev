package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Printf("%v\n", ctx)
	fmt.Printf("\t%#v\n", ctx)

}
