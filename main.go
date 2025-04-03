package main

import (
	"fmt"
)

func Hello() string {
	return "Hello World!"
}

func main() {
	fmt.Println(Hello())

	fmt.Println("Oh noes!! There's a terrible bug in v1.0")
	panic("Why ...")
}
