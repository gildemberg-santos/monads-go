package main

import (
	"fmt"

	"github.com/gildemberg-santos/monads-go/result"
)

func divide(a, b int) *result.Result {
	if b == 0 {
		return result.NewFailure("division by zero")
	}

	return result.NewSuccess(a / b)
}

func main() {
	result := divide(10, 2)
	if result.IsSuccess() {
		fmt.Println(result.GetValue())
	} else {
		fmt.Println(result.GetValue())
	}

	result = divide(10, 0)
	if result.IsSuccess() {
		fmt.Println(result.GetValue())
	} else {
		fmt.Println(result.GetValue())
	}

}
