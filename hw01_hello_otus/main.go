package main

import (
	"fmt"                                //nolint:gci
	"golang.org/x/example/hello/reverse" //nolint:gci
)

func main() {
	const greeting = "Hello, OTUS!"
	fmt.Println(reverse.String(greeting))
}
