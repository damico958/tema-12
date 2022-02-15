package calculator

import (
	"com.damico958.devops/operations"
	"log"
)

var History []operations.Operation

func Calculate(firstNumber float64, secondNumber float64, chosenOperation string) float64 {
	var op operations.Operation

	switch chosenOperation {
	case "sum":
		op = operations.Add(firstNumber, secondNumber)
	case "sub":
		op = operations.Sub(firstNumber, secondNumber)
	case "mul":
		op = operations.Mul(firstNumber, secondNumber)
	case "div":
		op = operations.Div(firstNumber, secondNumber)
	default:
		log.Fatal("Fail! Could not proceed with your calculation! Please, choose a valid operation!")
	}

	History = append(History, op)

	return op.Calculate()
}
