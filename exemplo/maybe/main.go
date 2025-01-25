package main

import (
	"fmt"

	"github.com/gildemberg-santos/monads-go/maybe"
)

func main() {
	someValue := maybe.NewSome("Hello, World!")
	noneValue := maybe.NewNone()

	if someValue.IsSome() {
		fmt.Println("Some value:", someValue.GetValue())
	} else {
		fmt.Println("No value")
	}

	if noneValue.IsNone() {
		fmt.Println("No value")
	} else {
		fmt.Println("Some value:", noneValue.GetValue())
	}
}
