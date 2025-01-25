package main

import (
	"fmt"

	"github.com/gildemberg-santos/monads-go/maybe"
)

func divide(a, b int) *maybe.Maybe {
	if b == 0 {
		return maybe.NewNone()
	}

	return maybe.NewSome(a / b)
}

func main() {
	result := divide(10, 2)
	if result.IsSome() {
		fmt.Println(result.GetValue())
	} else {
		fmt.Println("division by zero")
	}

	result = divide(10, 0)
	if result.IsSome() {
		fmt.Println(result.GetValue())
	} else {
		fmt.Println("division by zero")
	}
}
