package main

import "golang.org/x/exp/constraints"

type Store[K constraints.Ordered] struct {
	items []K
}

func main() {

}
