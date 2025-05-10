package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type MyMap map[string]string

type User struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	var myMap MyMap = MyMap{"mario": "Mario"}
	fmt.Println(myMap)

	var user User = User{"Mario", 50}
	fmt.Println(user)

	var maria User = User{
		Name: "Maria",
		Age:  78,
	}
	fmt.Println(maria)

	enc := json.NewEncoder(os.Stdout)

	if err := enc.Encode(maria); err != nil {
		log.Fatal(err)
	}

}
