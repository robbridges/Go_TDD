package DependencyInjection

import (
	"fmt"
	"io"
)

func Greet(writer io.Writer, name string) {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	if err != nil {
		fmt.Println(err)
	}
}
