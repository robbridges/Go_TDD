package main

import (
	"fmt"
	walk "go_tdd/Reflection"
)

func main() {
	x := struct {
		Name string
		Age  int
	}{"Rob", 35}
	var got []string
	walk.Walk(x, func(input string) {
		got = append(got, input)
	})

	fmt.Printf("We got the value %s back", string(got[0]))

}
