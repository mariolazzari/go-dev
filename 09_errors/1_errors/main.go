package main

import (
	"errors"
	"fmt"
)

func good() error {
	return nil
}

func good2() (int, error) {
	return 1, nil
}

func bad() (error, int) {
	return nil, 1
}

func main() {
	err := errors.New("BOM!")
	fmt.Println(err)

	err2 := fmt.Errorf("Format error")
	fmt.Println(err2)

	err = good()
	good2()
	bad()
}

func Get(key string) (string, error) {
	m := map[string]string{
		"a": "A",
		"b": "B",
	}

	if v, ok := m[key]; ok {
		return v, nil
	}

	return "", fmt.Errorf("Key not found: %s", key)
}
