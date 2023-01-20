package hello_world

import "fmt"

func Hello_World(name string) string {
	greeting := fmt.Sprintf("hello %s", name)
	return greeting
}
