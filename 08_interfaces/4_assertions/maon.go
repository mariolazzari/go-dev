package main

import (
	"fmt"
	"io"
)

func WriteNow(i any, s string) error {
	w, ok := i.(io.Writer)
	if !ok {
		return fmt.Errorf("i is not a io.Writer")
	}

	_, err := fmt.Fprintln(w, s)
	if err != nil {
		return err
	}

	return nil
}

func main() {}
