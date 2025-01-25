package main

import (
	"fmt"

	"github.com/gildemberg-santos/monads-go/result"
)

func divide(a, b float64) *result.Result {
	if b == 0 {
		return result.NewFailure("divisão por zero")
	}
	return result.NewSuccess(a / b)
}

func main() {
	successResult := divide(10, 2)
	failureResult := divide(10, 0)

	fmt.Println("Success Result:")
	if successResult.IsSuccess() {
		fmt.Println("Resultado:", successResult.GetValue())
	} else {
		fmt.Println("Erro:", successResult.Failure())
	}

	fmt.Println("Failure Result:")
	if failureResult.IsSuccess() {
		fmt.Println("Resultado:", failureResult.GetValue())
	} else {
		fmt.Println("Erro:", failureResult.Failure())
	}
}
