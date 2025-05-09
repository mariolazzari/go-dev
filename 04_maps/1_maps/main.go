package main

import (
	"fmt"
	"sort"
)

func main() {
	users := map[string]string{}
	users["Mario"] = "mario.lazzari@gmail.com"
	users["Luigi"] = "luigi@gmail.com"

	fmt.Println(users)
	fmt.Println(users["Mario"])
	fmt.Println(len(users))

	// init
	users2 := map[string]string{
		"Mario": "mario",
		"Luigi": "luigi",
		"Peach": "peach",
	}
	fmt.Println(users2)

	// make
	users3 := make(map[string]string)
	users3["Mario"] = "mario"
	users3["Luigi"] = "luigi"
	users3["Peach"] = "peach"
	fmt.Println(users3)

	// order is not guaranteed
	for k, v := range users {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	// delete
	delete(users, "Mario")
	fmt.Println(users)

	// update key
	users["Luigi"] = "Gino"
	fmt.Println(users)

	// extract keys
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}
	keys := make([]int, 0, len(months))
	for k := range months {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println(keys)
}
