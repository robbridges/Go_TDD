package main

import (
	"fmt"
	hello "go_tdd/hello_world"
	int "go_tdd/ints"
)

func main() {
	msg := hello.Hello_World("Rob")
	fmt.Println(msg)
	addition := int.AddSum(1, 2)
	fmt.Println(addition)
}
