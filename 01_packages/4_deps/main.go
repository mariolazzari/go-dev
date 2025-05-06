package main

import "github.com/gofrs/uuid/v5"

func main() {
	u, err := uuid.NewV4()
	if err != nil {
		println("Failed to generate UUID:", err)
		return
	}
	println(u.String())
}
